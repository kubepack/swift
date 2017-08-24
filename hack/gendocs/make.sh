#!/usr/bin/env bash

pushd $GOPATH/src/github.com/appscode/swift/hack/gendocs
go run main.go

cd $GOPATH/src/github.com/appscode/swift/docs/reference
sed -i 's/######\ Auto\ generated\ by.*//g' *
popd
