// SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>
//
// SPDX-License-Identifier: MIT

package main

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
}
