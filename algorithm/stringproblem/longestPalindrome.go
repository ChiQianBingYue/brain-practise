package stringproblem

func longestPalindrome(s string) int {
	freqMap := make(map[byte]uint32)
	for i := 0; i < len(s); i++ {
		freqMap[s[i]]++
	}
	count := 0
	for _, freq := range freqMap {
		// 将所有偶数次计入count
		count += int(freq &^ 1)
	}
	if count < len(s) {
		return count + 1
	}
	return count

}

func longestPalindromeSubStr(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len := max(
			expandAroundCenter(s, i, i),
			expandAroundCenter(s, i, i+1),
		)
		if len > (end - start + 1) {
			start = i - (len-1)/2
			end = i - len/2
		}
	}
	return s[start : end+1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

func longestPalindromeSubStr2(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		len := max2(expend(s, i, i), expend(s, i, i+1))
		if len > (right - left + 1) {
			left = i - (len-1)/2
			right = i + len/2
		}
	}
	return s[left : right+1]
}

func max2(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func expend(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
