package release

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	urllib "net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/appscode/go/log"
	"github.com/pkg/errors"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"k8s.io/helm/pkg/repo"
)

const (
	TMP_DIR       = "/tmp"
	DIR_PREFIX    = "swift"
	INDEX_URL     = "https://kubernetes-charts.storage.googleapis.com/index.yaml"
	INDEX_FILE    = "index.yaml"
	STABLE_PREFIX = "stable/"
)

type chartInfo struct {
	ChartURL           string
	CaBundle           []byte
	Username           string
	Password           string
	Token              string
	ClientCertificate  []byte
	ClientKey          []byte
	InsecureSkipVerify bool
}

func prepareChart(repo chartInfo, values *chart.Config) (*chart.Chart, error) {
	// create tmp dir for kupbeapp index file and chart archive file
	chartDir, err := ioutil.TempDir(TMP_DIR, DIR_PREFIX)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(chartDir) // clean up tmp dir
	log.Infoln("Chart dir:", chartDir)

	if strings.HasPrefix(repo.ChartURL, STABLE_PREFIX) {
		stableUrl, err := kubeappUrl(repo.ChartURL, chartDir)
		if err != nil {
			return nil, err
		}
		repo.ChartURL = stableUrl
	}
	log.Infoln("Chart url:", repo.ChartURL)

	chart, err := pullChart(repo, chartDir)
	if err != nil {
		return nil, err
	}

	// ref: k8s.io/helm/pkg/helm/client.go (func InstallReleaseFromChart)
	err = chartutil.ProcessRequirementsEnabled(chart, values)
	if err != nil {
		return nil, err
	}
	err = chartutil.ProcessRequirementsImportValues(chart)
	if err != nil {
		return nil, err
	}

	return chart, nil
}

func pullChart(repo chartInfo, dir string) (*chart.Chart, error) {
	if repo.ChartURL == "" {
		return nil, errors.New("url not specified")
	}

	tokens := strings.Split(repo.ChartURL, "/")
	fileName := filepath.Join(dir, tokens[len(tokens)-1])

	err := downloadFile(repo, fileName, true)
	if err != nil {
		return nil, err
	}

	return chartFromPath(fileName)
}

// ref: k8s.io/helm/cmd/helm/install.go (func run)
func chartFromPath(path string) (*chart.Chart, error) {
	chartRequested, err := chartutil.Load(path)
	if err != nil {
		return nil, err
	}

	if req, err := chartutil.LoadRequirements(chartRequested); err == nil {
		if err := checkDependencies(chartRequested, req); err != nil {
			return nil, err
		}
	} else if err != chartutil.ErrRequirementsNotFound {
		return nil, errors.Wrap(err, "cannot load requirements")
	}

	return chartRequested, nil
}

// ref: k8s.io/helm/cmd/helm/install.go (func checkDependencies)
func checkDependencies(ch *chart.Chart, reqs *chartutil.Requirements) error {
	missing := []string{}

	deps := ch.GetDependencies()
	for _, r := range reqs.Dependencies {
		found := false
		for _, d := range deps {
			if d.Metadata.Name == r.Name {
				found = true
				break
			}
		}
		if !found {
			missing = append(missing, r.Name)
		}
	}

	if len(missing) > 0 {
		return errors.Errorf("found in requirements.yaml, but missing in charts/ directory: %s", strings.Join(missing, ", "))
	}
	return nil
}

func kubeappUrl(name string, dir string) (string, error) {
	indexPath := filepath.Join(dir, INDEX_FILE)
	err := downloadFile(chartInfo{ChartURL: INDEX_URL}, indexPath, true)
	if err != nil {
		return "", err
	}

	index, err := repo.LoadIndexFile(indexPath)
	if err != nil {
		return "", err
	}

	name = strings.TrimPrefix(name, STABLE_PREFIX)
	version := ""

	tokens := strings.Split(name, "/")
	if len(tokens) > 1 {
		name = tokens[0]
		version = tokens[1]
	}

	log.Infoln("kubeApp:", name, version)

	chartVersion, err := index.Get(name, version)
	if err != nil {
		return "", err
	}

	return chartVersion.URLs[0], nil
}

// filePath: /tmp/swift311308022/index.yaml
// filePath: /tmp/swift311308022/test-chart-0.1.0.tgz
func downloadFile(repo chartInfo, filePath string, replace bool) error {
	if !replace {
		if _, err := os.Stat(filePath); err == nil {
			log.Infoln("File already exists:", filePath)
			return nil
		}
	}

	dir := path.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, 0777); err != nil {
			return err
		}
	}

	log.Infoln("Downloading", repo.ChartURL, "to", filePath)

	output, err := os.Create(filePath)
	if err != nil {
		log.Infoln("Error while creating", filePath, "-", err)
		return err
	}
	defer output.Close()

	u, err := urllib.Parse(repo.ChartURL)
	if err != nil {
		log.Infof("failed to parse url. reason: %s", err)
		return err
	}

	var auth string
	if repo.Token != "" {
		auth = "Bearer " + repo.Token
	} else if repo.Username != "" && repo.Password != "" {
		auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(repo.Username+":"+repo.Password))
	} else if u.User != nil {
		username := u.User.Username()
		password, _ := u.User.Password()
		auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
		u.User = nil
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	if auth != "" {
		req.Header.Set("Authorization", auth)
	}
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	tlsConfig := &tls.Config{}
	if repo.InsecureSkipVerify {
		tlsConfig.InsecureSkipVerify = true
	} else {
		if len(repo.CaBundle) > 0 {
			pool := x509.NewCertPool()
			pool.AppendCertsFromPEM(repo.CaBundle)
			tlsConfig.RootCAs = pool
		}
		if len(repo.ClientCertificate) > 0 && len(repo.ClientKey) > 0 {
			pair, err := tls.X509KeyPair(repo.ClientCertificate, repo.ClientKey)
			if err != nil {
				return errors.Wrap(err, "invalid client cert pair")
			}
			tlsConfig.Certificates = []tls.Certificate{pair}
		}
	}

	// ref: https://github.com/golang/go/blob/release-branch.go1.9/src/net/http/transport.go#L40
	var transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	}
	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		log.Infoln("Error while downloading", u.String(), "-", err)
		return err
	}
	defer resp.Body.Close()

	n, err := io.Copy(output, resp.Body)
	if err != nil {
		log.Infoln("Error while downloading", u.String(), "-", err)
		return err
	}

	log.Infoln(n, "bytes downloaded")
	return nil
}
