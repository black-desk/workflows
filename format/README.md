<!--
SPDX-FileCopyrightText: Copyright 2026 Chen Linxuan <me@black-desk.cn>

SPDX-License-Identifier: MIT
-->

# Format checks

Checks project formatting without scanning submodules.

The action checks tracked regular non-symlink files reported by `git ls-files`
from the selected `working-directory`. This lets CI initialize submodules such
as `.format` for shared formatter configuration while avoiding formatting files
owned by those submodules.

Submodule gitlinks and tracked symlinks are filtered out before file lists are
passed to formatting tools. This prevents tools such as Prettier from receiving
an explicit submodule directory path, which would make them recursively scan
that submodule.

## Inputs

| Input                         | Description                                                  | Required | Default  |
| ----------------------------- | ------------------------------------------------------------ | -------- | -------- |
| `working-directory`           | Directory to run formatter checks in.                        | no       | `.`      |
| `clang-format-extra-patterns` | Extra git pathspecs checked by `clang-format`, one per line. | no       | empty    |
| `clang-format-version`        | `clang-format` PyPI package version constraint.              | no       | `22.1.8` |

## Checks

- Rust: `cargo fmt --check` when `Cargo.toml` exists.
- Go: `gofmt -l` on tracked regular non-symlink `*.go` files.
- C/C++/Objective-C:
  `clang-format --style=file --fallback-style=none --dry-run --Werror` on
  tracked regular non-symlink files with conventional suffixes.
- Python: `black --check --diff` on tracked regular non-symlink `*.py` and
  `*.pyi` files when `.black.toml` or `pyproject.toml` exists. `.black.toml` is
  passed explicitly with `--config`.
- Prettier: `prettier --check --ignore-unknown` on tracked regular non-symlink
  web, Markdown, JSON, and YAML files when a Prettier config exists.

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

The extra patterns are still passed through tracked regular file discovery, so
they do not expand into submodule contents.

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
