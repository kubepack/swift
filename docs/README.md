[![Go Report Card](https://goreportcard.com/badge/github.com/appscode/wheel)](https://goreportcard.com/report/github.com/appscode/wheel)

# wheel
Ajax friendly [Helm](https://github.com/kubernetes/helm) Tiller proxy using [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).

## API Reference

- **Tiller Version** 
```
GET http://127.0.0.1:9855/tiller/v2/version/json
```

- **Summarize releases** 
```
GET http://127.0.0.1:9855/tiller/v2/releases/json
```

- **List releases** 
```
GET http://127.0.0.1:9855/tiller/v2/releases/list/json
```

- **Release status**
```
GET http://127.0.0.1:9855/tiller/v2/releases/my-release/status/json
```

- **Release content**
```
GET http://127.0.0.1:9855/tiller/v2/releases/my-release/content/json
```

- **Release history**
```
GET http://127.0.0.1:9855/tiller/v2/releases/my-release/json
```

- **Rollback release**
```
GET http://127.0.0.1:9855/tiller/v2/releases/my-release/rollback/json
```

- **Install release from url**

```
POST http://127.0.0.1:9855/tiller/v2/releases/my-release/json

{
	"chart_url": "https://github.com/tamalsaha/test-chart/raw/master/test-chart-0.1.0.tgz",
	"values": {
		"raw": "{\"ns\":\"c10\",\"clusterName\":\"h505\"}"
	}
}
```

- **Install release from stable kubeapps (most recent version)**

```
POST http://127.0.0.1:9855/tiller/v2/releases/my-release/json

{
	"chart_url": "stable/fluent-bit"
}
```

- **Install release from stable kubeapps (specific version)**

```
POST http://127.0.0.1:9855/tiller/v2/releases/my-release/json

{
	"chart_url": "stable/fluent-bit/0.1.2"
}
```

- **Update release**

```
PUT http://127.0.0.1:9855/tiller/v2/releases/my-release/json

{
	"chart_url": "https://github.com/tamalsaha/test-chart/raw/master/test-chart-0.1.0.tgz",
	"values": {
		"raw": "{\"ns\":\"c15\",\"clusterName\":\"h505\"}"
	}
}
```

- **Uninstall release**

```
DELETE http://127.0.0.1:9855/tiller/v2/releases/my-release/json
```

- **Uninstall & purge release**

```
DELETE http://127.0.0.1:9855/tiller/v2/releases/my-release/json?purge=true
```

## Supported Versions
Kubernetes 1.5+ . Helm Tiller server [checks for version compatibility](https://github.com/kubernetes/helm/blob/master/pkg/version/compatible.go#L27). Please pick a version of Wheel that matches your Tiller server.

| Wheel Version | Helm Version |
|---------------|--------------|
| 0.1.0         | 2.5.x        |

## Installation
To install Wheel, please follow the guide [here](/docs/install.md).

## Contribution guidelines
Want to help improve Wheel? Please start [here](/CONTRIBUTING.md).

---

**The wheel server collects anonymous usage statistics to help us learn how the software is being used and how we can improve it. To disable stats collection, run the operator with the flag** `--analytics=false`.

---

## Support
If you have any questions, you can reach out to us.
* [Slack](https://slack.appscode.com)
* [Twitter](https://twitter.com/AppsCodeHQ)
* [Website](https://appscode.com)
