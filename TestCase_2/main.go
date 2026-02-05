package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)

	if len(s) <= 1 {
		return true
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	input := bufio.NewReader(os.Stdin)

	fmt.Println("=== PALINDROME CHECKER ===")
	fmt.Println("Masukkan string untuk dicek:")
	fmt.Print("> ")

	var text string
	_, err := fmt.Fscan(input, &text)
	if err != nil {
		fmt.Println("Error: invalid input")
		os.Exit(1)
	}

	result := IsPalindrome(text)
	if result {
		fmt.Printf("'%s' adalah palindrome\n", text)
	} else {
		fmt.Printf("'%s' bukan palindrome\n", text)
	}
}
