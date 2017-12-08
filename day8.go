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

func eval(regs map[string]int, reg, cmp, val string) bool {
	cval, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}

	switch cmp {
	case ">":
		return regs[reg] > cval
	case ">=":
		return regs[reg] >= cval
	case "<":
		return regs[reg] < cval
	case "<=":
		return regs[reg] <= cval
	case "==":
		return regs[reg] == cval
	case "!=":
		return regs[reg] != cval
	default:
		log.Fatalf("unknown cmp - %s", cmp)
	}

	return false // make compiler happy
}

func maxRegs(regs map[string]int) int {
	max := 0
	for _, val := range regs {
		if val > max {
			max = val
		}
	}
	return max
}

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

		if !eval(regs, fields[4], fields[5], fields[6]) {
			continue
		}

		val, err := strconv.Atoi(fields[2])
		if err != nil {
			log.Fatal(err)
		}

		reg, op := fields[0], fields[1]

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

	fmt.Println(maxRegs(regs))
	fmt.Println(maxReg)
}
