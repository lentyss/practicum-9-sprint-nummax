package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		size     int
		hasError bool
	}{
		{1000, false},
		{0, true},
		{-1, true},
		{100_000_001, true},
	}

	for _, m := range tests {
		result, err := generateRandomElements(m.size)

		if m.hasError {
			assert.Error(t, err)
			assert.Nil(t, result)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, m.size, len(result))
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		hasError bool
	}{
		{[]int{}, 0, true},
		{[]int{42}, 42, false},
		{[]int{10, 5, 20, 15}, 20, false},
		{[]int{0, 0, 0, 0}, 0, false},
		{[]int{-5, -10, -3, -1}, -1, false},
		{[]int{3, 1, 4, 1, 5, 9, 2, 6}, 9, false},
	}

	for _, m := range tests {
		result, err := maximum(m.input)

		if m.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, m.expected, result)
		}
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		hasError bool
	}{
		{[]int{}, 0, true},
		{[]int{42}, 42, false},
		{[]int{1, 2, 3, 4, 5}, 5, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 16, false},
	}

	for _, m := range tests {
		result, err := maxChunks(m.input)

		if m.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, m.expected, result)
		}
	}
}

func TestMaximumAndMaxChunksEqual(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{5, 3, 8, 1, 9, 2, 7, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{100, 200, 300, 400, 500, 600, 700, 800}},
	}

	for _, m := range tests {
		resultSingle, err1 := maximum(m.input)
		resultChunks, err2 := maxChunks(m.input)

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.Equal(t, resultSingle, resultChunks)
	}
}
