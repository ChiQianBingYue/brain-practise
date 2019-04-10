const maxAreaOfIsland = require('./maxAreaOfIsland');

test('simple', () => {
  expect(
    maxAreaOfIsland([
      [0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
      [0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0],
      [0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0]
    ])
  ).toBe(6);
});

test('zero', () => {
  expect(maxAreaOfIsland([])).toBe(0);
  expect(maxAreaOfIsland([0])).toBe(0);
  expect(maxAreaOfIsland([0, 0, 0])).toBe(0);
  expect(maxAreaOfIsland([0], [0])).toBe(0);
});

test('one', () => {
  expect(maxAreaOfIsland([[1]])).toBe(1);
});
