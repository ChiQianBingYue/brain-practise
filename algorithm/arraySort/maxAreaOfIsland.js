const bfs = (grid, i, j, visited) => {
  const q = [[i, j]];
  visited[i][j] = 1; // eslint-disable-line
  let res = 0;
  const dx = [0, 1, 0, -1];
  const dy = [1, 0, -1, 0];
  while (q.length) {
    const [oX, oY] = q.pop();
    res += 1;
    for (let k = 0; k < 4; k += 1) {
      const nX = oX + dx[k];
      const nY = oY + dy[k];
      if (nX < 0 || nX >= grid.length || nY < 0 || nY >= grid[0].length) {
        continue; // eslint-disable-line
      }
      if (grid[nX][nY] === 1 && !visited[nX][nY]) {
        q.push([nX, nY]);
        visited[nX][nY] = 1; // eslint-disable-line
      }
    }
  }
  return res;
};

/**
 * @param {number[][]} grid
 * @return {number}
 */
const maxAreaOfIsland = grid => {
  if (grid.length < 1 || grid[0].length < 1) {
    return 0;
  }
  const m = grid.length;
  const n = grid[0].length;
  // 0 is not visited, 1 is visited
  const visited = new Array(m);
  for (let a = 0; a < m; a += 1) {
    visited[a] = new Array(n).fill(0);
  }
  let ans = 0;
  for (let i = 0; i < m; i += 1) {
    for (let j = 0; j < n; j += 1) {
      if (grid[i][j] === 1 && !visited[i][j]) {
        ans = Math.max(ans, bfs(grid, i, j, visited));
      }
    }
  }
  return ans;
};

setTimeout(() => {
  console.log('timer');
}, 0);

setImmediate(() => {
  console.log('setImmediate');
}, 0);

// fs.open('./maxAreaOfIsland.go', (err, fd) => {
//   console.log('fs opne');
//   Promise.resolve().then(() => {
//     console.log('fs resolve');
//   });
//   process.nextTick(() => {
//     console.log('fs tick');
//   });
//   fs.close(fd, () => {
//     console.log('fs close');
//     Promise.resolve().then(() => {
//       console.log('fs c resolve');
//     });
//     process.nextTick(() => {
//       console.log('fs c tick');
//     });
//   });
// });

module.exports = maxAreaOfIsland;
