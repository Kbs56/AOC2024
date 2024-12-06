package main

import (
	"fmt"
	"os"
	"strings"
)

func hasXmas(i, j int, lines []string) bool {
	n := len(lines)
	m := len(lines[0])

	if !(1 <= i && i < n-1 && 1 <= j && j < m-1) {
		return false
	}

	if lines[i][j] != 'A' {
		return false
	}

	diag1 := string([]byte{lines[i-1][j-1], lines[i+1][j+1]})
	diag2 := string([]byte{lines[i-1][j+1], lines[i+1][j-1]})

	return (diag1 == "MS" || diag1 == "SM") && (diag2 == "MS" || diag2 == "SM")
}

func main() {
	content, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	fileContent := strings.TrimSpace(string(content))
	lines := strings.Split(fileContent, "\n")

	n := len(lines)
	m := len(lines[0])

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if hasXmas(i, j, lines) {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
