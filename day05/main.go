package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func getRow(input string) int {
	min := 0
	max := 127
	for _, r := range input {
    min, max = getRange(min, max, string(r))
	}
	return min
}

func getRange(min int, max int, kind string) (int, int) {
	var newMin, newMax int
	interval := ((max - min) + 1) / 2
	if kind == "L" || kind == "F" {
		newMax = max - interval
		newMin = min
	}
	if kind == "R" || kind == "B" {
		newMin = min + interval
		newMax = max
	}
	return newMin, newMax
}

func getColumn(input string) int {
	min := 0
	max := 7
	for _, r := range input {
    min, max = getRange(min, max, string(r))
	}
	return min
}

func Boarding(input string) int {
	row := getRow(input[0:7])
	column := getColumn(input[7:10])
	// fmt.Printf("Row: %d Column: %d\n", row, column)
	return row * 8 + column
}

func main() {
	file, err := os.Open("input")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []int
	var out int
	for scanner.Scan() {
			s := scanner.Text()
			out = Boarding(s)
			data = append(data, out)
	}
	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
	sort.Ints(data)
	fmt.Println("Part 1")
	fmt.Println(data[len(data)-1])
	fmt.Println("Part 2")
	for i, v := range data {
		if i+23 != v {
			fmt.Println(v-1)
			break;
		}
	}

}
