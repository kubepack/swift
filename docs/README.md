---
title: Readme | Swift
description: Readme of Swift v0.6.0
menu:
  product_swift_0.6.0:
    identifier: readme-0.6.0
    name: Overview
    parent: getting-started
    weight: 20
product_name: swift
left_menu: product_swift_0.6.0
aliases:
  - /products/swift/0.6.0/
url: /products/swift/0.6.0/getting-started/
section_menu_id: getting-started
---


[![Go Report Card](https://goreportcard.com/badge/github.com/appscode/swift)](https://goreportcard.com/report/github.com/appscode/swift)

# swift
Swift is an Ajax friendly [Helm](https://github.com/kubernetes/helm) Tiller proxy using [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway). It was previously called Wheel.

## API Reference

- **Summarize releases**
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

- **Release status**
```
GET http://127.0.0.1:9855/tiller/v2/releases/my-release/status/json
```

- **Release content**
```

GET http://127.0.0.1:9855/tiller/v2/releases/my-release/content/json
GET http://127.0.0.1:9855/tiller/v2/releases/my-release/content/json?format_values_as_json=true

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
Kubernetes 1.5+ . Helm Tiller server [checks for version compatibility](https://github.com/kubernetes/helm/blob/master/pkg/version/compatible.go#L27). Please pick a version of Swift that matches your Tiller server.

| Swift Version                                                           | Docs                                                                 | Helm/Tiller Version |
|-------------------------------------------------------------------------|----------------------------------------------------------------------|---------------------|
| [0.5.0](https://github.com/appscode/swift/releases/tag/0.5.0)           | [User Guide](https://github.com/appscode/swift/tree/0.5.0/docs)      | 2.7.0               |
| [0.3.1](https://github.com/appscode/swift/releases/tag/0.3.1)           | [User Guide](https://github.com/appscode/swift/tree/0.3.1/docs)      | 2.5.x, 2.6.x        |
| [0.2.0](https://github.com/appscode/swift/releases/tag/0.2.0)           | [User Guide](https://github.com/appscode/swift/tree/0.2.0/docs)      | 2.5.x, 2.6.x        |
| [0.1.0](https://github.com/appscode/swift/releases/tag/0.1.0)           | [User Guide](https://github.com/appscode/swift/tree/0.1.0/docs)      | 2.5.x, 2.6.x        |


## Installation
To install Swift, please follow the guide [here](/docs/install.md).

## Contribution guidelines
Want to help improve Swift? Please start [here](/CONTRIBUTING.md).

---

**The swift server collects anonymous usage statistics to help us learn how the software is being used and how we can improve it. To disable stats collection, run the operator with the flag** `--analytics=false`.

---

## Support
If you have any questions, you can reach out to us.
* [Slack](https://slack.appscode.com)
* [Twitter](https://twitter.com/AppsCodeHQ)
