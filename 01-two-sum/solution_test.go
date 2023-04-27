package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_case1(t *testing.T) {
	assert.Equal(t, solve([]int{2, 7, 11, 15}, 9), []int{0, 1}, "case 1")
}

func Test_case2(t *testing.T) {
	assert.Equal(t, solve([]int{3, 2, 4}, 6), []int{1, 2}, "case 2")
}

func Test_case3(t *testing.T) {
	assert.Equal(t, solve([]int{3, 3}, 6), []int{0, 1}, "case 3")
}
