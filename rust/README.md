<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Continuous integration for rust projects

## Permissions

``` yaml
contents: write
id-token: write
```

## Inputs

| Input | Description | Required | Default |
| --- | --- | --- | --- |
| `working-directory` | Directory to run the build/test commands in. | no | `.` |

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
    permissions:
      contents: write
      id-token: write
    steps:
      - uses: black-desk/workflows/rust@master
```
