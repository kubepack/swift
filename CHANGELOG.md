## 0.2.0-rc.0 / 2017.08.31

Swift 0.2.0-rc.0 makes some API changes, fixes chart directory bug and performance improvement for incluster mode.

Download docker image via:
```
docker pull appscode/swift:0.2.0
```

**Noteable Changes**

- Removes `List Releases` API. #24
- Changes `status-codes` parameter type to string in `Summerize Releases` API. Supported values are `UNKNOWN, DEPLOYED, DELETED, SUPERSEDED, FAILED, DELETING`. #33
- Adds option to format release values as json in `Release Content` API. #34
- Makes the chart-archive directory temporary. Now creates separate directory for each install request to download and store chart-archive file and deletes it after request is served. Previously we cached the chart-archive files, but two different releases may have same chart-archive file name with different content. #36
- Mounts an EmptyDir as scratch volume to store chart-archive files for incluster-connector mode. This improves storage performance. #35