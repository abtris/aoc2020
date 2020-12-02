package main

import "testing"

func TestVerifyPassword(t *testing.T) {
	tests := []struct {
		name   string
		input  ElfIssue
		result bool
}{
		{name: "1st", input: ElfIssue{min: 1, max: 3, testValue: "a", password: "abcde"}, result: true},
		{name: "2nd", input: ElfIssue{min: 1, max: 3, testValue: "b", password: "cdefg"}, result: false},
		{name: "3rd", input: ElfIssue{min: 2, max: 9, testValue: "c", password: "ccccccccc"}, result: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := VerifyPassword(test.input.min, test.input.max, test.input.testValue, test.input.password)
			if result != test.result {
				t.Errorf("Expected %v and real %v)", test.result, result)
			}
		})
	}
}
func TestVerifyPasswordNewPolicy(t *testing.T) {
	tests := []struct {
		name   string
		input  ElfIssue
		result bool
}{
		{name: "1st", input: ElfIssue{min: 1, max: 3, testValue: "a", password: "abcde"}, result: true},
		{name: "2nd", input: ElfIssue{min: 1, max: 3, testValue: "b", password: "cdefg"}, result: false},
		{name: "3rd", input: ElfIssue{min: 2, max: 9, testValue: "c", password: "ccccccccc"}, result: false},
		{name: "4rd", input: ElfIssue{min: 2, max: 14, testValue: "m", password: "mnmmmmmmmmmmmgm"}, result: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := VerifyPasswordNewPolicy(test.input.min, test.input.max, test.input.testValue, test.input.password)
			if result != test.result {
				t.Errorf("Expected %v and real %v)", test.result, result)
			}
		})
	}
}
