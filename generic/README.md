<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Generic github action

This action runs the housekeeping checks shared across all projects. It is
intended to be included in every repository's CI to keep the code clean and
consistent.

The following checks run on every trigger:

- **[black-desk/clean](https://github.com/black-desk/clean-action)** — project-specific cleanup / style rules.
- **[black-desk/up2date](https://github.com/black-desk/up2date-action)** — verifies dependencies are up to date.
- **[REUSE](https://github.com/fsfe/reuse-action)** — checks [REUSE](https://reuse.software/) compliance for licensing metadata.
- **[gitleaks](https://github.com/gitleaks/gitleaks-action)** — scans the repository for leaked secrets.

Additionally, on `pull_request` events it lints commit messages with
[commitlint](https://commitlint.js.org/) using the
`@commitlint/config-conventional` ruleset. If the repository does not provide a
`commitlint.config.*` file, a default config extending `config-conventional` is
generated on the fly.

## Example

``` yaml
name: Continuous integration for master

on:
  push:
    branches:
      - master

jobs:
  generic:
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/generic@master
```
