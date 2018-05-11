---
title: Release | Swift
description: swift Release
menu:
  product_swift_0.8.1:
    identifier: release
    name: Release
    parent: developer-guide
    weight: 15
product_name: swift
menu_name: product_swift_0.8.1
section_menu_id: setup
---

# Release Process

The following steps must be done from a Linux x64 bit machine.

- Do a global replacement of tags so that docs point to the next release.
- Push changes to the `release-x` branch and apply new tag.
- Push all the changes to remote repo.
- Build and push swift docker image:

```console
$ cd ~/go/src/github.com/appscode/swift
./hack/release.sh
```

- Now, update the release notes in Github. See previous release notes to get an idea what to include there.
