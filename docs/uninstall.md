> New to Wheel? Please start [here](/docs/tutorials/README.md).

# Uninstall Wheel
Please follow the steps below to uninstall Wheel:

1. Delete the various objects created for Wheel operator.
```console
$ ./hack/deploy/uninstall.sh
+ kubectl delete deployment -l app=wheel -n kube-system
deployment "wheel" deleted
+ kubectl delete service -l app=wheel -n kube-system
service "wheel" deleted
+ kubectl delete serviceaccount -l app=wheel -n kube-system
No resources found
+ kubectl delete clusterrolebindings -l app=wheel -n kube-system
No resources found
+ kubectl delete clusterrole -l app=wheel -n kube-system
No resources found
```

2. Now, wait several seconds for Wheel to stop running. To confirm that Wheel operator pod(s) have stopped running, run:
```console
$ kubectl get pods --all-namespaces -l app=wheel
```
