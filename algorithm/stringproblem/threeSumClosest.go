package stringproblem

import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum > target {
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < target {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return res
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
