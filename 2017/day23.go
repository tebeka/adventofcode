package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type inst struct {
	op   string
	dest string
	val  string
}

var isReg = regexp.MustCompile("^[a-h]$").MatchString

func valOf(val string, regs map[string]int) (int, error) {
	if isReg(val) {
		return regs[val], nil
	}

	return strconv.Atoi(val)
}

func run(prog []inst) (int, error) {
	pc := 0
	regs := map[string]int{}
	nmul := 0

	for pc >= 0 && pc < len(prog) {
		inst := prog[pc]
		offset := 1
		val, err := valOf(inst.val, regs)
		if err != nil {
			return 0, fmt.Errorf("%d: bad inst %v", pc, inst)
		}
		switch inst.op {
		case "set":
			regs[inst.dest] = val
		case "sub":
			regs[inst.dest] -= val
		case "mul":
			regs[inst.dest] *= val
			nmul++
		case "jnz":
			condVal, err := valOf(inst.dest, regs)
			if err != nil {
				return 0, fmt.Errorf("%d: bad inst %v", pc, inst)
			}
			if condVal != 0 {
				offset = val
			}
		}
		pc += offset
	}

	return nmul, nil
}

func parse(in io.Reader) ([]inst, error) {
	scan := bufio.NewScanner(in)
	lnum := 0
	var prog []inst
	for scan.Scan() {
		lnum++
		fields := strings.Fields(strings.TrimSpace(scan.Text()))
		if len(fields) != 3 {
			return nil, fmt.Errorf("%d: bad line - %s", lnum, scan.Text())
		}
		prog = append(prog, inst{fields[0], fields[1], fields[2]})
	}

	return prog, scan.Err()
}

func main() {
	file, err := os.Open("input-23.txt")
	if err != nil {
		panic(err)
	}

	prog, err := parse(file)
	if err != nil {
		panic(err)
	}

	val, err := run(prog)
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
