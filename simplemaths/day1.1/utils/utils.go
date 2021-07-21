package utils

//countDigits count number of digits in the input
func CountDigits(input int) (digits int) {
	for input != 0 {
		digits = digits + 1
		input = input / 10
	}

	return
}
