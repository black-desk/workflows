<!--
SPDX-FileCopyrightText: Copyright 2026 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Format checks

Checks project formatting without scanning submodules.

The action checks only files reported by `git ls-files` from the selected
`working-directory`. This lets CI initialize submodules such as `.format` for
shared formatter configuration while avoiding formatting files owned by those
submodules.

## Inputs

| Input                         | Description                                                  | Required | Default  |
| ----------------------------- | ------------------------------------------------------------ | -------- | -------- |
| `working-directory`           | Directory to run formatter checks in.                        | no       | `.`      |
| `clang-format-extra-patterns` | Extra git pathspecs checked by `clang-format`, one per line. | no       | empty    |
| `clang-format-version`        | `clang-format` PyPI package version constraint.              | no       | `22.1.8` |

## Checks

- Rust: `cargo fmt --check` when `Cargo.toml` exists.
- Go: `gofmt -l` on tracked `*.go` files.
- C/C++/Objective-C:
  `clang-format --style=file --fallback-style=none --dry-run --Werror` on
  tracked files with conventional suffixes.
- Python: `black --check --diff` on tracked `*.py` and `*.pyi` files when
  `.black.toml` or `pyproject.toml` exists. `.black.toml` is passed explicitly
  with `--config`.
- Prettier: `prettier --check --ignore-unknown` on tracked web, Markdown, JSON,
  and YAML files when a Prettier config exists.

## Headers Without Suffixes

`clang-format` cannot reliably discover C/C++ headers that have no conventional
suffix. Add project-specific git pathspecs when needed:

```yaml
steps:
  - uses: black-desk/workflows/format@master
    with:
      clang-format-extra-patterns: |
        include/**
        src/**/public/*
```

The extra patterns are still passed through `git ls-files`, so they do not
expand into submodule contents.

## Example

```yaml
name: Formatting

on:
  pull_request:

jobs:
  format:
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/format@master
```
