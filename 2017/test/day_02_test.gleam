import gleeunit

import gleam/list

import day_02

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input =
    "5 1 9 5
7 5 3
2 4 6 8"
  let inputs = [#(input, 18)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_02.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  let input =
    "5 9 2 8
9 4 7 3
3 8 6 5"
  let inputs = [#(input, 9)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_02.part2(input)
    assert result == expected
  })
}
