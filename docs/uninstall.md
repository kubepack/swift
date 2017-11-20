---
title: Uninstall Swift v0.4.0
description: Uninstall of Swift v0.4.0
menu:
  product_swift_0.4.0:
    identifier: uninstall-0.4.0
    name: Uninstall 0.4.0
    parent: getting-started
    weight: 40
product_name: swift
left_menu: product_swift_0.4.0
url: /products/swift/0.4.0/getting-started/uninstall/
section_menu_id: getting-started
---

# Uninstall Swift
Please follow the steps below to uninstall Swift:

1. Delete the various objects created for Swift operator.
```console
$ ./hack/deploy/uninstall.sh
+ kubectl delete deployment -l app=swift -n kube-system
deployment "swift" deleted
+ kubectl delete service -l app=swift -n kube-system
service "swift" deleted
+ kubectl delete serviceaccount -l app=swift -n kube-system
No resources found
+ kubectl delete clusterrolebindings -l app=swift -n kube-system
No resources found
+ kubectl delete clusterrole -l app=swift -n kube-system
No resources found
```

2. Now, wait several seconds for Swift to stop running. To confirm that Swift operator pod(s) have stopped running, run:
```console
$ kubectl get pods --all-namespaces -l app=swift
```
