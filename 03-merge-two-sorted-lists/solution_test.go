package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_case1(t *testing.T) {
	assert.Equal(t, solve([]int{1, 2, 4}, []int{1, 3, 4}), []int{1, 1, 2, 3, 4, 4}, "case 1")
}

func Test_case2(t *testing.T) {
	assert.Equal(t, solve([]int{}, []int{}), []int{}, "case 2")
}

func Test_case3(t *testing.T) {
	assert.Equal(t, solve([]int{}, []int{0}), []int{0}, "case 3")
}
