package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	// Step C must be finished before step A can begin.
	lineRe = regexp.MustCompile("Step ([A-Z]+).*step ([A-Z]+)")
)

type Graph map[string]map[string]bool

func parseInput(r io.Reader) (Graph, error) {
	graph := make(Graph)
	scan := bufio.NewScanner(r)

	lNum := 1
	for scan.Scan() {
		fields := lineRe.FindStringSubmatch(scan.Text())
		if len(fields) != 3 {
			return nil, fmt.Errorf("%d: bad line - %s", lNum, scan.Text())
		}
		src, dest := fields[1], fields[2]
		nodes, ok := graph[src]
		if !ok {
			graph[src] = map[string]bool{dest: true}
		} else {
			nodes[dest] = true
		}

		if _, ok = graph[dest]; !ok {
			graph[dest] = make(map[string]bool)
		}
	}

	if err := scan.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func graphRoots(graph Graph) []string {
	if len(graph) == 1 {
		for node := range graph {
			return []string{node}
		}
	}

	hasIncoming := make(map[string]bool)
	for _, dests := range graph {
		for node := range dests {
			hasIncoming[node] = true
		}
	}

	var roots []string
	for node := range graph {
		if !hasIncoming[node] {
			roots = append(roots, node)
		}
	}

	return roots
}

// A combination of toplogical and lexicographical sort
func topoLexiSort(graph Graph) []string {
	var nodes []string
	roots := graphRoots(graph)
	for len(roots) > 0 {
		sort.Strings(roots)
		node := roots[0]
		nodes = append(nodes, node)
		delete(graph, node)
		roots = graphRoots(graph)
	}

	return nodes
}

var data = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

func duration(step string, delta int) int {
	return int(step[0]-'A') + 1 + delta
}

func allZero(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func processTime(n int, nodes []string, delta int) int {
	if n > len(nodes) {
		n = len(nodes)
	}

	times := make([]int, n)
	jobs := make([]string, n)
	for i := 0; i < n; i++ {
		times[i] = duration(nodes[i], delta)
		jobs[i] = nodes[i]
	}

	time := 0
	for {
		if allZero(times) && n >= len(nodes) {
			return time
		}
		time++
		fmt.Printf("%02d: %v\n", time, jobs)
		for i, t := range times {
			if t > 0 {
				t--
			}
			if t == 0 && n < len(nodes) {
				fmt.Println(time, nodes[n])
				t = duration(nodes[n], delta)
				jobs[i] = nodes[n]
				n++
			}
			times[i] = t
		}
	}
}

// A combination of toplogical and lexicographical timing
func topoLexiTime(graph Graph, n int, delta int) int {
	time := 0
	for len(graph) > 0 {
		roots := graphRoots(graph)
		sort.Strings(roots)
		time += processTime(n, roots, delta)
		for _, node := range roots {
			delete(graph, node)
		}
	}

	return time
}

func openInput(debug bool) (io.Reader, error) {
	if debug {
		return strings.NewReader(data), nil
	}

	return os.Open("day-7.txt")
}

func main() {
	debug := true
	n, delta := 2, 0
	if !debug {
		n, delta = 4, 60
	}

	r, err := openInput(debug)
	if err != nil {
		panic(err)
	}

	graph, err := parseInput(r)
	if err != nil {
		panic(err)
	}

	// part 1
	nodes := topoLexiSort(graph)
	fmt.Println(strings.Join(nodes, ""))

	// part 2 - FIXME
	r, err = openInput(debug)
	if err != nil {
		panic(err)
	}

	graph, err = parseInput(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(topoLexiTime(graph, n, delta))
}
