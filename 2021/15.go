package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func loadMat(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	var mat [][]int
	for s.Scan() {
		var row []int
		for _, c := range s.Text() {
			row = append(row, int(c-'0'))
		}
		mat = append(mat, row)
	}

	return mat, nil
}

type point struct {
	x int
	y int
}

func shortestPath(mat [][]int, p point, cache map[point]int) int {
	if p.x == len(mat)-1 && p.y == len(mat[0])-1 {
		return mat[p.x][p.y]
	}

	if n, ok := cache[p]; ok {
		return n
	}

	px, py := -1, -1

	if p.x < len(mat)-1 {
		px = shortestPath(mat, point{p.x + 1, p.y}, cache)
	}

	if p.y < len(mat[0])-1 {
		py = shortestPath(mat, point{p.x, p.y + 1}, cache)
	}

	n := 0
	if p.x != 0 || p.y != 0 {
		n = mat[p.x][p.y]
	}

	switch {
	case px == -1:
		n += py
	case py == -1:
		n += px
	default:
		if px < py {
			n += px
		} else {
			n += py
		}
	}
	cache[p] = n
	return n
}

func main() {
	// fileName := "15_small.txt"
	fileName := "15.txt"

	mat, err := loadMat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	cache := make(map[point]int)
	n := shortestPath(mat, point{0, 0}, cache)
	fmt.Println(n)
}
