---
title: Install | Swift
description: Installation of Swift
menu:
  product_swift_0.8.1:
    identifier: install
    name: Install
    parent: setup
    weight: 10
product_name: swift
menu_name: product_swift_0.8.1
section_menu_id: setup
---

# Installation Guide

Swift proxy server can connect to [Helm](https://github.com/kubernetes/helm) Tiller gRPC server in a number of different ways depending on the [`--connector`](/docs/reference/swift_run.md) flag.


## `incluster` Connector
Swift can proxy Tiller server running in the same Kubernetes cluster using `incluster` connector.

Swift operator can be installed via a script or as a Helm chart.

<ul class="nav nav-tabs" id="installerTab" role="tablist">
  <li class="nav-item">
    <a class="nav-link active" id="script-tab" data-toggle="tab" href="#script" role="tab" aria-controls="script" aria-selected="true">Script</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" id="helm-tab" data-toggle="tab" href="#helm" role="tab" aria-controls="helm" aria-selected="false">Helm</a>
  </li>
</ul>
<div class="tab-content" id="installerTabContent">
  <div class="tab-pane fade show active" id="script" role="tabpanel" aria-labelledby="script-tab">

### Using Script
Swift can be installed via installer script included in the [/hack/deploy](https://github.com/appscode/swift/tree/0.8.1/hack/deploy) folder.

```console
$ curl -fsSL https://raw.githubusercontent.com/appscode/swift/0.8.1/hack/deploy/swift.sh | bash
```

#### Customizing Installer

You can see the full list of flags available to installer using `-h` flag.

```console
$ curl -fsSL https://raw.githubusercontent.com/appscode/swift/0.8.1/hack/deploy/swift.sh | bash -s -- -h
swift.sh - install Ajax friendly Helm Tiller Proxy

swift.sh [options]

options:
-h, --help                         show brief help
-n, --namespace=NAMESPACE          specify namespace (default: kube-system)
    --rbac                         create RBAC roles and bindings (default: true)
    --docker-registry              docker registry used to pull swift images (default: appscode)
    --image-pull-secret            name of secret used to pull swift operator images
    --run-on-master                run swift operator on master
    --enable-analytics             send usage events to Google Analytics (default: true)
    --uninstall                    uninstall swift
```

If you would like to run Swift operator pod in `master` instances, pass the `--run-on-master` flag:

```console
$ curl -fsSL https://raw.githubusercontent.com/appscode/swift/0.8.1/hack/deploy/swift.sh \
    | bash -s -- --run-on-master [--rbac]
```

Swift operator will be installed in a `kube-system` namespace by default. If you would like to run Swift operator pod in `swift` namespace, pass the `--namespace=swift` flag:

```console
$ kubectl create namespace swift
$ curl -fsSL https://raw.githubusercontent.com/appscode/swift/0.8.1/hack/deploy/swift.sh \
    | bash -s -- --namespace=swift [--run-on-master] [--rbac]
```

If you are using a private Docker registry, you need to pull the following docker image:

 - [appscode/swift](https://hub.docker.com/r/appscode/swift)

To pass the address of your private registry and optionally a image pull secret use flags `--docker-registry` and `--image-pull-secret` respectively.

```console
$ kubectl create namespace swift
$ curl -fsSL https://raw.githubusercontent.com/appscode/swift/0.8.1/hack/deploy/swift.sh \
    | bash -s -- --docker-registry=MY_REGISTRY [--image-pull-secret=SECRET_NAME] [--rbac]
```

For detailed instructions on how to deploy __Swift in a RBAC enabled cluster__, please visit [here](/docs/setup/rbac.md).

</div>
<div class="tab-pane fade" id="helm" role="tabpanel" aria-labelledby="helm-tab">

### Using Helm
Swift can be installed via [Helm](https://helm.sh/) using the [chart](https://github.com/appscode/swift/tree/0.8.1/chart/swift) from [AppsCode Charts Repository](https://github.com/appscode/charts). To install the chart with the release name `my-release`:

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search appscode/swift
NAME            CHART VERSION APP VERSION DESCRIPTION
appscode/swift  0.8.1         0.8.1       Swift by AppsCode - Ajax friendly Helm Tiller Proxy

$ helm install appscode/swift --name swift --version 0.8.1
```

To see the detailed configuration options, visit [here](https://github.com/appscode/swift/tree/0.8.1/chart/swift/).

</div>

### Verify installation
To check if Swift proxy pods have started, run the following command:

```console
$ kubectl get pods --all-namespaces -l app=swift --watch
```

Once the proxy pods are running, you can cancel the above command by typing `Ctrl+C`.


## `kubeconfig` Connector
Swift can proxy Tiller server running in a remote Kubernetes cluster using `kubeconfig` connector. In this mode, Swift open a tunnel between its own pod and Tiller server pod using Kubernetes api server. This is similar to how Helm cli connects to Tiller server today.

For example, if you are running a [Minikube](https://github.com/kubernetes/minikube) cluster locally, you can use the steps below to connect to a Tiller server running inside minikjube cluster from your workstation.

```console
$ minikube start

$ helm init

$ swift run --v=3 --connector=kubeconfig --kube-context=minikube
I0806 04:38:03.261749   20867 logs.go:19] FLAG: --alsologtostderr="false"
I0806 04:38:03.261796   20867 logs.go:19] FLAG: --enable-analytics="true"
I0806 04:38:03.261809   20867 logs.go:19] FLAG: --api-domain=""
I0806 04:38:03.261826   20867 logs.go:19] FLAG: --caCertFile=""
I0806 04:38:03.261835   20867 logs.go:19] FLAG: --certFile=""
I0806 04:38:03.261845   20867 logs.go:19] FLAG: --connector="kubeconfig"
I0806 04:38:03.261853   20867 logs.go:19] FLAG: --cors-origin-allow-subdomain="false"
I0806 04:38:03.261862   20867 logs.go:19] FLAG: --cors-origin-host=""
I0806 04:38:03.261870   20867 logs.go:19] FLAG: --enable-cors="false"
I0806 04:38:03.261879   20867 logs.go:19] FLAG: --enable-java-client="false"
I0806 04:38:03.261888   20867 logs.go:19] FLAG: --help="false"
I0806 04:38:03.261897   20867 logs.go:19] FLAG: --keyFile=""
I0806 04:38:03.261905   20867 logs.go:19] FLAG: --kube-context="minikube"
I0806 04:38:03.261916   20867 logs.go:19] FLAG: --log_backtrace_at=":0"
I0806 04:38:03.261925   20867 logs.go:19] FLAG: --log_dir=""
I0806 04:38:03.261935   20867 logs.go:19] FLAG: --logtostderr="true"
I0806 04:38:03.261943   20867 logs.go:19] FLAG: --plaintext-addr=":9855"
I0806 04:38:03.261952   20867 logs.go:19] FLAG: --secure-addr=":50055"
I0806 04:38:03.261961   20867 logs.go:19] FLAG: --stderrthreshold="2"
I0806 04:38:03.261970   20867 logs.go:19] FLAG: --tiller-endpoint=""
I0806 04:38:03.261978   20867 logs.go:19] FLAG: --v="3"
I0806 04:38:03.261987   20867 logs.go:19] FLAG: --vmodule=""
I0806 04:38:03.261996   20867 logs.go:19] FLAG: --web-addr=":56790"
I0806 04:38:03.409223   20867 server.go:241] Configuration: {
  "SecureAddr": ":50055",
  "PlaintextAddr": ":9855",
  "EnableJavaClient": false,
  "APIDomain": "",
  "CACertFile": "",
  "CertFile": "",
  "KeyFile": "",
  "EnableCORS": false,
  "CORSOriginHost": "",
  "CORSOriginAllowSubdomain": false,
  "WebAddr": ":56790",
  "EnableAnalytics": true,
  "Connector": "kubeconfig",
  "TillerEndpoint": "",
  "KubeContext": "minikube"
 }
I0806 04:38:03.409512   20867 server.go:89] [PROXYSERVER] Sarting Proxy Server at port [::]:9855
I0806 04:38:03.409591   20867 server.go:168] Registering endpoint: RegisterReleaseServiceHandlerFromEndpoint
I0806 04:38:03.409506   20867 server.go:83] [GRPCSERVER] Starting gRPC Server at addr [::]:9855
I0806 04:38:03.409737   20867 server.go:120] Registering server: *release.Server
```

## `direct` Connector
Swift can proxy Tiller server by directly connecting to it using `direct` connector.
```console
$ swift run --v=3 --connector=direct --tiller-endpoint=http://127.0.0.1:44134
```
