import gleeunit

import gleam/list

import day_19

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input =
    "    |          \n    |  +--+    \n    A  |  C    \nF---|----E|--+ \n    |  |  |  D \n    +B-+  +--+"
  let inputs = [#(input, "ABCDEF")]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_19.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  let input =
    "    |          \n    |  +--+    \n    A  |  C    \nF---|----E|--+ \n    |  |  |  D \n    +B-+  +--+"
  let inputs = [#(input, 38)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_19.part2(input)
    assert result == expected
  })
}
