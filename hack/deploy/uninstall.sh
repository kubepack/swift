#!/bin/bash
set -x

kubectl delete deployment -l app=swift -n kube-system
kubectl delete service -l app=swift -n kube-system

# Delete RBAC objects, if --rbac flag was used.
kubectl delete serviceaccount -l app=swift -n kube-system
kubectl delete clusterrolebindings -l app=swift -n kube-system
kubectl delete clusterrole -l app=swift -n kube-system
