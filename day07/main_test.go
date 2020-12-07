package main

import "testing"


func TestParseLines(t *testing.T) {
	expected := 4
	data := loadFile("sample")
	output := processData(data, "shiny gold bag")
	real := len(output)
	if real != expected {
		t.Errorf("Expected %v and real %v)", expected, real)
	}
}

