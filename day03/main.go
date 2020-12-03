package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
type slope struct {
	right int
	down int
}


type forest struct {
	places []string
}

func traverseForest(f forest, s slope) int {
	hits := 0
	lineLength := len(f.places[0])
	x := 0
	down := 0
	for row, line := range f.places {
		// 35 #
		// 46 .
		if row == down {
			if line[x] == 35 {
				hits++
			}
			down += s.down
			x += s.right
		}
		if x >= lineLength {
			x = x-lineLength
		}
	}
	return hits
}

func createForest(filename string) forest {
	file, err := os.Open(filename)
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var myForest forest
	for scanner.Scan() {
			s := scanner.Text()
			myForest.places = append(myForest.places, s)
	}
	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
	return myForest
}

func main() {
	input := createForest("input")
	fmt.Println("Part 1")
	oneslope := slope{right: 3, down: 1}
	fmt.Println(traverseForest(input, oneslope))
	fmt.Println("Part 2")
	slopes := []slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}
	real := 1
	for _, slope := range slopes {
		count := traverseForest(input, slope)
		real *= count
	}
	fmt.Println(real)
}
