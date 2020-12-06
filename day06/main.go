package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func SortString(w string) string {
		s := strings.Split(w, "")
		sort.Strings(s)
    return strings.Join(s, "")
}


func HowManyCharsPart2(input string) int {
	input = strings.TrimSpace(input)
	groups := strings.Split(input, "\n")
	input = strings.Replace(input, "\n", " ", -1)
	input = strings.Replace(input, " ", "", -1)
	input = SortString(input)
	if len(groups) == 1 {
		return len(input)
	}
	count := 0
	groupcount := 0
	for i, val := range input {
		count = 0
		for j, second := range input {
			if (j>=i) {
				// fmt.Printf("(%d,%d) %s - % s\n", i, j, string(val), string(second))
				if string(val) == string(second) {
					count++
				}
			}
		}
		if count == len(groups) {
			groupcount++
		}
	}
	return groupcount
}

func calculateUniqChars(input string) int {
	input = strings.Replace(input, " ", "", -1)
	input = SortString(input)
	prev := ""
	count := 0
	for _, val := range input {
		if prev != string(val) {
			count++
		}
		prev = string(val)
	}
	return count
}

func HowManyCharsPart1(input string) int {
	input = strings.Replace(input, "\n", " ", -1)
	return calculateUniqChars(input)
}

func loadFile(filename string) (int, int) {
	contentInBytes, err := ioutil.ReadFile(filename)
	if err != nil {
			log.Fatal(err)
	}
	count := 0
	count2 := 0
  content := string(contentInBytes)
	groups := strings.Split(content, "\n\n")
	for _, group := range groups {
		v := HowManyCharsPart1(group)
		count += v
		v2 := HowManyCharsPart2(group)
		count2 += v2

	}
	return count, count2
}

func main() {
	out1, out2 := loadFile("input")
	fmt.Println(out1)
	fmt.Println(out2)
}
