[Website](https://appscode.com) • [Slack](https://slack.appscode.com) • [Forum](https://discuss.appscode.com) • [Twitter](https://twitter.com/AppsCodeHQ)

# grpc-seed
Skeleton of gRPC service with HTTP gateway and JS ajax client

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
