// 思路
// 要记住上次执行的时间，然后计算当前时间和上次时间的差值与limit比较
// 差值比limit大，就是过了节流，执行
// 如果小，就不运行。但如果有trilling，则要在节流时间后执行。
// 默认的上次执行时间是0，这样就一定会执行首次方法
//    但如果 leading 是false，就把首次的上次是执行时间改为当期时间。
// context 和 args 在每次执行后清空

// polyfill Date.now
if (!Date.now) {
  Date.now = function now() {
    return new Date().getTime();
  };
}

/**
 * 限制函数执行的频率
 * @param {Function} func 方法
 * @param {number} limit 每多少秒执行一次
 * @param {object} [option] 可选项，是否 leading 或 trilling
 */
function throttle(func, limit, options = { leading: true, trilling: false }) {
  let timeout;
  let prev = 0;
  return function throttled(...args) {
    const now = Date.now();
    if (!options.leading && prev === 0) prev = now;
    const rest = limit - (now - prev);
    let result;
    if (rest <= 0) {
      clearTimeout(timeout); // 时间可能不准
      timeout = null;
      prev = now;
      result = func.apply(this, args);
    } else if (!timeout && options.trilling) {
      timeout = setTimeout(() => {
        prev = !options.leading ? 0 : Date.now();
        result = func.apply(this, args);
        timeout = null;
      }, rest);
    }
    return result;
  };
}

module.exports = throttle;
