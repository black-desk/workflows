# Continuous integration for rust projects

## Permissions

``` yaml
contents: write
id-token: write
```

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
