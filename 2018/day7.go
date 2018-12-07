package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
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
	}

	if err := scan.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func graphRoots(graph Graph) []string {
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

func topoLexiSort(graph Graph) []string {
	var nodes []string
	roots := graphRoots(graph)
	for len(roots) > 0 {
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

func main() {
	r := strings.NewReader(data)
	graph, err := parseInput(r)
	if err != nil {
		panic(err)
	}

	for k, v := range graph {
		fmt.Printf("%s -> %v\n", k, v)
	}

	fmt.Printf("%v\n", graphRoots(graph))
}
