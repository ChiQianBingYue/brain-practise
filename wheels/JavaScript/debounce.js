/**
 * 方法防抖，持续执行同一个方法，只执行一次，除非等待一段时间后重新执行，搜索建议
 * @param {Function} func 执行的方法
 * @param {*} delay 等待的时间延时
 * @param {*} immediate 是立即执行，还是最后执行
 */
function debounce(func, delay, immediate) {
  let timeout;
  return function debounced(...args) {
    if (immediate && !timeout) func.apply(this, args);
    clearTimeout(timeout);
    timeout = setTimeout(() => {
      timeout = null;
      if (!immediate) func.apply(this, args);
    }, delay);
  };
}

module.exports = debounce;
