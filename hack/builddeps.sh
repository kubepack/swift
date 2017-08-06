#!/bin/bash

set -x

echo 'Setting up dependencies for compiling protos...'
$(GO env GOPATH)/src/github.com/appscode/wheel/_proto/hack/builddeps.sh

echo '---'
echo '--'
echo '.'
echo 'Setting up dependencies for compiling wheel...'
# https://github.com/ellisonbg/antipackage
pip install git+https://github.com/ellisonbg/antipackage.git#egg=antipackage

go get -u golang.org/x/tools/cmd/goimports
go get github.com/Masterminds/glide
go get github.com/sgotti/glide-vc
go install github.com/progrium/go-extpoints
