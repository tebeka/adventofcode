package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func load(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	var nums []int
	for {
		var n int
		_, err := fmt.Fscanf(file, "%d", &n)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func upCount(nums []int) int {
	c := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			c++
		}
	}
	return c
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func wins(nums []int) []int {
	var w []int
	for i := 0; i < len(nums)-1; i++ {
		w = append(w, sum(nums[i:i+3]))
	}
	return w
}

func main() {
	// const fileName = "1_small.txt"
	const fileName = "1.txt"
	nums, err := load(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(upCount(nums))
	fmt.Println(upCount(wins(nums)))
}
