// FIXME: Wrong answer
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
)

// Rect is rectangle
type Rect struct {
	//ID      int
	TopX    int
	TopY    int
	BottomX int
	BottomY int
}

func parseLine(line string) (*Rect, error) {
	fields := lineRe.FindStringSubmatch(line)
	if len(fields) != 6 {
		return nil, fmt.Errorf("wrong number of fields - %d", len(fields))
	}

	var err error
	r := &Rect{}
	/*
		r.ID, err = strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("bad id - %s", fields[1])
		}
	*/

	r.TopX, err = strconv.Atoi(fields[2])
	if err != nil {
		return nil, fmt.Errorf("bad top x - %s", fields[2])
	}

	r.TopY, err = strconv.Atoi(fields[3])
	if err != nil {
		return nil, fmt.Errorf("bad top y - %s", fields[3])
	}

	r.BottomX, err = strconv.Atoi(fields[4])
	if err != nil {
		return nil, fmt.Errorf("bad width - %s", fields[4])
	}
	r.BottomX += r.TopX

	r.BottomY, err = strconv.Atoi(fields[5])
	if err != nil {
		return nil, fmt.Errorf("bad height - %s", fields[5])
	}
	r.BottomY += r.TopY

	return r, nil
}

func parseFile(r io.Reader) ([]*Rect, error) {
	var rects []*Rect
	lnum := 0
	s := bufio.NewScanner(r)
	for s.Scan() {
		lnum++
		r, err := parseLine(s.Text())
		if err != nil {
			return nil, fmt.Errorf("%d: %s", lnum, err)
		}
		rects = append(rects, r)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return rects, nil
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

func inRect(r *Rect, x int, y int) bool {
	if x < r.TopX || x > r.BottomX {
		return false
	}

	if y < r.TopY || y > r.BottomY {
		return false
	}

	return true
}

func overlap(r1, r2 *Rect) int {
	// Max(0, Min(XA2, XB2) - Max(XA1, XB1))
	// *
	// Max(0, Min(YA2, YB2) - Max(YA1, YB1))
	width := max(0, min(r1.BottomX, r2.BottomX)-max(r1.TopX, r2.TopX))
	height := max(0, min(r1.BottomY, r2.BottomY)-max(r1.TopY, r2.TopY))
	return width * height
}

var data = `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`

func main() {
	file, err := os.Open("day-3.txt")
	if err != nil {
		panic(err)
	}

	//file := strings.NewReader(data)

	rects, err := parseFile(file)
	if err != nil {
		panic(err)
	}

	n := 0
	for i, r1 := range rects {
		for _, r2 := range rects[i+1:] {
			n += overlap(r1, r2)
		}
	}
	fmt.Println(n)
}
