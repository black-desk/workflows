<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Automatic (pre)release via goreleaser

## Permissions

``` yaml
contents: write
```

## Example

``` yaml
name: Automatic (pre)release via goreleaser

on:
  push:
    tags:
      - v*-*
    branches:
      - master

jobs:
  goreleaser:
    permissions:
      contents: write
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/goreleaser@master
```
