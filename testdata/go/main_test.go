// SPDX-FileCopyrightText: Copyright 2025 Chen Linxuan <me@black-desk.cn>
//
// SPDX-License-Identifier: MIT

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 5, Add(2, 3), "Add(2, 3) should equal 5")
}
