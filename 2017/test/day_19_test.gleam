import gleeunit

import gleam/list

import day_19

pub fn main() -> Nil {
  gleeunit.main()
}

const input = "    |          \n    |  +--+    \n    A  |  C    \nF---|----E|--+ \n    |  |  |  D \n    +B-+  +--+"

pub fn grid_debug_test() {
  let expected =
    "    ~         \n    ~  ~~~~   \n    A  ~  C   \nF~~~~~~~~E~~~~\n    ~  ~  ~  D\n    ~B~~  ~~~~"
  let input = day_19.parse(input)
  let result = day_19.grid_debug(input)
  assert result == expected
}

pub fn part1_test() {
  let inputs = [#(input, "ABCDEF")]

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_19.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  let inputs = [#(input, 38)]

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_19.part2(input)
    assert result == expected
  })
}
