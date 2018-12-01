package main

import (
	"fmt"
)

const (
	div = 2147483647
)

func gen(initial, factor uint) chan uint {
	ch := make(chan uint)

	val := initial
	go func() {
		for {
			val *= factor
			val %= div
			ch <- val
		}
	}()

	return ch
}

func gen2(initial, factor, m uint) chan uint {
	ch := make(chan uint)

	val := initial
	go func() {
		for {
			val *= factor
			val %= div
			if val%m == 0 {
				ch <- val
			}
		}
	}()

	return ch
}

func main() {

	//aStart, bStart := uint(65), uint(8921)
	aStart, bStart := uint(618), uint(814)
	//n := 5
	/*
		genA := gen(aStart, 16807)
		genB := gen(bStart, 48271)
		n := 40000000
	*/
	genA := gen2(aStart, 16807, 4)
	genB := gen2(bStart, 48271, 8)
	n := 5000000
	count := 0

	for i := 0; i < n; i++ {
		a, b := <-genA, <-genB
		if a&0xFFFF == b&0xFFFF {
			count++
		}
	}
	fmt.Printf("COUNT: %d\n", count)
}
