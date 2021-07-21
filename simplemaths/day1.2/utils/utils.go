package utils

//IsPalinDrome return true if given number is palindrome else return false
func IsPalinDrome(input int) (isPalinDrome bool) {
	isPalinDrome = input == reverseNumber(input)
	return
}

//reverseNumber return the number formed by reversing the order of digits of input
func reverseNumber(input int) (reverseNum int) {
	for input != 0 {
		reverseNum = reverseNum*10 + input%10
		input = input / 10
	}

	return
}
