#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT=$DIR/..

rm -rf $ROOT/apis $ROOT/schemas

pushd $GOPATH/src/github.com/appscode/lever/_proto
# copy files
mkdir -p $ROOT/apis $ROOT/schemas
./hack/make.sh js
find . -name '*.gw.js' | cpio -pdm $ROOT/apis
find . -name '*.schema.json' | cpio -pdm $ROOT/schemas
(find . | grep gw.js | xargs rm) || true
# regenerate index.js
$DIR/browserify.py $ROOT

popd
