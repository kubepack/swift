#!/usr/bin/env bash

pushd $GOPATH/src/github.com/appscode/wheel/hack/gendocs
go run main.go

cd $GOPATH/src/github.com/appscode/wheel/docs/reference
sed -i 's/######\ Auto\ generated\ by.*//g' *
popd
