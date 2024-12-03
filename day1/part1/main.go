package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	file, _ := os.Open("../input.txt")

	scanner := bufio.NewScanner(file)

	arr1 := []int{}
	arr2 := []int{}

	for scanner.Scan() {
		var x, y int
		fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		arr1 = append(arr1, x)
		arr2 = append(arr2, y)
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	res := 0
	for i := 0; i < len(arr1); i++ {
		res += int(math.Abs(float64(arr1[i] - arr2[i])))
	}

	fmt.Println(res)
}
