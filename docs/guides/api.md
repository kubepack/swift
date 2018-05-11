---
title: API Reference
description: API Reference
menu:
  product_swift_0.8.1:
    identifier: guides-apiserver
    name: API Reference
    parent: guides
    weight: 10
product_name: swift
menu_name: product_swift_0.8.1
section_menu_id: guides
---

# API Reference

## Tiller Version
```
GET http://127.0.0.1:9855/tiller/v2/version/json
```

## Summarize releases
```
# List releases with status `DEPLOYED` from all namespaces
GET http://127.0.0.1:9855/tiller/v2/releases/json

# List releases with status `DEPLOYED` from `default` namespace
GET http://127.0.0.1:9855/tiller/v2/releases/json?namespace=default

# List releases from all namespaces for a list of statuses
GET http://127.0.0.1:9855/tiller/v2/releases/json?status_codes=DEPLOYED&&status_codes=DELETED

# List releases from `default` namespace for a list of statuses
GET http://127.0.0.1:9855/tiller/v2/releases/json?namespace=default&status_codes=DEPLOYED&&status_codes=DELETED

# List releases with any status from all namespaces
GET http://127.0.0.1:9855/tiller/v2/releases/json?all=true

Available query parameters:
  namespace=<name of namespace>|EMPTY(for all namespaces)
  sort_by=NAME|LAST_RELEASED
  all=true|false
  sort_order=ASC|DESC
  status_codes=UNKNOWN, DEPLOYED, DELETED, SUPERSEDED, FAILED, DELETING
```

## Release status
```
GET http://127.0.0.1:9855/tiller/v2/releases/my-release/status/json
```

## Release content
```

GET http://127.0.0.1:9855/tiller/v2/releases/my-release/content/json
GET http://127.0.0.1:9855/tiller/v2/releases/my-release/content/json?format_values_as_json=true

```

## Release history
```
GET http://127.0.0.1:9855/tiller/v2/releases/my-release/json?max=10
```

## Rollback release
```
GET http://127.0.0.1:9855/tiller/v2/releases/my-release/rollback/json
```

## Install release from url

```
# Install chart in default namespace
POST http://127.0.0.1:9855/tiller/v2/releases/my-release/json

{
	"chart_url": "https://github.com/tamalsaha/test-chart/raw/master/test-chart-0.1.0.tgz",
	"values": {
		"raw": "{\"ns\":\"c10\",\"clusterName\":\"h505\"}"
	}
}

# Install chart in custom "kube-system" namespace
POST http://127.0.0.1:9855/tiller/v2/releases/my-release/json

{
	"chart_url": "https://github.com/tamalsaha/test-chart/raw/master/test-chart-0.1.0.tgz",
	"namespace": "kube-system",
	"values": {
		"raw": "{\"ns\":\"c10\",\"clusterName\":\"h505\"}"
	}
}

# Install chart in custom "kube-system" namespace with custom values.yaml

## values.yaml
proxy:
  secretToken: mytoken
rbac:
   enabled: false

## convert values.yaml to json format and pass as string in "values.raw"
{
  "proxy": {
    "secretToken": "mytoken"
  },
  "rbac": {
    "enabled": false
  }
}

POST http://127.0.0.1:9855/tiller/v2/releases/my-release/json

{
	"chart_url": "https://github.com/tamalsaha/test-chart/raw/master/test-chart-0.1.0.tgz",
	"namespace": "kube-system",
	"values": {
		"raw": "{ \"proxy\": { \"secretToken\": \"mytoken\" }, \"rbac\": { \"enabled\": false } }"
	}
}
```

## Install release from stable kubeapps (most recent version)

```
POST http://127.0.0.1:9855/tiller/v2/releases/my-release/json

{
	"chart_url": "stable/fluent-bit"
}
```

## Install release from stable kubeapps (specific version)

```
POST http://127.0.0.1:9855/tiller/v2/releases/my-release/json

{
	"chart_url": "stable/fluent-bit/0.1.2"
}
```

## Update release

```
PUT http://127.0.0.1:9855/tiller/v2/releases/my-release/json

{
	"chart_url": "https://github.com/tamalsaha/test-chart/raw/master/test-chart-0.1.0.tgz",
	"values": {
		"raw": "{\"ns\":\"c15\",\"clusterName\":\"h505\"}"
	}
}
```

## Uninstall release

```
DELETE http://127.0.0.1:9855/tiller/v2/releases/my-release/json
```

## Uninstall & purge release

```
DELETE http://127.0.0.1:9855/tiller/v2/releases/my-release/json?purge=true
```
