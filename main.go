package main

import "fmt"

const HEIGHT, WIDTH = 6, 8

func main() {
	block := [][]string{
		{"#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", ".", ".", ".", ".", ".", ".", "#"},
		{"#", ".", "#", "#", "#", ".", ".", "#"},
		{"#", ".", ".", ".", "#", ".", "#", "#"},
		{"#", "X", "#", ".", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#"},
	}
	height, width := len(block), len(block[0])
	maxStep := height
	if height > width {
		maxStep = width
	}
	x, y := 0, 0
	for i, v := range block {
		for j, w := range v {
			if w == "X" {
				x = i
				y = j
			}
		}
	}

	counter := 0
	var i int
	for i = y; i <= maxStep; i++ {
		if i <= x && block[x][y+i] == "." {
			counter++
		}
	}

	fmt.Printf("kemungkinan titik %d\n", counter)
}
