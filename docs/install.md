# Installation Guide

Wheel proxy server can connect to [Helm](https://github.com/kubernetes/helm) Tiller gRPC server in a number of different ways depending on the `[--connector](/docs/reference/wheel_run.md)` flag.


## `incluster` Connector

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


## `direct` Connector

