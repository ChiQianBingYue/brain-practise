package mathandnumber

import "math"

func reverse(x int) int {
	if x == 0 || x >= math.MaxInt32 || x <= math.MinInt32 {
		return 0
	}
	res := 0
	for x != 0 {
		res = res*10 + x%10
		if res >= math.MaxInt32 || res <= math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return res
}
