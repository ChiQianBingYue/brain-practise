const swap = (nums, i, j) => {
  [nums[i], nums[j]] = [nums[j], nums[i]]; // eslint-disable-line
};

const qSortPartition = (nums, l, r) => {
  if (l >= r) return;
  const pivot = nums[l + Math.trunc((r - l) / 2)];
  let i = l;
  let j = r;
  while (i <= j) {
    while (i <= j && nums[i] < pivot) {
      i += 1;
    }
    while (i <= j && nums[j] > pivot) {
      j -= 1;
    }
    if (i <= j) {
      swap(nums, i, j);
      i += 1;
      j -= 1;
    }
  }
  qSortPartition(nums, l, j);
  qSortPartition(nums, i, r);
};

const qSort = nums => {
  return qSortPartition(nums, 0, nums.length - 1);
};

const threeSum = nums => {
  qSort(nums);
  // nums.sort((a, b) => a - b);
  const res = [];
  for (let i = 0; i < nums.length - 2; i += 1) {
    if (i === 0 || nums[i] !== nums[i - 1]) {
      let l = i + 1;
      let r = nums.length - 1;
      while (l < r) {
        const sum = nums[i] + nums[l] + nums[r];
        if (sum === 0) {
          res.push([nums[i], nums[l], nums[r]]);
          do {
            l += 1;
          } while (nums[l - 1] === nums[l]);
          do {
            r -= 1;
          } while (nums[r] === nums[r + 1]);
        } else if (sum < 0) {
          do {
            l += 1;
          } while (nums[l - 1] === nums[l]);
        } else if (sum > 0) {
          do {
            r -= 1;
          } while (nums[r] === nums[r + 1]);
        }
      }
    }
  }
  return res;
};

module.exports = threeSum;
