package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	//"strings"
)

var (
	// #1 @ 861,330: 20x10
	lineRe = regexp.MustCompile("([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	data   = `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`
)

// Square on fabric
type Square struct {
	X int
	Y int
}

// Claim is a claim on grid
type Claim struct {
	ID          int
	TopLeft     Square
	BottomRight Square
}

func parseLine(line string) (*Claim, error) {
	fields := lineRe.FindStringSubmatch(line)
	if len(fields) != 6 {
		return nil, fmt.Errorf("wrong number of fields - %d", len(fields))
	}

	id, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("bad id - %s", fields[1])
	}

	topX, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, fmt.Errorf("bad top x - %s", fields[2])
	}

	topY, err := strconv.Atoi(fields[3])
	if err != nil {
		return nil, fmt.Errorf("bad top y - %s", fields[3])
	}

	width, err := strconv.Atoi(fields[4])
	if err != nil {
		return nil, fmt.Errorf("bad width - %s", fields[4])
	}

	height, err := strconv.Atoi(fields[5])
	if err != nil {
		return nil, fmt.Errorf("bad height - %s", fields[5])
	}

	claim := &Claim{
		ID:          id,
		TopLeft:     Square{topX, topY},
		BottomRight: Square{topX + width - 1, topY + height - 1},
	}

	return claim, nil
}

func parseFile(r io.Reader) ([]*Claim, error) {
	var claims []*Claim
	lnum := 0
	s := bufio.NewScanner(r)
	for s.Scan() {
		lnum++
		c, err := parseLine(s.Text())
		if err != nil {
			return nil, fmt.Errorf("%d: %s", lnum, err)
		}
		claims = append(claims, c)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return claims, nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func overlap(c1, c2 *Claim) []Square {
	minX := max(c1.TopLeft.X, c2.TopLeft.X)
	maxX := min(c1.BottomRight.X, c2.BottomRight.X)
	minY := max(c1.TopLeft.Y, c2.TopLeft.Y)
	maxY := min(c1.BottomRight.Y, c2.BottomRight.Y)

	var sqs []Square
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			sqs = append(sqs, Square{x, y})
		}
	}

	return sqs
}

func part1(claims []*Claim) {
	dups := make(map[Square]bool)
	for i, c1 := range claims {
		for _, c2 := range claims[i+1:] {
			for _, s := range overlap(c1, c2) {
				dups[s] = true
			}
		}
	}

	fmt.Println(len(dups))
}

func hasOverlap(c *Claim, claims []*Claim) bool {
	for _, other := range claims {
		if other.ID == c.ID {
			continue
		}
		if len(overlap(other, c)) > 0 {
			return true
		}
	}
	return false
}

func part2(claims []*Claim) {
	for _, c := range claims {
		if !hasOverlap(c, claims) {
			fmt.Println(c.ID)
			return
		}
	}
}

func main() {
	file, err := os.Open("day-3.txt")
	if err != nil {
		panic(err)
	}
	// file := strings.NewReader(data)
	claims, err := parseFile(file)
	if err != nil {
		panic(err)
	}

	part1(claims)
	part2(claims)
}
