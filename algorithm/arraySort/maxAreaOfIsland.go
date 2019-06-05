package arraySort

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	visited := make([]int, m) // use bit map for row
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && visited[i]&(1<<uint(j)) == 0 {
				ans = max(bfs(grid, visited, i, j), ans)
			}
		}
	}
	return ans
}

type point struct {
	x int
	y int
}

func bfs(grid [][]int, visited []int, i, j int) int {
	queue := []point{}
	res := 0
	queue = append(queue, point{i, j})
	visited[i] = visited[i] | (1 << uint(j))
	dx := [...]int{-1, 0, 1, 0}
	dy := [...]int{0, -1, 0, 1}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		res++
		x, y := p.x, p.y
		for k := 0; k < 4; k++ {
			nx, ny := x+dx[k], y+dy[k]
			if nx < 0 || nx == len(grid) || ny < 0 || ny == len(grid[0]) {
				continue
			}
			if grid[nx][ny] == 1 && visited[nx]&(1<<uint(ny)) == 0 {
				queue = append(queue, point{nx, ny})
				visited[nx] = visited[nx] | (1 << uint(ny))
			}
		}
	}
	return res
}
