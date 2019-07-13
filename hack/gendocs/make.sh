#!/usr/bin/env bash

pushd $GOPATH/src/kubepack.dev/swift/hack/gendocs
go run main.go
popd
