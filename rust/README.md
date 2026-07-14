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
| `cargo-flags` | Flags passed to `cargo`. | no | `--all-features` |

## Checks

The action installs `rustfmt`, runs `cargo fmt --check`, then runs tests and
coverage with `cargo llvm-cov`.

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

## Passing extra flags

Pass any `cargo` flags via the `cargo-flags` input, which defaults to
`--all-features`.

``` yaml
steps:
  - uses: black-desk/workflows/rust@master
    with:
      cargo-flags: --no-default-features --features serde,tokio
```
