## 0.3.0 / 2017.09.20
Swift 0.3.0 updates dependencies `k8s.io/client-go` to 4.0.0 and `k8s.io/helm` to 2.6.1. There is no user visible change in api.

- Check for returned pods or services before connecting. [\#45](https://github.com/appscode/swift/pull/45) ([tamalsaha](https://github.com/tamalsaha))
- Use client-go 4.0.0 [\#43](https://github.com/appscode/swift/pull/43) ([tamalsaha](https://github.com/tamalsaha))
- Fix command in Developer-guide [\#42](https://github.com/appscode/swift/pull/42) ([the-redback](https://github.com/the-redback))
- Move analytics to common GA project [\#41](https://github.com/appscode/swift/pull/41) ([tamalsaha](https://github.com/tamalsaha))


## 0.2.0 / 2017.09.08
0.2.0-rc.0 is now marked as 0.2.0.


## 0.2.0-rc.0 / 2017.08.31
Wheel has been renamed to Swift, because it gets you Tiller Swift. :) Swift 0.2.0 makes some backward incompatible api changes.

- Removes `List Releases` API. The URL path for this api could conflict with the `GetHistory` api. The replacement should be to use `SummarizeReleases` api. #24
- Changes `status_codes` parameter type to string in `SummarizeReleases` API. Supported values are `UNKNOWN, DEPLOYED, DELETED, SUPERSEDED, FAILED, DELETING`. `all` field has been added to request object of this proto. This allows to get all releases (including deleted but not purged ones) without specifying all the status codes. #33
- Adds option to format release values as json in `ReleaseContent` API. #34
- Deployment scripts now mount an `EmptyDir` at `/tmp` path. This is used as scratch volume to store downloaded chart archives files. The downloaded files are also deleted, after request is served. Previously we cached the chart-archive files, but two different releases may have same chart-archive file name with different content. This improves storage performance. #35, #36


## 0.1.0 / 2017.08.06
First public release of Wheel. To install, please visit [here](/docs/install.md).

 - JSON proxy for Helm Tiller apis. Tested with Helm 2.5.0 .
 - Supports connecting to Tiller server in [3 different modes](/docs/install.md).
 - Install and update api uses chart tarballs instead of binary chart objects.
 - Includes a [npm ready](https://www.npmjs.com/package/@appscode/tiller-js-client) Javascript client using promises.
