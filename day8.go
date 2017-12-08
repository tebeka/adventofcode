package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var code = `
b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10
`

func main() {
	file, err := os.Open("input-8.txt")
	if err != nil {
		log.Fatal(err)
	}

	maxReg := 0
	regs := map[string]int{}
	//scanner := bufio.NewScanner(strings.NewReader(code))
	scanner := bufio.NewScanner(file)
	for lnum := 1; scanner.Scan(); lnum++ {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		fields := strings.Fields(scanner.Text())
		if len(fields) != 7 {
			log.Fatalf("bad line: %d: %s", lnum, scanner.Text())
		}
		reg, op, creg, cmp := fields[0], fields[1], fields[4], fields[5]
		val, err := strconv.Atoi(fields[2])
		if err != nil {
			log.Fatal(err)
		}
		cval, err := strconv.Atoi(fields[6])
		if err != nil {
			log.Fatal(err)
		}
		cregV := regs[creg]
		var cond bool
		switch cmp {
		case ">":
			cond = cregV > cval
		case ">=":
			cond = cregV >= cval
		case "<":
			cond = cregV < cval
		case "<=":
			cond = cregV <= cval
		case "==":
			cond = cregV == cval
		case "!=":
			cond = cregV != cval
		default:
			log.Fatalf("%d: unknown cmp - %s", lnum, cmp)
		}
		if !cond {
			continue
		}
		switch op {
		case "inc":
			regs[reg] += val
		case "dec":
			regs[reg] -= val
		default:
			log.Fatalf("%d: unknown op - %s", lnum, op)
		}

		if regs[reg] > maxReg {
			maxReg = regs[reg]
		}
	}

	max := 0
	for _, val := range regs {
		if val > max {
			max = val
		}
	}

	fmt.Println(max)
	fmt.Println(maxReg)
}
