package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

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

func processData(data map[string][]string, bag string) []string {
	bag = strings.TrimSpace(bag)
	// fmt.Printf("INPUT BAG: %+v\n", bag)
	// fmt.Printf("INPUT: %+v\n", data)
	results := []string{}
	for key, values := range data {
		bags := []string{}
		for _, value := range values {
			bags = append(bags, strings.TrimSpace(value)[2:])
		}
		// fmt.Printf("BAGS(%s): %+v\n", bag, bags)
		if !InArray(bag, bags) {
			continue
		}
		if InArray(bag, results) {
			results = append(results, processData(data, key)...)
		}
		results = append(results, bag)
	}
	// fmt.Printf("RESULTS: %+v\n", results)
	return results
}

func loadFile(filename string) map[string][]string {
	file, err := os.Open(filename)
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := map[string][]string{}
	for scanner.Scan() {
			s := scanner.Text()
			if !strings.Contains(s, "no other bag") {
				s = strings.ReplaceAll(s, "bags", "bag")
				s = strings.ReplaceAll(s, ".", "")
				out := strings.Split(s, "contain")
				left := out[0]
				right := out[1]
				data[left] = strings.Split(right, ",")
			}
	}
	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
	return data
}

func main() {
	data := loadFile("input")
	fmt.Println(len(data))
	real := processData(data, "shiny gold bag")
	fmt.Println(len(real))
}
