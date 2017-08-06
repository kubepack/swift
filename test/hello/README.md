## Install

```console
$ helm install test/hello --name=tester
$ helm ls
```

## Upgrade
```console
$ helm upgrade --set user=c2 tester test/hello
```
