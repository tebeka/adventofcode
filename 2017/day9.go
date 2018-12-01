package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func skipGrabage(s string, i int) (int, int) {
	cancelled := 0
	start := i
	i++ // skip initial <
	for ; ; i++ {
		if s[i] == '>' {
			return i, i - start - cancelled - 1
		}
		if s[i] == '!' {
			i++
			cancelled += 2
		}
	}
}

func score(s string) (int, int) {
	stack := 0
	deleted := 0
	score := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '!':
			i++
		case '<':
			var d int
			i, d = skipGrabage(s, i)
			deleted += d
		case '{':
			stack++
		case '}':
			score += stack
			stack--
		}
	}

	if stack != 0 {
		log.Fatal("unbalanced stack")
	}

	return score, deleted
}

func testScore() {
	tests := []struct {
		s string
		c int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}

	for _, test := range tests {
		fmt.Println(test.s)
		if c, _ := score(test.s); c != test.c {
			log.Fatalf("%d != %d", c, test.c)
		}
	}
}

func testDeleted() {
	tests := []struct {
		s string
		c int
	}{
		{"<>", 0},
		{"<random characters>", 17},
		{"<<<<>", 3},
		{"<{!>}>", 2},
		{"<!!>", 0},
		{"<!!!>>", 0},
		{"<{o\"i!a,<{i<a>", 10},
	}

	for _, test := range tests {
		fmt.Println(test.s)
		if _, c := skipGrabage(test.s, 0); c != test.c {
			log.Fatalf("%d != %d", c, test.c)
		}
	}
}

func main() {
	data, err := ioutil.ReadFile("input-9.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(score(string(data)))

}
