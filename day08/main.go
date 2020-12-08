package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseLine(input string) (string, int) {
	if len(input)>1 {
		if i, err := strconv.Atoi(input[4:]); err == nil {
			return input[:3], i
		} else {
			fmt.Println(err)
		}
	}
	return "",0
}

func readLines(input string) map[int]map[string]int {
	output := make(map[int]map[string]int)
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	for i, line := range lines {
		op, num := parseLine(line)
		// fmt.Printf("Index: %d Op: %s Num: %d\n", i, op, num)
		l:= make(map[string]int)
		l[op] = num
		output[i] = l
	}
	return output
}

func readFile(filename string) string {
	contentInBytes, err := ioutil.ReadFile(filename)
	if err != nil {
			log.Fatal(err)
	}
	return string(contentInBytes)
}

func getAcc(instructions map[int]map[string]int) int {
	acc := 0
	i := 0
	loopDetected := false
	jmps := []int{}
	for loopDetected != true {
		for key, value := range instructions[i] {
			jmps = append(jmps, i)
			if key == "acc" {
					acc += value
					// fmt.Println("acc:", acc)
			}
			if key == "jmp" {
					i += value
					// fmt.Println("jmp(i):", value, i)
			} else {
				i++
			}
		}
		for _, j := range jmps {
			// fmt.Println(jmps)
			if i == j {
				loopDetected = true
			}
		}
	}
	return acc
}

func main() {
	sample := readFile("input")
	instructions := readLines(sample)
	real := getAcc(instructions)
	fmt.Println(real)
}
