package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func makeIntSlice(arr []string) []int {
	res := []int{}
	for _, val := range arr {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		res = append(res, num)
	}
	return res
}

func evalReport(arr []int) bool {
	isIncreaseing := true
	isDecreasing := true
	safeDiff := true
	for i := 0; i < len(arr)-1; i++ {
		diff := int(math.Abs(float64(arr[i]) - float64(arr[i+1])))
		if arr[i] > arr[i+1] {
			isIncreaseing = false
		} else if arr[i] < arr[i+1] {
			isDecreasing = false
		}

		if !isIncreaseing && !isDecreasing {
			safeDiff = false
			break
		}

		if diff < 1 || diff > 3 {
			safeDiff = false
			break
		}

	}

	return safeDiff
}

func main() {
	file, _ := os.Open("../input.txt")

	scanner := bufio.NewScanner(file)

	safeReports := 0
	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Fields(line)
		res := makeIntSlice(report)
		safeDiff := evalReport(res)
		if safeDiff {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
