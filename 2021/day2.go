package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	loc, depth := 0, 0
	aim := 0
	/* part 1
	d := map[string]func(int){
		"forward": func(n int) { loc += n },
		"up":      func(n int) { depth -= n },
		"down":    func(n int) { depth += n },
	}
	*/
	d := map[string]func(int){
		"forward": func(n int) {
			loc += n
			depth += aim * n
		},
		"up":   func(n int) { aim -= n },
		"down": func(n int) { aim += n },
	}

	// file, err := os.Open("2_small.txt")
	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for {
		var cmd string
		var n int
		_, err := fmt.Fscanf(file, "%s %d", &cmd, &n)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		d[cmd](n)
	}
	fmt.Println(loc * depth)
}
