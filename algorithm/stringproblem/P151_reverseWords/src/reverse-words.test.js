const reverseWords = require('./reverse-words');

test('the sky is blue', () => {
  expect(reverseWords('the sky is blue')).toBe('blue is sky the');
});

test('a good   example', () => {
  expect(reverseWords('a good   example')).toBe('example good a');
});

test('  hello world!  ', () => {
  expect(reverseWords('  hello world!  ')).toBe('world! hello');
});
