package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	file, _ := os.Open("../input.txt")

	scanner := bufio.NewScanner(file)

	arr1 := []int{}
	arr2 := []int{}
	numMap := make(map[int]int)

	for scanner.Scan() {
		var x, y int
		fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		arr1 = append(arr1, x)
		arr2 = append(arr2, y)
		numMap[y]++
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	res := 0
	for i := 0; i < len(arr1); i++ {
		n, ok := numMap[arr1[i]]
		if ok {
			m := arr1[i] * n
			res += m
		}
	}

	fmt.Println(res)
}
