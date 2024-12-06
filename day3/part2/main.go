package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseProcedures(values []string, add *bool) [][]int {
	r := [][]int{}
	doCount := 0
	dontCount := 0
	for _, v := range values {
		if v == "don't()" {
			*add = false
			dontCount++
			continue
		} else if v == "do()" {
			*add = true
			doCount++
			continue
		}

		if *add {
			re := regexp.MustCompile(`\d{1,3}`)
			matches := re.FindAllString(v, -1)
			numbers := []int{}
			for _, match := range matches {
				num, _ := strconv.Atoi(match)
				numbers = append(numbers, num)
			}
			r = append(r, numbers)
		}
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
	add := true
	for scanner.Scan() {
		re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
		matches := re.FindAllString(scanner.Text(), -1)

		r := parseProcedures(matches, &add)

		n += calculateValue(r)
	}
	fmt.Println(n)
}
