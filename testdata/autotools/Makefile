# SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>
#
# SPDX-License-Identifier: MIT

CC := gcc
CFLAGS := -Wall -Wextra --coverage
LDFLAGS := --coverage

.PHONY: all test clean

all: libfoo.a

libfoo.a: src/foo.o
	ar rcs $@ $^

src/foo.o: src/foo.c src/foo.h
	$(CC) $(CFLAGS) -c src/foo.c -o $@

test: all
	$(CC) $(CFLAGS) tests/test_foo.c -L. -lfoo -o test_foo
	./test_foo
	lcov --capture --directory . --output-file lcov.info --no-external

clean:
	rm -f src/*.o libfoo.a test_foo *.gcno *.gcda lcov.info
