package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var input int
	fmt.Print("Input: ")
	fmt.Scan(&input)

	result := factorial(input)
	fmt.Println("Output:", result)
}
