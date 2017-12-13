package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Layer struct {
	Depth int
	Range int
	Pos   int
	Step  int
}

type Firewall map[int]*Layer

func NewLayer(depth, rng int) *Layer {
	return &Layer{
		Depth: depth,
		Range: rng,
		Pos:   0,
		Step:  1,
	}
}

func (lr *Layer) Tick() {
	if lr.Pos == 0 {
		lr.Step = 1
	}
	if lr.Pos == lr.Range-1 {
		lr.Step = -1
	}
	lr.Pos += lr.Step
}

func (lr *Layer) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, " %02d ", lr.Depth)
	for i := 0; i < lr.Range; i++ {
		if i == lr.Pos {
			fmt.Fprintf(&buf, "[S]")
		} else {
			fmt.Fprintf(&buf, "[ ]")
		}
	}
	return buf.String()
}

func (ly *Layer) Reset() {
	ly.Pos = 0
	ly.Step = 0
}

var data = `
0: 3
1: 2
4: 4
6: 4
`

func parse(rdr io.Reader) (Firewall, error) {
	fw := make(Firewall)

	scanner := bufio.NewScanner(rdr)
	for lnum := 1; scanner.Scan(); lnum++ {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}

		fields := strings.Split(scanner.Text(), ": ")
		if len(fields) != 2 {
			return nil, fmt.Errorf("%d: bad line: %q", lnum, scanner.Text())
		}
		depth, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		rng, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		fw[depth] = NewLayer(depth, rng)
	}

	return fw, scanner.Err()
}

func (fw Firewall) Depth() int {
	depth := 0
	for _, layer := range fw {
		if layer.Depth > depth {
			depth = layer.Depth
		}
	}
	return depth
}

func (fw Firewall) String() string {
	var buf bytes.Buffer

	for i := 0; i <= fw.Depth(); i++ {
		layer, ok := fw[i]
		if ok {
			fmt.Fprintln(&buf, layer)
		} else {
			fmt.Fprintf(&buf, " %02d · · · · \n", i)
		}
	}

	return buf.String()
}

func (fw Firewall) Tick() {
	for _, layer := range fw {
		layer.Tick()
	}
}

func (fw Firewall) Reset() {
	for _, layer := range fw {
		layer.Reset()
	}
}

func step1(fw Firewall) int {
	severity := 0

	for d := 0; d <= fw.Depth(); d++ {
		layer, ok := fw[d]
		if ok && layer.Pos == 0 {
			severity += layer.Depth * layer.Range
		}
		fw.Tick()
	}

	return severity
}

func step2(fw Firewall) int {
	for i := 0; ; i++ {
		if i > 0 && i%1000 == 0 {
			fmt.Printf(">>> %d\n", i)
		}
		fw.Reset()
		for j := 0; j < i; j++ {
			fw.Tick()
		}
		if step1(fw) == 0 {
			return i
		}
	}
}

func main() {
	file, err := os.Open("input-13.txt")
	if err != nil {
		log.Fatal(err)
	}
	fw, err := parse(file)
	if err != nil {
		log.Fatal(err)
	}

	fw, err = parse(strings.NewReader(data))

	/*
		for i := 0; i < 5; i++ {
			fmt.Println(fw)
			fmt.Println("--------------")
			fw.Tick()
		}
	*/
	fmt.Println(step1(fw))
	// TODO: This is wrong
	fmt.Println(step2(fw))
}
