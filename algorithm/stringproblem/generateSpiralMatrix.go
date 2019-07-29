package stringproblem

func generateSpiralMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	l, r, t, b := 0, n-1, 0, n-1
	x, y := 0, 0
	dir := 0
	dirs := [][]int{
		[]int{1, 0},  // to right
		[]int{0, 1},  // to bottom
		[]int{-1, 0}, // to left
		[]int{0, -1}, // to top
	}
	for i := 1; i <= n*n; i++ {
		res[y][x] = i
		if dir == 0 && x+1 > r {
			dir = 1
			t++
		} else if dir == 1 && y+1 > b {
			dir = 2
			r--
		} else if dir == 2 && x-1 < l {
			dir = 3
			b--
		} else if dir == 3 && y-1 < t {
			dir = 0
			l++
		}
		x, y = x+dirs[dir][0], y+dirs[dir][1]
	}
	return res
}
