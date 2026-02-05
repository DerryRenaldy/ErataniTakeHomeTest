package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func isAccepted(n int) bool {
	if n%3 == 0 {
		return false
	}

	if n%10 == 3 {
		return false
	}

	return true
}

func FindAcceptedValue(k int) int {
	if k == 0 {
		return 0
	}

	count := 0
	num := 0

	for count < k {
		num++
		if isAccepted(num) {
			count++
		}
	}

	return num
}

func validateTestCaseCount(n int) error {
	if n < 1 || n > 100 {
		return fmt.Errorf("Error: number of test cases must be between 1 and 100, got %d", n)
	}
	return nil
}

func validateTestCaseValue(v int) error {
	if v < 1 || v > 1000 {
		return fmt.Errorf("Error: test case value must be between 1 and 1000, got %d", v)
	}
	return nil
}

type Result struct {
	idx   int
	value int
}

func main() {
	input := bufio.NewReader(os.Stdin)

	var numberOfTestCases int
	var testCases []int

	fmt.Println("Enter the number of test cases:")

	_, err := fmt.Fscan(input, &numberOfTestCases)
	if err != nil {
		fmt.Println("Error: invalid input, please enter a number")
		os.Exit(1)
	}

	if err := validateTestCaseCount(numberOfTestCases); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("=== INPUTS ===")

	for i := 0; i < numberOfTestCases; i++ {
		var inputNumber int
		fmt.Printf("Test case %d: ", i+1)
		_, err := fmt.Fscan(input, &inputNumber)
		if err != nil {
			fmt.Printf("Error: invalid input for test case %d, please enter a number\n", i+1)
			os.Exit(1)
		}

		if err := validateTestCaseValue(inputNumber); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		testCases = append(testCases, inputNumber)
	}

	fmt.Println("\n=== ANSWERS ===")

	results := make(chan Result, numberOfTestCases)
	var wg sync.WaitGroup

	// Create a goroutine for each test case
	for idx, value := range testCases {
		wg.Add(1)
		go func(idx, val int) {
			defer wg.Done()

			result := FindAcceptedValue(val)
			results <- Result{idx, result}
		}(idx, value)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	output := make([]Result, numberOfTestCases)
	for r := range results {
		output[r.idx] = r
	}

	for i := 0; i < numberOfTestCases; i++ {
		fmt.Printf("Test case %d: %d\n", i+1, output[i].value)
	}
}
