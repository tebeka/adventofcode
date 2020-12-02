package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var prev []int

	s := bufio.NewScanner(file)
scan:
	for s.Scan() {
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		for _, p := range prev {
			if val+p == 2020 {
				fmt.Println(val * p)
				break scan
			}
		}

		prev = append(prev, val)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
