## Install

```sh
$ helm install test/hello --name=tester
$ helm ls
```

## Upgrade
```sh
helm upgrade --set user=c2 tester test/hello
```
