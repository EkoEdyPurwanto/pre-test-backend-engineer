package main

import (
	"fmt"
)

func findMax(arr []int) int {
	type s string
	if len(arr) < 0 {
		return 0
	}
	maximum := arr[0]
	for _, num := range arr {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}

func main() {
	var numCount int
	fmt.Print("Enter total of elements: ")
	_, err := fmt.Scan(&numCount)
	if err != nil {
		return
	}

	input := make([]int, numCount)
	for i := 0; i < numCount; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		_, err := fmt.Scan(&input[i])
		if err != nil {
			return
		}
	}

	fmt.Println("Input:", input)
	fmt.Println("Output:", findMax(input))
}
