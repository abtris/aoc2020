package main

import (
	"testing"
)


func TestLines(t *testing.T) {
	tests := []struct {
		name   string
		in      string
		out     string
		outNum	int
}{
		{name: "simple", in:"nop +0", out: "nop", outNum: 0},
		{name: "simple", in:"acc +1", out: "acc", outNum: 1},
		{name: "simple", in:"acc -99", out: "acc", outNum: -99},
		{name: "simple", in:"jmp -3", out: "jmp", outNum: -3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, outNum := parseLine(test.in)
			if out != test.out {
				t.Errorf("Expected %v and real %v)", test.out, out)
			}
			if outNum != test.outNum {
				t.Errorf("Expected %v and real %v)", test.outNum, outNum)
			}
		})
	}
}


func TestSample(t *testing.T) {
	expected := 5
	sample := readFile("sample")
	instructions := readLines(sample)
	real := getAcc(instructions)
	if real != expected {
		t.Errorf("Expected %v and real %v)", expected, real)
	}
}
