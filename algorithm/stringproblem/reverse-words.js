const reverseWords = str =>
  str
    .trim()
    .split(' ')
    .filter(s => !!s)
    .reverse()
    .join(' ');

module.exports = reverseWords;
