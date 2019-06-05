package arraySort

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := nums1, nums2
	l1, l2 := len(n1), len(n2)
	if l1 > l2 {
		l1, l2 = l2, l1
		n1, n2 = n2, n1
	}
	left, right := 0, l1
	k := (l1 + l2 + 1) / 2
	for left <= right {
		i := left + (right-left)/2
		j := k - i
		if i > 0 && n1[i-1] > n2[j] {
			right = i - 1
		} else if i < l1 && n2[j-1] > n1[i] {
			left = i + 1
		} else {
			// i is good
			var lmax int
			if i == 0 {
				lmax = n2[j-1]
			} else if j == 0 {
				lmax = n1[i-1]
			} else {
				lmax = max(n2[j-1], n1[i-1])
			}
			if (l1+l2)%2 == 1 {
				return float64(lmax)
			}
			var rmin int
			if i == l1 {
				rmin = n2[j]
			} else if j == l2 {
				rmin = n1[i]
			} else {
				rmin = min(n1[i], n2[j])
			}
			fmt.Println(lmax, rmin)
			return float64(lmax+rmin) / 2
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
