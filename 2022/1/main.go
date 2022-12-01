package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	max, curr := 0, 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := strings.TrimSpace(s.Text())
		if v == "" {
			if curr > max {
				max = curr
			}
			curr = 0
			continue
		}

		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		curr += n
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	if curr > max {
		max = curr
	}
	fmt.Println(max)

}

func main() {
	var values []int
	var curr int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := strings.TrimSpace(s.Text())
		if v == "" {
			values = append(values, curr)
			curr = 0
			continue
		}

		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		curr += n
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	values = append(values, curr)
	sort.Ints(values)
	sum := 0
	for _, v := range values[len(values)-3:] {
		sum += v
	}
	fmt.Println(sum)
}
