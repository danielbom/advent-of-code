import * as $iv from "../iv/iv.mjs";

export function countSteps1(steps, index, view) {
  let xs = $iv.to_list(view).toArray();
  while (0 <= index && index < xs.length) {
    let count = xs[index];
    xs[index] = count + 1;
    index += count;
    steps++;
  }
  return steps;
}

export function countSteps2(steps, index, view) {
  let xs = $iv.to_list(view).toArray();
  while (0 <= index && index < xs.length) {
    let count = xs[index];
    xs[index] = count >= 3 ? count - 1 : count + 1;
    index += count;
    steps++;
  }
  return steps;
}
