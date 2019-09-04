const restoreIpAddressesRecursive = (restStr, ans, partCount, pre) => {
  if (partCount === 4) {
    if (restStr.length === 0) ans.push(pre);
    return;
  }
  for (let i = 1; i < 4; i += 1) {
    if (restStr.length < i) break;
    const newPartStr = restStr.slice(0, i);
    const newPart = parseInt(newPartStr, 10);
    if (newPart > 255 || (i > 1 && newPartStr[0] === '0')) break;
    restoreIpAddressesRecursive(
      restStr.slice(i),
      ans,
      partCount + 1,
      pre + newPartStr + (partCount === 3 ? '' : '.')
    );
  }
};

/**
 * @param {string} s
 * @return {string[]}
 */
const restoreIpAddresses = s => {
  const ans = [];
  restoreIpAddressesRecursive(s, ans, 0, '');
  return ans;
};

export default restoreIpAddresses;
