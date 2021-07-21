package main

import (
	"IamMaheshDere/datastructure/simplemaths/day1.2/utils"
	"fmt"
)

func main() {
	input := 12321
	fmt.Printf("number: %d is palindrome: %v\n", input, utils.IsPalinDrome(input))

	input = 123
	fmt.Printf("number: %d is palindrome: %v\n", input, utils.IsPalinDrome(input))

}
