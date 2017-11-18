package palindrome_number

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Log(isPalindrome(0))
	t.Log(isPalindrome(1))
	t.Log(isPalindrome(11))
	t.Log(isPalindrome(121))
	t.Log(isPalindrome(12))
	t.Log(isPalindrome(122))
	t.Log(isPalindrome(-2147447412))
	t.Log(isPalindrome(2147447412))
}
