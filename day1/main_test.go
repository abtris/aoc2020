package main

import (
	"testing"
)

func TestTwoNumbers(t *testing.T) {
	expected := 514579
	input := []int{1721,979,366,299,675,1456}
	if SumAndMultiply(input) != expected {
		t.Fail();
	}
}


func TestThreeNumbers(t *testing.T) {
	expected := 241861950
	input := []int{1721,979,366,299,675,1456}
	if SumAndMultiplySecond(input) != expected {
		t.Fail();
	}
}
