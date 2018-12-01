package main

import (
	"fmt"
	"strings"
)

func mem2key(mem []int) string {
	var smem []string
	for _, val := range mem {
		smem = append(smem, fmt.Sprintf("%d", val))
	}

	return strings.Join(smem, ",")

}

func maxBank(mem []int) (int, int) {
	imax, vmax := 0, mem[0]
	for i, v := range mem[1:] {
		if v > vmax {
			vmax = v
			imax = i + 1
		}
	}
	return imax, vmax
}

func step1() {
	//	mem := []int{0, 2, 7, 0}
	mem := []int{14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4}
	seen := map[string]bool{
		mem2key(mem): true,
	}

	for count := 1; ; count++ {
		i, v := maxBank(mem)

		mem[i] = 0
		for i = (i + 1) % len(mem); v > 0; v, i = v-1, (i+1)%len(mem) {
			mem[i]++
		}
		key := mem2key(mem)
		if _, ok := seen[key]; ok {
			fmt.Println(count)
			break
		}
		seen[key] = true
	}
}

func step2() {
	//mem := []int{0, 2, 7, 0}
	mem := []int{14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4}
	seen := map[string]int{
		mem2key(mem): 1,
	}

	for {
		i, v := maxBank(mem)

		mem[i] = 0
		for i = (i + 1) % len(mem); v > 0; v, i = v-1, (i+1)%len(mem) {
			mem[i]++
		}
		key := mem2key(mem)
		if _, ok := seen[key]; ok {
			fmt.Println(seen[key])
			break
		}
		for k := range seen {
			seen[k]++
		}
		seen[key] = 1
	}
}

func main() {
	step1()
	step2()
}
