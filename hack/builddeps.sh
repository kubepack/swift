#!/bin/bash

# echo 'Setting up dependencies for compiling protos...'
# "$(go env GOPATH)"/src/github.com/appscode/swift/_proto/hack/builddeps.sh

echo '---'
echo '--'
echo '.'
echo 'Setting up dependencies for compiling swift...'
# https://github.com/ellisonbg/antipackage
pip install git+https://github.com/ellisonbg/antipackage.git#egg=antipackage

go get -u golang.org/x/tools/cmd/goimports
go get github.com/Masterminds/glide
go get github.com/sgotti/glide-vc
go install github.com/progrium/go-extpoints
