package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	sum := 0
	xx := x

	count := 10
	for xx != 0 {
		sum = sum*count + xx%10
		xx = xx / 10
	}

	if sum != x {
		return false
	}

	return true
}
