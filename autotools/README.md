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

## Coverage

This action uploads coverage to Codecov after testing. Codecov looks for
`lcov.info` at the repository root, so your `make test` target must generate
it there, typically by collecting gcov data with `lcov`:

    lcov --capture --directory . --output-file lcov.info

Note that this action runs `./configure` with no extra flags, so coverage
generation must not rely on a `--enable-coverage` configure option — have the
`test` target produce `lcov.info` on its own.

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
  autotools:
    runs-on: ubuntu-latest
    permissions:
      contents: write
      id-token: write
    steps:
      - uses: black-desk/workflows/autotools@master
```
