## Install

```sh
$ git clone https://github.com/TamalSaha/test-chart.git

$ ./bin/helm install test-chart
$ ./bin/helm ls
```

## Upgrade
```sh
helm upgrade --set ns=c2 <RELEASE_NAME> test-chart
```
