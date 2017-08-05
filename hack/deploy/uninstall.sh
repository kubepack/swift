#!/bin/bash
set -x

kubectl delete deployment -l app=wheel -n kube-system
kubectl delete service -l app=wheel -n kube-system

# Delete RBAC objects, if --rbac flag was used.
kubectl delete serviceaccount -l app=wheel -n kube-system
kubectl delete clusterrolebindings -l app=wheel -n kube-system
kubectl delete clusterrole -l app=wheel -n kube-system
