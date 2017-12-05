package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func show(i int, steps []int) {
	for j, v := range steps {
		if j == i {
			fmt.Printf("(%d) ", v)
		} else {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println("")
}

func execute(steps []int) int {
	// Make a copy
	orig := steps
	steps = make([]int, len(steps))
	copy(steps, orig)

	count := 0
	i := 0

	for {
		count++
		// show(i, steps)
		val := steps[i]
		steps[i]++
		i += val
		if i < 0 || i >= len(steps) {
			return count
		}
	}
}

func readInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var steps []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		steps = append(steps, i)
	}

	return steps, scanner.Err()
}

func main() {
	steps, err := readInput("input-5.txt")
	if err != nil {
		log.Fatal(err)
	}
	// steps := []int{0, 3, 0, 1, -3}

	n := execute(steps)
	fmt.Println(n)
}
