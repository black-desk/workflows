<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Automatic (pre)release via goreleaser

## Permissions

```yaml
contents: write
```

## Project setup

`working-directory` defaults to `.` and selects the directory containing the
GoReleaser configuration and version script or Cargo package. For Go projects,
Go is installed when that directory contains `go.mod`; its dependency/build
cache is enabled only when `go.sum` also exists there.

Version scripts must print the complete tag, such as `v0.3.4` or `v0.3.4-rc.1`,
to stdout. The Cargo fallback adds the `v` prefix automatically. Tags must match
SemVer 2.0.0 with a `v` prefix, including optional prerelease and build metadata
(for example, `v1.2.3-rc.1+build.01`). Invalid versions fail before any tag is
created or pushed.

## Release flow

On a branch push, the action reads the project version using
`./scripts/get-project-version.sh`, `./tools/get-project-version.sh`, or
`cargo pkgid --locked`. It creates and pushes that tag before running
GoReleaser. An existing tag at a different commit skips the release; an existing
tag at the current commit allows a failed release to be retried. Tags are never
overwritten.

Stable and prerelease versions use the same flow for all languages. For example,
changing the project version to `0.1.1-rc.2` and pushing the release branch
automatically creates `v0.1.1-rc.2`; no manual tag is needed. Configure
GoReleaser to mark these GitHub Releases as prereleases:

```yaml
release:
  prerelease: auto
```

Crates.io identifies prereleases by their Cargo version suffix (such as
`0.1.1-rc.2`); there is no separate registry prerelease flag.

Only branch push events release; tag push events are ignored. Use GitHub Actions
**Re-run jobs** to retry a failed run. Run the action only after CI has passed
for the checked-out commit.

## Optional crates.io publishing

Set `publish-crate: "true"` to run `cargo publish --locked` after GoReleaser
succeeds. It defaults to `"false"`, so existing callers do not publish a crate.
This uses GoReleaser OSS with a separate Cargo step, not Pro-only after hooks.
Your `.goreleaser.yaml` must configure the Rust builder for Rust binaries; this
input only enables registry publication. Install any additional build tools,
such as `cargo-zigbuild`, in the project's `before.hooks`.

The release tag must match `v<package.version>` from the selected Cargo package.
Keep `Cargo.lock` up to date when changing versions. Select a non-root package
with `working-directory`; virtual workspace manifests are not supported. The
runner needs `curl` and `jq` (both available on `ubuntu-latest`).

Grant `id-token: write` in addition to `contents: write`. Publish the initial
crate version manually, then configure a crates.io Trusted Publisher with the
repository and the caller's workflow filename. No long-lived crates.io token is
needed. See [the crate action prerequisites](../crate/README.md#prerequisites).

Before requesting credentials, the action checks whether the version is already
published. Existing versions, including yanked versions, are skipped. HTTP 404
means publication is needed; other HTTP errors or invalid responses fail the
job. If Cargo publication fails after GitHub Release succeeds, rerun the
workflow for the same commit. A later branch commit without a version bump skips
release.

Serialize release jobs using a shared concurrency group with
`cancel-in-progress: false`. Tags pushed using `GITHUB_TOKEN` do not trigger a
second workflow; GitHub Release and crates.io publication run in this action.

## Example

```yaml
name: Automatic (pre)release via goreleaser

on:
  push:
    branches:
      - master

concurrency:
  group: release
  cancel-in-progress: false

jobs:
  goreleaser:
    permissions:
      contents: write
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/goreleaser@master
```

## Rust release with crates.io publishing

```yaml
name: Release

on:
  push:
    branches: [master]

concurrency:
  group: release
  cancel-in-progress: false

jobs:
  release:
    permissions:
      contents: write
      id-token: write
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/goreleaser@master
        with:
          publish-crate: "true"
```

The caller must gate release jobs on CI success for the same commit; these
trigger examples do not themselves provide that gate. Branch protection that
requires PR checks does not make a separate CD workflow wait for push CI.
