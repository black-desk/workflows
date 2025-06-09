# Build a container image then push to ghcr.io

## Example

``` yaml
name: Continuous integration for master

on:
  push:
    branches:
      - master

jobs:
  container:
    permissions:
      contents: read
      packages: write
    runs-on: ubuntu-latest
    steps:
      - uses: black-desk/workflows/container@master
```
