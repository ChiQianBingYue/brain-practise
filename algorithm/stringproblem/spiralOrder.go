package stringproblem

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	m := len(matrix)

	if m == 0 {
		return res
	}
	n := len(matrix[0])
	if n == 0 {
		return res
	}
	dir := [][]int{
		[]int{1, 0},
		[]int{0, 1},
		[]int{-1, 0},
		[]int{0, -1},
	}
	l, r, t, b := 0, n-1, 0, m-1
	i, j := 0, 0
	res = append(res, matrix[j][i])
	d := 0
	for {
		if d == 0 && i+1 > r {
			d = 1
			t++
		}
		if d == 1 && j+1 > b {
			d = 2
			r--
		}
		if d == 2 && i-1 < l {
			d = 3
			b--
		}
		if d == 3 && j-1 < t {
			d = 0
			l++
		}
		if l > r || t > b {
			break
		}
		i, j = i+dir[d][0], j+dir[d][1]

		res = append(res, matrix[j][i])
	}
	return res
}
