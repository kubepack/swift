[Website](https://appscode.com) • [Slack](https://slack.appscode.com) • [Twitter](https://twitter.com/AppsCodeHQ)

# wheel
Ajax friendly [Helm](https://github.com/kubernetes/helm) Tiller service

## Build Instructions
```sh
# Install/Update dependency (needs glide)
glide slow

# build & run
./hack/make.py
```

## Test api

Run the server, by running:
```sh
wheel run --v=10
```

Get version:

<http://127.0.0.1:50066/_appscode/api/seed/v1beta1/apps/get-version>

List releases:

<http://127.0.0.1:50066/_appscode/api/seed/v1beta1/apps/list-releases>

Get Release status:

<http://127.0.0.1:50066/_appscode/api/seed/v1beta1/apps/get-release-status?name=wheel-test>

Get Release content:

<http://127.0.0.1:50066/_appscode/api/seed/v1beta1/apps/get-release-content?name=wheel-test>

Get Release history:

<http://127.0.0.1:50066/_appscode/api/seed/v1beta1/apps/get-history?name=wheel-test&&max=3>

Uninstall release:

<http://127.0.0.1:50066/_appscode/api/seed/v1beta1/apps/uninstall-release?name=wheel-test>

Rollback release:

<http://127.0.0.1:50066/_appscode/api/seed/v1beta1/apps/rollback-release?name=wheel-test>

Install release:

<http://127.0.0.1:50066/_appscode/api/seed/v1beta1/apps/install-release?name=wheel-test&chart_url=https://kubernetes-charts.storage.googleapis.com/g2-0.1.0.tgz>

Update release:

<http://127.0.0.1:50066/_appscode/api/seed/v1beta1/apps/update-release>

Run test release:

<http://127.0.0.1:50066/_appscode/api/seed/v1beta1/apps/run-test-release>
