<!--
SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Workflows

This repository provides a collection of reusable GitHub Actions workflow templates. Each subdirectory contains a specific type of workflow or action that can be used to automate various tasks in your projects.

## Directory Structure

- `autotools/`  — Workflows for autotools (configure/make) projects.
- `codecov/`    — Report code coverage to Codecov.
- `commitlint/` — Lint commit messages using conventional commits.
- `common/`     — Common reusable workflow components.
- `container/`  — Build and push container images to ghcr.io.
- `crate/`      — Publish Rust crates to crates.io.
- `generic/`    — Generic workflows for general automation needs.
- `go/`         — Workflows for Go projects.
- `goreleaser/` — Automatic releases via goreleaser for Go projects.
- `meson/`      — Workflows for meson projects.
- `rust/`       — Workflows for Rust projects.
- `tag/`        — Automatically generate version tags.
- `testdata/`   — Minimal example projects for end-to-end testing of the build actions.

Each directory contains an `action.yml` file describing the action or workflow, and some have a `README.md` with further details.

## License

This repository is licensed under the MIT License.
