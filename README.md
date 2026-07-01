<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Workflows

This repository provides a collection of reusable GitHub Actions workflow templates. Each subdirectory contains a specific type of workflow or action that can be used to automate various tasks in your projects.

## Directory Structure

- `autotools/`  — Workflows for autotools (configure/make) projects.
- `container/`  — Build and push container images to ghcr.io.
- `crate/`      — Publish Rust crates to crates.io.
- `generic/`    — Generic workflows for general automation needs.
- `go/`         — Workflows for Go projects.
- `goreleaser/` — Automatic releases via goreleaser for Go projects.
- `meson/`      — Workflows for meson projects.
- `rust/`       — Workflows for Rust projects.
- `testdata/`   — Minimal example projects for end-to-end testing of the build actions.

Each action directory contains an `action.yml` file describing the action or workflow, plus a `README.md` with further details. The `testdata/` directory holds the fixtures consumed by the repository's own `test-fixtures` workflow.

## Release / tagging convention

The release actions (`crate`, `goreleaser`) share a common tagging flow (inlined into each release action):

- **Stable release** — push to `master`. The release action reads the version from the project file (e.g. `Cargo.toml`) and, if it is a new version, creates the tag automatically, then publishes.
- **Pre-release** — push a tag like `v1.0.0-rc.1` by hand. The release action publishes that version (driven by the Git ref, since the tag already exists).

## License

This repository is licensed under the MIT License.
