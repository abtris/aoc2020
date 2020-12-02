package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// SumAndMultiply ...
func SumAndMultiply(numbers []int) int {
	sort.Ints(numbers)
	for _, val := range numbers {
		for _, compareVal := range numbers {
			if val + compareVal == 2020 {
				return val * compareVal
			}
		}
	}
	return 0
}
// SumAndMultiplySecond ...
func SumAndMultiplySecond(numbers []int) int {
	sort.Ints(numbers)
	for _, val := range numbers {
		for _, compareVal := range numbers {
			for _, compareVal2 := range numbers {
				if val + compareVal +  compareVal2 == 2020 {
					return val * compareVal * compareVal2
				}
			}
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []int
	for scanner.Scan() {
			s := scanner.Text()
			if i, err := strconv.Atoi(s); err == nil {
				data = append(data, i)
			}
	}
	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
	fmt.Println("Part 1")
	fmt.Println(SumAndMultiply(data))
	fmt.Println("Part 2")
	fmt.Println(SumAndMultiplySecond(data))
}
