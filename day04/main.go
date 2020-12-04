package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// InArray function test if string is in array of string
func InArray(val string, array []string) bool {
	r, err := regexp.Compile(`(?i)` + regexp.QuoteMeta(val))
	if err != nil {
		fmt.Printf("InArray: Error in compile regexp for value %v", val)
		return false
	}
	for _, v := range array {
		match := r.MatchString(v)
		if match {
			return true
		}
	}
	return false
}

func validValues(token string, value string) bool {
	switch token {
		case "byr":
			if i, err := strconv.Atoi(value); err == nil {
				if i >= 1920 && i <= 2002 {
					return true
				} else {
					fmt.Println("Invalid byr (1920-2002)", value)
				}
			}
		case "iyr":
			if i, err := strconv.Atoi(value); err == nil {
				if i >= 2010 && i <= 2020 {
					return true
				} else {
					fmt.Println("Invalid iyr (2010-2020)", value)
				}
			}
		case "eyr":
			if i, err := strconv.Atoi(value); err == nil {
				if i >= 2020 && i <= 2030 {
					return true
				} else {
					fmt.Println("Invalid eyr (2020-2030)", value)
				}
			}
		case "hgt": // cm (150-193) or in (59-76)
			var ret = regexp.MustCompile(`^(?P<num>\d+)(?P<type>in|cm)$`)
			if ret.MatchString(value) {
				output := ret.FindStringSubmatch(value)
				height := 0
				if i, err := strconv.Atoi(output[1]); err == nil {
					height = i
				}
				if output[2] == "cm" {
					if height >= 150 && height <= 193 {
						return true
					}
				}
				if output[2] == "in" {
					if height >= 59 && height <= 76 {
						return true
					}
				}
			}
		case "hcl":
			var re = regexp.MustCompile(`^#[0-9a-f]{6}$`)
			if re.MatchString(value) {
					return true
			} else {
					fmt.Println("Invalid hcl", value)
			}
		case "ecl":
			validColors := []string{"amb","blu","brn","gry","grn","hzl","oth"}
			if len(value)>0 && InArray(value, validColors) {
				return true
			} else {
					fmt.Println("Invalid ecl", value)
			}
		case "pid":
			var reg = regexp.MustCompile(`^[0-9]{9}$`)
			if reg.MatchString(value) {
					return true
			} else {
					fmt.Println("Invalid pid", value)
			}
	}
	return false
}

func validPassport(p string, extraValidation bool) bool {
	validTokens := []string{"byr","iyr","eyr","hgt","hcl","ecl","pid"}
	tokens := strings.Split(p, " ")
	var headers []string
	var validationTokens []string
	for _, token := range tokens {
		headers = strings.Split(token, ":")
		if len(headers[0])>0 && InArray(headers[0], validTokens) {
			if extraValidation {
				if validValues(headers[0], headers[1]) {
					validationTokens = append(validationTokens, headers[0])
				}
			} else {
				validationTokens = append(validationTokens, headers[0])
			}
		}
	}
	sort.Strings(validTokens)
	sort.Strings(validationTokens)
	size := len(validationTokens)
	output := false
	if reflect.DeepEqual(validationTokens, validTokens) {
		output = true
	}
	if output == false && size == 6 {
		fmt.Printf("%s - %d - %v\n", p, size, output)
	}
	return output
}

func createPassport(filename string, extraValidation bool) int {
  contentInBytes, err := ioutil.ReadFile(filename)
	if err != nil {
			log.Fatal(err)
	}
	valid := 0
  content := string(contentInBytes)
	passports := strings.Split(content, "\n\n")
	for _, passport := range passports {
		passport = strings.Replace(passport, "\n", " ", -1)
		if len(passport) > 1 {
			v := validPassport(passport, extraValidation)
			if v {
				valid++
			}
		}
	}
	return valid
}

func main() {
	real := createPassport("input", false)
	fmt.Println(real)
	real2 := createPassport("input", true)
	fmt.Println(real2)
}
