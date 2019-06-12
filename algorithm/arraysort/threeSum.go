package arraysort

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	if nums[0] <= 0 && nums[len(nums)-1] >= 0 {
		for i := 0; i < len(nums)-2; i++ {
			if i == 0 || nums[i] != nums[i-1] {
				l, r := i+1, len(nums)-1
				for l < r {
					sum := nums[i] + nums[l] + nums[r]
					if sum == 0 {
						res = append(res, []int{nums[i], nums[l], nums[r]})
						r--
						for l < r && nums[r] == nums[r+1] {
							r--
						}
						l++
						for l < r && nums[l] == nums[l-1] {
							l++
						}
					} else if sum > 0 {
						r--
						for l < r && nums[r] == nums[r+1] {
							r--
						}
					} else {
						l++
						for l < r && nums[l] == nums[l-1] {
							l++
						}
					}
				}
			}
		}
	}
	return res
}
