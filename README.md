<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Workflows

This repository provides a collection of reusable GitHub Actions workflow
templates. Each subdirectory contains a specific type of workflow or action that
can be used to automate various tasks in your projects.

## Directory Structure

- `autotools/` — Workflows for autotools (configure/make) projects.
- `container/` — Build and push container images to ghcr.io.
- `crate/` — Publish Rust crates to crates.io.
- `format/` — Formatting checks for source files.
- `generic/` — Generic workflows for general automation needs.
- `go/` — Workflows for Go projects.
- `goreleaser/` — Automatic releases via goreleaser for Go projects.
- `meson/` — Workflows for meson projects.
- `rust/` — Workflows for Rust projects.
- `testdata/` — Minimal example projects for end-to-end testing of the build
  actions.

Each action directory contains an `action.yml` file describing the action or
workflow, plus a `README.md` with further details. The `testdata/` directory
holds the fixtures consumed by the repository's own `test-fixtures` workflow.

## Recommended usage

The actions here are meant to be composed inside a project's own
`.github/workflows/`, not invoked in isolation. The layout below is distilled
from real consumers and works well as a starting point.

### Separate CI from CD

Keep continuous integration and continuous delivery in two files.

**`.github/workflows/ci.yaml`** — runs on pushes to the default branch, on PRs,
and on a weekly schedule (so cached dependencies do not rot silently):

```yaml
name: Continuous integration

on:
  push:
    branches:
      - master
  pull_request:
  schedule:
    - cron: "0 0 * * 0" # weekly, catches rotting dependencies
  workflow_dispatch:

jobs:
  generic:
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/generic@master

  format:
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/format@master

  build:
    permissions:
      contents: write
      id-token: write
    runs-on: ubuntu-latest
    steps:
      # Install any project-specific prerequisites the language action does
      # not ship (system libraries, autotools, extra toolchains, ...).
      - uses: black-desk/workflows/<lang>@master # autotools / go / meson / rust

  pass:
    name: pass
    if: always()
    needs: [generic, format, build]
    runs-on: ubuntu-latest
    steps:
      - uses: re-actors/alls-green@release/v1
        with:
          jobs: ${{ toJSON(needs) }}
```

**`.github/workflows/cd.yaml`** — runs on the default branch for both stable and
prerelease versions (see the tagging convention below):

```yaml
name: Continuous deployment

on:
  push:
    branches:
      - master

jobs:
  release:
    permissions:
      contents: write
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/goreleaser@master # or crate / container
```

### Conventions

- **Run `generic` first.** It is the shared housekeeping baseline. Tune it with
  environment variables such as `CLEAN_IGNORE` instead of forking the action.
- **Run `format` separately.** It checks language formatters without scanning
  submodule contents, while still allowing shared formatting config from
  submodules such as `.format`.
- **One job per language/build system.** A project that mixes Go and autotools,
  or Rust and autotools, simply adds one job per toolchain.
- **Install prerequisites before the action.** The language actions handle
  checkout, build, test and coverage — _not_ system packages or extra
  toolchains; provide those in a preceding `run` / `setup` step.
- **Scope permissions per job.** Only jobs that upload coverage (OIDC) or create
  tags need `id-token: write` / `contents: write`; leave `generic` with the
  default token.
- **Gate branch protection on a single `pass` check.** `re-actors/alls-green`
  folds every needed job into one status a branch-protection rule can require.
- **Customise without forking.** Pass `working-directory` for monorepo layouts,
  or environment variables (`GO_TAGS`, `VCPKG_ROOT`, …) that your build system
  already reads.
- **Re-run on demand and on a schedule.** Add `workflow_dispatch` so changes can
  be verified from the Actions tab without a push/PR, and a weekly `schedule` to
  catch upstream action updates no PR picked up.

## Release / tagging convention

The `goreleaser` action uses the same version-driven flow for stable releases
and prereleases, regardless of the project's language:

- Push a project version change to `master`. The action reads the version from
  the project file or version script, creates and pushes its tag, then releases.
- A version such as `0.1.1-rc.2` automatically produces `v0.1.1-rc.2`, just as
  `0.1.1` produces `v0.1.1`. Manual prerelease tags are not required.
- Configure `release.prerelease: auto` in `.goreleaser.yaml` to mark GitHub
  prereleases. Optional crates.io publication uses the Cargo version suffix to
  identify prereleases.
- Only branch push events release; tag push events are ignored. Failed runs can
  be retried using GitHub Actions **Re-run jobs**.

See [the GoReleaser action](goreleaser/README.md) for configuration and retry
behavior. The standalone `crate` action is unchanged and does not provide the
same tag pushing and retry behavior.

## License

This repository is licensed under the MIT License.
