---
title: Changelog | Swift
description: Changelog
menu:
  product_swift_0.7.1:
    identifier: changelog-swift
    name: Changelog
    parent: welcome
    weight: 10
product_name: swift
menu_name: product_swift_0.7.1
section_menu_id: welcome
url: /products/swift/0.7.0/welcome/changelog/
aliases:
  - /products/swift/0.7.0/CHANGELOG/
---

# Change Log

## [0.7.1](https://github.com/appscode/swift/releases/tag/0.7.1) / 2018.02.19
Swift 0.7.1 adds --tiller-timeout flag with default 5 min deadline. There is no user visible change in api.

- Use github.com/pkg/errors [\#99](https://github.com/appscode/swift/pull/99)
- Add --tiller-timeout flag with default 5 min deadline [\#97](https://github.com/appscode/swift/pull/97)
- Rename factory package to connectors [\#96](https://github.com/appscode/swift/pull/96)


## [0.7.0](https://github.com/appscode/swift/releases/tag/0.7.0) / 2018.02.12
Swift 0.7.0 adds support for SSL enable Tiller. `InstallRelease` and `UpdateRelease` apis have been updated in backward compatible manner to support downloading charts from secure chart repository.

- Update installer script [\#95](https://github.com/appscode/swift/pull/95)
- Document Swift SSL options [\#93](https://github.com/appscode/swift/pull/93)
- Support self-signed ca certificate for Tiller [\#92](https://github.com/appscode/swift/pull/92)
- Do not write SETTINGS in response to ACKs in Cmux [\#91](https://github.com/appscode/swift/pull/91)
- Support SSL for chart repository [\#90](https://github.com/appscode/swift/pull/90)
- Pass username/password in chart URL as basic auth header [\#89](https://github.com/appscode/swift/pull/89)
- Update grpc-go to v1.9.2 [\#85](https://github.com/appscode/swift/pull/85)


## [0.6.0](https://github.com/appscode/swift/releases/tag/0.6.0) / 2018.01.23
Swift 0.6.0 updates Helm dependency to 2.8.0. There is no user visible change in api.

- Revendor to Helm 2.8 [\#83](https://github.com/appscode/swift/pull/83)


## [0.5.2](https://github.com/appscode/swift/releases/tag/0.5.2) / 2018.01.06
Swift 0.5.2 closes tunnel after api call for `kubeconfig` connector. We recommend upgrading to this version. There is no user visible change in api.

- Close tunnel after api call [\#79](https://github.com/appscode/swift/pull/79)
- Add front matter for docs [\#81](https://github.com/appscode/swift/pull/81)
- Fix analytics client id detection [\#80](https://github.com/appscode/swift/pull/80)
- Use tunnel tools from kutil [\#78](https://github.com/appscode/swift/pull/78)
- Set ClientID for analytics [\#77](https://github.com/appscode/swift/pull/77)


## [0.5.1](https://github.com/appscode/swift/releases/tag/0.5.1) / 2017.11.27
Swift 0.5.1 fixes connection leakage in proxy server. We recommend upgrading to this version. There is no user visible change in api.

- Close connection after usage. [\#68](https://github.com/appscode/swift/pull/68)
- Make chart namespaced [\#64](https://github.com/appscode/swift/pull/64)
- Add front matter for swift cli [\#69](https://github.com/appscode/swift/pull/69)
- Add Front matter of docs [\#65](https://github.com/appscode/swift/pull/65)


## [0.5.0](https://github.com/appscode/swift/releases/tag/0.5.0) / 2017.10.29
Swift 0.5.0 makes backward incompatible change to `SummarizeReleases` api.

- `SummarizeReleases` api will not set namespace to `default` by default any more. If no namespace is set, it will return releases from all namespaces. To get releases from a given namespace pass query parameter `namespace=<name>`. [\#61](https://github.com/appscode/swift/pull/61)
- Add [tutorial](/docs/rbac.md) for RBAC enabled cluster [\#59](https://github.com/appscode/swift/pull/59)
- Fix service port in installer yamls [\#58](https://github.com/appscode/swift/pull/58)


## [0.4.0](https://github.com/appscode/swift/releases/tag/0.4.0) / 2017.10.24
0.4.0-rc.0 is now marked as 0.4.0.


## [0.4.0](https://github.com/appscode/swift/releases/tag/0.4.0)-rc.0 / 2017.10.16
Swift 0.4.0-rc.0 updates Helm dependency to 2.7.0-rc.1. There is no user visible change in api.


## [0.3.1](https://github.com/appscode/swift/releases/tag/0.3.1) / 2017.09.21
Swift 0.3.1 fixes RBAC issues with installer yamls. There is no user visible change in api.

- Set service account name in RBAC mode. [\#48](https://github.com/appscode/swift/pull/48)
- Update installer RBAC to support listing services [\#47](https://github.com/appscode/swift/pull/47)


## [0.3.0](https://github.com/appscode/swift/releases/tag/0.3.0) / 2017.09.20
Swift 0.3.0 updates dependencies `k8s.io/client-go` to 4.0.0 and `k8s.io/helm` to 2.6.1. There is no user visible change in api.

- Check for returned pods or services before connecting. [\#45](https://github.com/appscode/swift/pull/45)
- Use client-go 4.0.0 [\#43](https://github.com/appscode/swift/pull/43)
- Fix command in Developer-guide [\#42](https://github.com/appscode/swift/pull/42)
- Move analytics to common GA project [\#41](https://github.com/appscode/swift/pull/41)


## [0.2.0](https://github.com/appscode/swift/releases/tag/0.2.0) / 2017.09.08
0.2.0-rc.0 is now marked as 0.2.0.


## [0.2.0](https://github.com/appscode/swift/releases/tag/0.2.0)-rc.0 / 2017.08.31
Wheel has been renamed to Swift, because it gets you Tiller Swift. :) Swift 0.2.0 makes some backward incompatible api changes.

- Removes `List Releases` API. The URL path for this api could conflict with the `GetHistory` api. The replacement should be to use `SummarizeReleases` api. #24
- Changes `status_codes` parameter type to string in `SummarizeReleases` API. Supported values are `UNKNOWN, DEPLOYED, DELETED, SUPERSEDED, FAILED, DELETING`. `all` field has been added to request object of this proto. This allows to get all releases (including deleted but not purged ones) without specifying all the status codes. #33
- Adds option to format release values as json in `ReleaseContent` API. #34
- Deployment scripts now mount an `EmptyDir` at `/tmp` path. This is used as scratch volume to store downloaded chart archives files. The downloaded files are also deleted, after request is served. Previously we cached the chart-archive files, but two different releases may have same chart-archive file name with different content. This improves storage performance. #35, #36


## [0.1.0](https://github.com/appscode/swift/releases/tag/0.1.0) / 2017.08.06
First public release of Wheel. To install, please visit [here](/docs/setup/install.md).

 - JSON proxy for Helm Tiller apis. Tested with Helm 2.5.0 .
 - Supports connecting to Tiller server in [3 different modes](/docs/setup/install.md).
 - Install and update api uses chart tarballs instead of binary chart objects.
 - Includes a [npm ready](https://www.npmjs.com/package/@appscode/tiller-js-client) Javascript client using promises.
