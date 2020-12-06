package main

import "testing"

func TestLoadFile(t *testing.T) {
		real1, real2 := loadFile("sample")
		expected1 := 11
		expected2 := 6
		if real1 != expected1 {
			t.Errorf("Expected part1 %v and real %v)", expected1, real1)
		}
		if real2 != expected2 {
			t.Errorf("Expected part2 %v and real %v)", expected2, real2)
		}

}
