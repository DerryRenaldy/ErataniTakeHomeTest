package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := make([]int, 0, len(arr))
	middle := make([]int, 0, len(arr))
	right := make([]int, 0, len(arr))

	for _, v := range arr {
		switch {
		case v < pivot:
			left = append(left, v)
		case v == pivot:
			middle = append(middle, v)
		case v > pivot:
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	result := append(left, middle...)
	result = append(result, right...)

	return result
}

func generateRandomNumbers(count, min, max int) []int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := make([]int, count)
	for i := range numbers {
		numbers[i] = rng.Intn(max-min+1) + min
	}
	return numbers
}

func main() {
	randNumbers := generateRandomNumbers(10, 1, 100)

	fmt.Println("=== TestCase_4: Sorting Algorithm ===")
	fmt.Println()
	fmt.Println("Original numbers:", randNumbers)

	sortedNumbers := quickSort(randNumbers)

	fmt.Println("Sorted numbers: ", sortedNumbers)
	fmt.Println()
	fmt.Println("Algorithm used: QuickSort")
}
