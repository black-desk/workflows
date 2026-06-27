<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Automatic (pre)release to crates.io

This action publishes to crates.io using [Trusted Publishing](https://crates.io/docs/trusted-publishing)
(OpenID Connect), so it does **not** require a long-lived `CARGO_REGISTRY_TOKEN` secret.

## Permissions

``` yaml
contents: write
id-token: write
```

## Prerequisites

Trusted Publishing authenticates the workflow to crates.io via OIDC, but it can
only publish versions of a crate that already exists on the registry:

1. **Publish the first version manually.** The very first version of a crate
   must be published by hand using an API token. Trusted Publishing can then
   publish all subsequent versions automatically.
2. **Register the crate as a trusted publisher.** On crates.io, open the crate's
   *Settings → Trusted Publishing*, add GitHub, and fill in the repository
   owner, repository name, and the workflow filename (the `.yml` file in your
   project that invokes this action).
3. **Grant `id-token: write`** to the publishing job (see Permissions above).

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
      id-token: write
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/crate@master
```
