package main

import (
	"testing"
)



func TestSimpleTraverseForest(t *testing.T) {
	sample := createForest("sample")
	expected := 7
	slope := slope{right: 3, down: 1}
	real := traverseForest(sample, slope)
	if expected != real {
		t.Fail()
	}
}


func TestTraverseForest(t *testing.T) {
	sample := createForest("sample")
	expected := 336
	slopes := []slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}
	real := 1
	for _, slope := range slopes {
		count := traverseForest(sample, slope)
		real *= count
	}
	if expected != real {
			t.Errorf("Expected %v and real %v)", expected, real)
	}
}


