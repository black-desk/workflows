<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Generic github action

## Example

``` yaml
name: Continuous integration for master

on:
  push:
    branches:
      - master

jobs:
  rust:
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/generic@master
```
