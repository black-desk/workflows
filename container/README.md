<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Build a container image then push to ghcr.io

## Permissions

``` yaml
contents: read
packages: write
```

## Example

``` yaml
name: Continuous integration for master

on:
  push:
    branches:
      - master

jobs:
  container:
    permissions:
      contents: read
      packages: write
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/container@master
```
