package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseProcedures(values []string) [][]int {
	r := [][]int{}
	re := regexp.MustCompile(`\d{1,3}`)
	for _, v := range values {
		matches := re.FindAllString(v, -1)
		numbers := []int{}
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			numbers = append(numbers, num)
		}
		r = append(r, numbers)
	}
	return r
}

func calculateValue(values [][]int) int {
	r := 0
	for i := range values {
		r += (values[i][0] * values[i][1])
	}
	return r
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	n := 0
	for scanner.Scan() {
		re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		matches := re.FindAllString(scanner.Text(), -1)

		r := parseProcedures(matches)

		n += calculateValue(r)
	}
	fmt.Println(n)
}
