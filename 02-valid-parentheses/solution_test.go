package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_case1(t *testing.T) {
	assert.Equal(t, solve("()"), true, "case 1")
}

func Test_case2(t *testing.T) {
	assert.Equal(t, solve("()[]{}"), true, "case 2")
}

func Test_case3(t *testing.T) {
	assert.Equal(t, solve("(]"), false, "case 3")
}

func Test_case4(t *testing.T) {
	assert.Equal(t, solve("]"), false, "case 3")
}

func Test_case5(t *testing.T) {
	assert.Equal(t, solve("([}}])"), false, "case 3")
}
