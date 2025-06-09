<!--
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>
-->

# Continuous integration for golang projects

## Permissions

``` yaml
contents: write
id-token: write
```

## Example

``` yaml
name: Continuous integration for master

on:
  push:
    branches:
      - master

jobs:
  go:
    runs-on: ubuntu-latest
    permissions:
      contents: write
      id-token: write
    steps:
      - uses: black-desk/workflows/go@master
```
