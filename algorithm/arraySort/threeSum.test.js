const threeSum = require('./threeSum');

test('simple', () => {
  expect(threeSum([-1, 0, 1, 2, -1, -4])).toEqual([[-1, -1, 2], [-1, 0, 1]]);
});

test('duplication', () => {
  expect(threeSum([1, 1, -1, 1, -1, 0])).toEqual([[-1, 0, 1]]);
});
