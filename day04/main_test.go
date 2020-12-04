package main

import "testing"

func TestValidPassport(t *testing.T) {
		real := createPassport("sample", false)
		expected := 3
		if real != expected {
			t.Errorf("Expected %v and real %v)", expected, real)
		}
}

func TestValidationValues(t *testing.T) {
	tests := []struct {
		name   string
		x      string
		y      string
		result bool
}{
		{name: "byr valid", x:"byr", y: "2002", result: true},
		{name: "byr invalid", x:"byr", y: "2003", result: false},
		{name: "hgt valid", x:"hgt", y: "60in", result: true},
		{name: "hgt valid", x:"hgt", y: "190cm", result: true},
		{name: "hgt invalid", x:"hgt", y: "190in", result: false},
		{name: "hgt invalid", x:"hgt", y: "190", result: false},
		{name: "hcl valid", x:"hcl", y: "#123abc", result: true},
		{name: "hcl invalid", x:"hcl", y: "#123abz", result: false},
		{name: "hcl invalid", x:"hcl", y: "123abc", result: false},
		{name: "ecl valid", x:"ecl", y: "brn", result: true},
		{name: "ecl invalid", x:"ecl", y: "wat", result: false},
		{name: "pid valid", x:"pid", y: "000000001", result: true},
		{name: "pid invalid", x:"pid", y: "0123456789", result: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := validValues(test.x, test.y)
			if result != test.result {
				t.Errorf("Expected %v and real %v)", test.result, result)
			}
		})
	}
}


func TestInValidPassports(t *testing.T) {
		real := createPassport("sample2", true)
		expected := 0
		if real != expected {
			t.Errorf("Expected %v and real %v)", expected, real)
		}
}

func TestValidPassports(t *testing.T) {
		real := createPassport("sample3", true)
		expected := 4
		if real != expected {
			t.Errorf("Expected %v and real %v)", expected, real)
		}
}
