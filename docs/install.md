# Installation Guide

Wheel proxy server can connect to [Helm](https://github.com/kubernetes/helm) Tiller gRPC server in a number of different ways depending on the [`--connector`](/docs/reference/wheel_run.md) flag.


## `incluster` Connector
Wheel can proxy Tiller server running in the same Kubernetes cluster using `incluster` connector.

### Using YAML
Wheel can be installed using YAML files includes in the [/hack/deploy](/hack/deploy) folder.

```console
# Install without RBAC roles
$ curl https://raw.githubusercontent.com/appscode/wheel/master/hack/deploy/without-rbac.yaml \
  | kubectl apply -f -


# Install with RBAC roles
$ curl https://raw.githubusercontent.com/appscode/wheel/master/hack/deploy/with-rbac.yaml \
  | kubectl apply -f -
```

### Verify installation
To check if Wheel operator pods have started, run the following command:
```console
$ kubectl get pods --all-namespaces -l app=wheel --watch
```

Once the operator pods are running, you can cancel the above command by typing `Ctrl+C`.


## `kubeconfig` Connector
Wheel can proxy Tiller server running in a remote Kubernetes cluster using `kubeconfig` connector. In this mode, Wheel open a tunnel between its own pod and Tiller server pod using Kubernetes api server. This is similar to how Helm cli connects to Tiller server today.

For example, if you are running a [Minikube](https://github.com/kubernetes/minikube) cluster locally, you can use the steps below to connect to a Tiller server running inside minikjube cluster from your workstation.

```console
$ minikube start

$ helm init

$ wheel run --v=3 --connector=kubeconfig --kube-context=minikube
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

# Running PostgreSQL
This tutorial will show you how to use KubeDB to run a PostgreSQL database.

## Before You Begin
At first, you need to have a Kubernetes cluster, and the kubectl command-line tool must be configured to communicate with your cluster. If you do not already have a cluster, you can create one by using [Minikube](https://github.com/kubernetes/minikube). 

Now, install KubeDB cli on your workstation and KubeDB operator in your cluster following the steps [here](/docs/install.md).

To keep things isolated, this tutorial uses a separate namespace called `demo` throughout this tutorial. This tutorial will also use a PGAdmin to connect and test PostgreSQL database, once it is running. Run the following command to prepare your cluster for this tutorial:

```console
$ kubectl create -f ./docs/examples/postgres/demo-0.yaml
namespace "demo" created
deployment "pgadmin" created
service "pgadmin" created

$ kubectl get pods -n demo --watch
NAME                      READY     STATUS              RESTARTS   AGE
pgadmin-538449054-s046r   0/1       ContainerCreating   0          13s
pgadmin-538449054-s046r   1/1       Running   0          1m
^C⏎                                                                                                                                                             

$ kubectl get service -n demo
NAME      CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
pgadmin   10.0.0.92    <pending>     80:31188/TCP   1m

$ minikube ip
192.168.99.100
```

Now, open your browser and go to the following URL: _http://{minikube-ip}:{pgadmin-svc-nodeport}_. According to the above example, this URL will be [http://192.168.99.100:31188](http://192.168.99.100:31188). To log into the PGAdmin, use username `admin` and password `admin`.













## `direct` Connector

