package main

import (
	"bytes"
	"testing"
)

func TestCounts(t *testing.T) {
	adj := bytes.Split(rawAdj, []byte("\n"))
	if len(adj) != int(AdjectiveCount) {
		t.Fail()
		t.Logf("Adjective Count, Expected: %d, Actual: %d", AdjectiveCount, len(adj))
	}

	nouns := bytes.Split(rawNoun, []byte("\n"))
	if len(nouns) != int(NounCount) {
		t.Fail()
		t.Logf("Noun Count, Expected: %d, Actual: %d", NounCount, len(nouns))
	}
}
