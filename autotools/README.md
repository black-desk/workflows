<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Continuous integration for autotools projects

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
  autotools:
    runs-on: ubuntu-latest
    permissions:
      contents: write
      id-token: write
    steps:
      - uses: black-desk/workflows/autotools@master
```
