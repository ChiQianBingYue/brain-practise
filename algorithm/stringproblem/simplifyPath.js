class Stack {
  constructor() {
    this.stack = [];
  }

  push(item) {
    this.stack.push(item);
  }

  pop() {
    return this.stack.pop();
  }

  top() {
    return this.stack[this.stack.length - 1];
  }

  get length() {
    return this.stack.length;
  }
}

class StringBuilder {
  constructor() {
    this.strings = [];
  }

  concat(value) {
    this.strings.push(value);
    return this;
  }

  leftConcat(value) {
    this.strings.unshift(value);
    return this;
  }

  string() {
    return this.strings.join('');
  }

  clear() {
    this.strings = [];
    return this;
  }

  isEmpty() {
    return this.strings.length === 0;
  }
}

function simplifyPath(path) {
  const stack = new Stack();
  const sb = new StringBuilder();
  function handleNewPath(str) {
    if (stack.top() === str || str === '..') {
      stack.pop();
    } else if (str !== '.') {
      stack.push(str);
    }
  }
  for (let i = 0; i < path.length; i += 1) {
    if (path[i] === '/' && !sb.isEmpty()) {
      const str = sb.string();
      handleNewPath(str);
      sb.clear();
    } else if (path[i] !== '' && path[i] !== '/') {
      sb.concat(path[i]);
    }
    // if '' or '.' continue
  }
  if (!sb.isEmpty()) {
    const str = sb.string();
    handleNewPath(str);
    sb.clear();
  }
  sb.clear();
  while (stack.length > 0) {
    sb.leftConcat(stack.pop());
    sb.leftConcat('/');
  }
  return sb.isEmpty() ? '/' : sb.string();
}
module.exports = simplifyPath;
