package main

import "testing"

func TestRanges(t *testing.T) {
	tests := []struct {
		name   string
		min    int
		max    int
		kind 	 string
		newMin int
		newMax int
}{
		{name: "simple 1", min: 0, max: 7, kind:"L", newMin: 0, newMax: 3},
		{name: "simple 2", min: 0, max: 7, kind:"R", newMin: 4, newMax: 7},
		{name: "simple 3", min: 4, max: 7, kind:"L", newMin: 4, newMax: 5},
		{name: "simple 4", min: 4, max: 5, kind:"R", newMin: 5, newMax: 5},
		{name: "simple 11", min: 0, max: 127, kind:"F", newMin: 0, newMax: 63},
		{name: "simple 12", min: 0, max: 63, kind:"B", newMin: 32, newMax: 63},
		{name: "simple 13", min: 32, max: 63, kind:"F", newMin: 32, newMax: 47},
		{name: "simple 14", min: 32, max: 47, kind:"B", newMin: 40, newMax: 47},
		{name: "simple 15", min: 40, max: 47, kind:"B", newMin: 44, newMax: 47},
		{name: "simple 16", min: 44, max: 47, kind:"F", newMin: 44, newMax: 45},
		{name: "simple 16", min: 44, max: 45, kind:"F", newMin: 44, newMax: 44},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			min, max  := getRange(test.min, test.max, test.kind)
			if min != test.newMin {
				t.Errorf("Expected %v and real %v)", test.newMin, min)
			}
			if max != test.newMax {
				t.Errorf("Expected %v and real %v)", test.newMax, max)
			}
		})
	}
}

func TestGetRow(t *testing.T) {
	tests := []struct {
		name   string
		x      string
		result int
}{
		{name: "simple", x:"BFFFBBF", result: 70},
		{name: "simple", x:"FFFBBBF", result: 14},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := getRow(test.x)
			if result != test.result {
				t.Errorf("Expected %v and real %v)", test.result, result)
			}
		})
	}
}

func TestGetColumn(t *testing.T) {
	tests := []struct {
		name   string
		x      string
		result int
}{
		{name: "simple", x:"RRR", result: 7},
		{name: "simple", x:"RLL", result: 4},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := getColumn(test.x)
			if result != test.result {
				t.Errorf("Expected %v and real %v)", test.result, result)
			}
		})
	}
}

func TestBoarding(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		row		 int
		column int
		result int
}{
		{name: "simple 1", in:"BFFFBBFRRR", row: 70, column: 7, result: 567},
		{name: "simple 2", in:"FFFBBBFRRR", row: 14, column: 7, result: 119},
		{name: "simple 3", in:"BBFFBBFRLL", row: 102, column: 4, result: 820},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Boarding(test.in)
			if result != test.result {
				t.Errorf("Expected %v and real %v)", test.result, result)
			}
		})
	}
}
