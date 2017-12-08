package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func step1(puzzle string) {
	re, err := regexp.Compile("[a-z]+")
	if err != nil {
		log.Fatal(err)
	}

	counts := map[string]int{}
	for _, word := range re.FindAllString(puzzle, -1) {
		counts[word]++
	}

	for word, count := range counts {
		if count == 1 {
			fmt.Println(word)
			break
		}
	}
}

func step2(puzzle string) {
	re, err := regexp.Compile(`([a-z]+) \(([0-9]+)\)( -> (.*))?`)
	if err != nil {
		log.Fatal(err)
	}

	weights := map[string]int{}
	tops := map[string][]string{}

	scanner := bufio.NewScanner(strings.NewReader(puzzle))

	for scanner.Scan() {
		groups := re.FindStringSubmatch(scanner.Text())
		name := groups[1]
		weight, err := strconv.Atoi(groups[2])
		if err != nil {
			log.Fatal(err)
		}
		weights[name] = weight
		if groups[4] != "" {
			tops[name] = strings.Split(groups[4], ", ")
		} else {
			tops[name] = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var diskWeight func(string) int

	diskWeight = func(name string) int {
		weight := weights[name]
		for _, top := range tops[name] {
			weight += diskWeight(top)
		}
		return weight
	}

	for prog, toppers := range tops {
		if len(toppers) == 0 {
			continue
		}

		var ws []int
		for _, name := range toppers {
			ws = append(ws, diskWeight(name))
		}
		sort.Ints(ws)
		if ws[0] == ws[len(ws)-1] {
			continue
		}

		diff := ws[0] - ws[len(ws)-1]
		fmt.Println(prog, weights[prog]-diff)
	}

}

func main() {
	data, err := ioutil.ReadFile("input-7.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle := string(data)
	/*
			puzzle = `pbga (66)
		xhth (57)
		ebii (61)
		havc (66)
		ktlj (57)
		fwft (72) -> ktlj, cntj, xhth
		qoyq (66)
		padx (45) -> pbga, havc, qoyq
		tknk (41) -> ugml, padx, fwft
		jptl (61)
		ugml (68) -> gyxo, ebii, jptl
		gyxo (61)
		cntj (57)
		`
	*/

	step1(puzzle)
	step2(puzzle)
}
