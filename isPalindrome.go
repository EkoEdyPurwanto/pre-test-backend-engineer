package main

import (
	"fmt"
	"strings"
)

func isPalindrome(word string) bool {
	word = strings.ToLower(strings.ReplaceAll(word, " ", ""))
	length := len(word)
	for i := 0; i < length/2; i++ {
		if word[i] != word[length-i-1] {
			return false
		}
	}
	return true
}
func main() {
	var (
		input1 string
		input2 string
	)
	fmt.Print("input word1 here:")
	fmt.Scanf("%s", &input1)
	fmt.Println(isPalindrome(input1))

	fmt.Print("input word2 here:")
	fmt.Scanf("%s\n", &input2)
	fmt.Println(isPalindrome(input2))

}
