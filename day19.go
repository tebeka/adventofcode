package main

import (
	"bufio"
	"fmt"
	"strings"
)

var data = `
     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 
`

func main() {
	scan := bufio.NewScanner(strings.NewReader(data))
	var letters []byte
	var maze [][]byte
	for scan.Scan() {
		line := scan.Text()
		if len(line) == 0 {
			continue
		}
		maze = append(maze, []byte(line))
	}

	dr, dc := 1, 0
	row, col := 0, 0
	// Find start
	for col = 0; col < len(maze[0]) && maze[0][col] == ' '; col++ {
	}

	fmt.Println(row, col)

	for {
		fmt.Println(row)
		if row < 0 || row >= len(maze) {
			break
		}
		if col < 0 || col >= len(maze[0]) {
			break
		}

		c := maze[row][col]
		switch {
		case c == '|' || c == '-':
			row, col = row+dr, col+dc
		case c >= 'A' && c <= 'Z':
			letters = append(letters, c)
			row, col = row+dr, col+dc
		case c == '+':
			// TODO: Find move
	}

}
