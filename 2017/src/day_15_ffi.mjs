const MASK = 0xffff;

export function incrementIfMatch(count, a, b) {
  return count + ((a & MASK) === (b & MASK));
}
