package arraysort

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		l1, l2 = l2, l1
		nums1, nums2 = nums2, nums1
	}
	left, right := 0, l1
	k := (l1 + l2 + 1) / 2
	for left <= right {
		i := left + (right-left)/2
		j := k - i
		if i > 0 && nums1[i-1] > nums2[j] {
			right = i - 1
		} else if i < l1 && nums1[i] < nums2[j-1] {
			left = i + 1
		} else {
			var maxl int
			if i == 0 {
				maxl = nums2[j-1]
			} else if j == 0 {
				maxl = nums1[i-1]
			} else {
				maxl = max(nums1[i-1], nums2[j-1])
			}
			if (l1+l2)%2 == 1 {
				return float64(maxl)
			}
			var minr int
			if i == l1 {
				minr = nums2[j]
			} else if j == l2 {
				minr = nums1[i]
			} else {
				minr = min(nums1[i], nums2[j])
			}
			return float64(maxl+minr) / 2
		}
	}
	return 0.0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
