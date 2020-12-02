package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// VerifyPassword ...
func VerifyPassword(min int, max int, test string, password string) bool {
	calculated := strings.Count(password, test)
	if calculated == 0 {
		return false
	}
	if min <= calculated {
		if max >= calculated {
			return true
		}
	}
	return false
}


// VerifyPasswordNewPolicy ...
func VerifyPasswordNewPolicy(min int, max int, test string, password string) bool {
	valid := min - 1
	invalid := max - 1
	if (string(password[valid]) == test || string(password[invalid]) == test) {
		if (string(password[valid]) != string(password[invalid])) {
			return true
		}
	}
	return false
}


// ElfIssue ...
type ElfIssue struct {
	min 				int
	max					int
	testValue		string
	password 		string
}

func main() {
	file, err := os.Open("input")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var data []ElfIssue

	for scanner.Scan() {
			s := scanner.Text()
			zp := regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<test>.): (?P<value>.+)`)
			output := zp.FindStringSubmatch(s)
			min := 0
			max := 0
			if i, err := strconv.Atoi(output[1]); err == nil {
				min = i
			}
			if i, err := strconv.Atoi(output[2]); err == nil {
				max = i
			}
			elf := ElfIssue{
				min: min,
				max: max,
				testValue: output[3],
				password: output[4],
			}
			data = append(data, elf)
	}
	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
	valid := 0
	for _, elf := range data {
		v := VerifyPassword(elf.min, elf.max, elf.testValue, elf.password)
		if v == true {
			valid++
		}
	}
	fmt.Println("Part 1")
	fmt.Println(valid)
	valid2 := 0
	for _, elf := range data {
		v := VerifyPasswordNewPolicy(elf.min, elf.max, elf.testValue, elf.password)
		if v == true {
			valid2++
		}
	}
	fmt.Println("Part 2")
	fmt.Println(valid2)
}
