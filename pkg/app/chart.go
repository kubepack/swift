/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/appscode/log"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/proto/hapi/chart"
)

// req.ChartUrl = "https://kubernetes-charts.storage.googleapis.com/g2-0.1.0.tgz" //kubeapps-g2-url

func prepareChart(chartUrl string, values *chart.Config) (*chart.Chart, error) {
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

	if _, err := os.Stat(fileName); err == nil {
		log.Infoln("File already exists:", fileName)
		return chartFromPath(fileName)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, 0777); err != nil {
			return nil, err
		}
	}

	log.Infoln("Downloading", url, "to", fileName)

	output, err := os.Create(fileName)
	if err != nil {
		log.Infoln("Error while creating", fileName, "-", err)
		return nil, err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Infoln("Error while downloading", url, "-", err)
		return nil, err
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Infoln("Error while downloading", url, "-", err)
		return nil, err
	}

	log.Infoln(n, "bytes downloaded")

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
