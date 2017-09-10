#!/usr/bin/env bash

pushd $GOPATH/src/github.com/appscode/swift/hack/gendocs
go run main.go
popd
