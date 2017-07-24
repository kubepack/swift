[Website](https://appscode.com) • [Slack](https://slack.appscode.com) • [Twitter](https://twitter.com/AppsCodeHQ)

# wheel
Ajax friendly [Helm](https://github.com/kubernetes/helm) Tiller service

## Build Instructions
```sh
# dev build
./hack/make.py

# Install/Update dependency (needs glide)
glide slow
```

## Test api
Run the server, by running:
```sh
seed-apis --v=10
```

Now open the following url:
http://127.0.0.1:50066/_appscode/api/seed/v1beta1/apps/hello?name=tamal
