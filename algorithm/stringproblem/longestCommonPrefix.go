package stringproblem

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minL := len(strs[0])
	for _, str := range strs {
		l := len(str)
		if l < minL {
			minL = l
		}
	}
	if minL == 0 {
		return ""
	}
	i := 0
	for ; i < minL; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0][0:i]
}
