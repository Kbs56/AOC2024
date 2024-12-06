package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = []struct{ dr, dc int }{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{1, 1},
	{-1, 1},
	{1, -1},
}

var targetWord = "XMAS"

func checkWord(r, c int, dr, dc int, grid []string) bool {
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < len(targetWord); i++ {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}

		if grid[r][c] != targetWord[i] {
			return false
		}

		r += dr
		c += dc
	}
	return true
}

func searchGrid(grid []string) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if checkWord(r, c, dir.dr, dir.dc, grid) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	grid := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	xmas := searchGrid(grid)
	fmt.Println(xmas)
}
