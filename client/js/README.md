# tiller-js-client
Tiller API JavaScript Client

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
