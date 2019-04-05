const simplifyPath = require('./simplifyPath');

test('/home/', () => {
  expect(simplifyPath('/home/')).toBe('/home');
});

test('/../', () => {
  expect(simplifyPath('/../')).toBe('/');
});

test('/home//foo/', () => {
  expect(simplifyPath('/home//foo/')).toBe('/home/foo');
});

test('/a//b////c/d//././/..', () => {
  expect(simplifyPath('/a//b////c/d//././/..')).toBe('/a/b/c');
});
