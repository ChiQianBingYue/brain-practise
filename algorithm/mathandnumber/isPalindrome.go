package mathandnumber

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reverseRightHalf := 0
	for x > reverseRightHalf {
		reverseRightHalf = reverseRightHalf*10 + x%10
		x = x / 10

	}
	if reverseRightHalf == x || x == reverseRightHalf/10 {
		return true
	}
	return false
}
