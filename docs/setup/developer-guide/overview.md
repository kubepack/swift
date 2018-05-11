---
title: Overview | Developer Guide
description: Developer Guide Overview
menu:
  product_swift_0.8.1:
    identifier: developer-guide-readme
    name: Overview
    parent: developer-guide
    weight: 15
product_name: swift
menu_name: product_swift_0.8.1
section_menu_id: setup
---

## Development Guide
This document is intended to be the canonical source of truth for things like supported toolchain versions for building Swift.
If you find a requirement that this doc does not capture, please submit an issue on github.

This document is intended to be relative to the branch in which it is found. It is guaranteed that requirements will change over time
for the development branch, but release branches of Swift should not change.

### Build Swift
Some of the Swift development helper scripts rely on a fairly up-to-date GNU tools environment, so most recent Linux distros should
work just fine out-of-the-box.

#### Setup GO
Swift is written in Google's GO programming language. Currently, Swift is developed and tested on **go 1.8.3**. If you haven't set up a GO
development environment, please follow [these instructions](https://golang.org/doc/code.html) to install GO.

#### Download Source

```console
$ go get github.com/appscode/swift
$ cd $(go env GOPATH)/src/github.com/appscode/swift
```

#### Install Dev tools
To install various dev tools for Swift, run the following command:

```console
# setting up dependencies for compiling protobufs...
$ ./_proto/hack/builddeps.sh

# setting up dependencies for compiling swift...
$ ./hack/builddeps.sh
```

Please note that this replaces various tools with specific versions needed to compile swift. You can find the full list [here](https://github.com/appscode/swift/blob/0.8.1/_proto/hack/builddeps.sh#L54.

#### Build Binary
```console
$ ./hack/make.py
$ swift version
```

#### Dependency management
Swift uses [Glide](https://github.com/Masterminds/glide) to manage dependencies. Dependencies are already checked in the `vendor` folder.
If you want to update/add dependencies, run:

```console
$ glide slow
```

#### Build Docker images
To build and push your custom Docker image, follow the steps below. To release a new version of Swift, please follow the [release guide](/docs/setup/developer-guide/release.md).

```console
# Build Docker image
$ ./hack/docker/setup.sh; ./hack/docker/setup.sh push

# Add docker tag for your repository
$ docker tag appscode/swift:<tag> <image>:<tag>

# Push Image
$ docker push <image>:<tag>
```

#### Generate CLI Reference Docs
```console
$ ./hack/gendocs/make.sh
```
