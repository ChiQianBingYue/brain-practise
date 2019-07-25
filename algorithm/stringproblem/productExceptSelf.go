package stringproblem

func productExceptSelf(nums []int) []int {
	nl := len(nums)
	res := make([]int, nl, nl)
	p, q := 1, 1
	for i := 0; i < nl; i++ {
		res[i] = p
		p *= nums[i]
	}
	for i := 0; i < nl; i++ {
		res[nl-i-1] *= q
		q *= nums[nl-i-1]
	}
	return res
}
