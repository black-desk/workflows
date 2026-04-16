<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Automatic (pre)release to crates.io

## Permissions

``` yaml
contents: write
```

## Example

``` yaml
name: Automatic release to crates.io

on:
  push:
    tags:
      - v*-*
    branches:
      - master

jobs:
  crate:
    permissions:
      contents: write
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/crate@master
```
