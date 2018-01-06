[![Go Report Card](https://goreportcard.com/badge/github.com/appscode/swift)](https://goreportcard.com/report/github.com/appscode/swift)

# Swift
Swift is an Ajax friendly [Helm](https://github.com/kubernetes/helm) Tiller proxy using [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway). It was previously called Wheel.


## Supported Versions
Kubernetes 1.5+ . Helm Tiller server [checks for version compatibility](https://github.com/kubernetes/helm/blob/master/pkg/version/compatible.go#L27). Please pick a version of Swift that matches your Tiller server.

| Swift Version                                                 | Docs                                                            | Helm/Tiller Version |
|---------------------------------------------------------------|-----------------------------------------------------------------|---------------------|
| [0.5.2](https://github.com/appscode/swift/releases/tag/0.5.2) | [User Guide](https://appscode.com/products/swift/0.5.2/)        | 2.7.0               |
| [0.3.1](https://github.com/appscode/swift/releases/tag/0.3.1) | [User Guide](https://github.com/appscode/swift/tree/0.3.1/docs) | 2.5.x, 2.6.x        |
| [0.2.0](https://github.com/appscode/swift/releases/tag/0.2.0) | [User Guide](https://github.com/appscode/swift/tree/0.2.0/docs) | 2.5.x, 2.6.x        |
| [0.1.0](https://github.com/appscode/swift/releases/tag/0.1.0) | [User Guide](https://github.com/appscode/swift/tree/0.1.0/docs) | 2.5.x, 2.6.x        |


## Installation
To install Swift, please follow the guide [here](https://appscode.com/products/swift/0.5.2/setup/install/).

## Using Swift
Want to learn how to use Swift? Please start [here](https://appscode.com/products/swift/0.5.2/).

## Contribution guidelines
Want to help improve Swift? Please start [here](https://appscode.com/products/swift/0.5.2/welcome/contributing/).

---

**Swift server collects anonymous usage statistics to help us learn how the software is being used and how we can improve it. To disable stats collection, run the operator with the flag** `--analytics=false`.

---

## Support
We use Slack for public discussions. To chit chat with us or the rest of the community, join us in the [AppsCode Slack team](https://appscode.slack.com/messages/C0XQFLGRM/details/) channel `#general`. To sign up, use our [Slack inviter](https://slack.appscode.com/).

If you have found a bug with Searchlight or want to request for new features, please [file an issue](https://github.com/appscode/swift/issues/new).
