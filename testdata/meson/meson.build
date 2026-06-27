# SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>
#
# SPDX-License-Identifier: MIT

project(
  'testdata-meson',
  'c',
  version: '0.1.0',
  default_options: ['warning_level=2'],
)

libfoo = static_library(
  'foo',
  'src/foo.c',
)

test(
  'foo-test',
  executable(
    'foo-test',
    'tests/test_foo.c',
    link_with: libfoo,
  ),
)
