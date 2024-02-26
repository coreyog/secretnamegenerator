package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounts(t *testing.T) {
	lineCount := uint32(strings.Count(string(rawAdj), "\n") + 1)
	assert.Equal(t, AdjectiveCount, lineCount)

	lineCount = uint32(strings.Count(string(rawNoun), "\n") + 1)
	assert.Equal(t, NounCount, lineCount)
}

func TestWordFromList(t *testing.T) {
	words := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}
	list := strings.Join(words, "\n")

	for range len(words) * 10 {
		word := wordFromList([]byte(list), uint32(len(words)))
		assert.Contains(t, words, word)
	}
}
