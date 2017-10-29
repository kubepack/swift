# Installation Guide

Swift proxy server can connect to [Helm](https://github.com/kubernetes/helm) Tiller gRPC server in a number of different ways depending on the [`--connector`](/docs/reference/swift_run.md) flag.


## `incluster` Connector
Swift can proxy Tiller server running in the same Kubernetes cluster using `incluster` connector.

### Using YAML
Swift can be installed using YAML files includes in the [/hack/deploy](/hack/deploy) folder.

```console
# Install without RBAC roles
$ curl https://raw.githubusercontent.com/appscode/swift/0.4.0/hack/deploy/without-rbac.yaml \
  | kubectl apply -f -


# Install with RBAC roles
$ curl https://raw.githubusercontent.com/appscode/swift/0.4.0/hack/deploy/with-rbac.yaml \
  | kubectl apply -f -
```

For detailed instructions on how to deploy __Swift in a RBAC enabled cluster__, please visit [here](/docs/rbac.md).

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
I0806 04:38:03.261796   20867 logs.go:19] FLAG: --analytics="true"
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
