package release

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/appscode/log"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"k8s.io/helm/pkg/repo"
)

const (
	WHEEL_ARCHIVE = "/tmp/wheel-archive/"
	INDEX_URL     = "https://kubernetes-charts.storage.googleapis.com/index.yaml"
	INDEX_PATH    = WHEEL_ARCHIVE + "index.yaml"
	STABLE_PREFIX = "stable/"
)

func prepareChart(chartUrl string, values *chart.Config) (*chart.Chart, error) {
	if strings.HasPrefix(chartUrl, STABLE_PREFIX) {
		stableUrl, err := kubeappUrl(chartUrl)
		if err != nil {
			return nil, err
		}
		chartUrl = stableUrl
	}

	log.Infoln("Chart url:", chartUrl)

	chart, err := chartFromUrl(chartUrl, WHEEL_ARCHIVE)
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

func chartFromUrl(url string, dir string) (*chart.Chart, error) {
	if url == "" {
		return nil, errors.New("Url not specified")
	}

	tokens := strings.Split(url, "/")
	fileName := dir + tokens[len(tokens)-1]

	err := downloadFile(url, fileName, false)
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
		return nil, fmt.Errorf("cannot load requirements: %v", err)
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
		return fmt.Errorf("found in requirements.yaml, but missing in charts/ directory: %s", strings.Join(missing, ", "))
	}
	return nil
}

func kubeappUrl(name string) (string, error) {
	err := downloadFile(INDEX_URL, INDEX_PATH, true)
	if err != nil {
		return "", err
	}

	index, err := repo.LoadIndexFile(INDEX_PATH)
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

func downloadFile(url, filePath string, replace bool) error {
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

	log.Infoln("Downloading", url, "to", filePath)

	output, err := os.Create(filePath)
	if err != nil {
		log.Infoln("Error while creating", filePath, "-", err)
		return err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Infoln("Error while downloading", url, "-", err)
		return err
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Infoln("Error while downloading", url, "-", err)
		return err
	}

	log.Infoln(n, "bytes downloaded")
	return nil
}
