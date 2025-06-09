<!--
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>
-->

# Continuous integration for golang projects

## Permissions

``` yaml
contents: write
```

## Example

``` yaml
name: Continuous integration for master

on:
  push:
    tag:
      - v*

jobs:
  go:
    runs-on: ubuntu-latest
    permissions:
      contents: write
    steps:
      - uses: black-desk/workflows/goreleaser@master
```
