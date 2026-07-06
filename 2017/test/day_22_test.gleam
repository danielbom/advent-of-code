import gleeunit

import gleam/list

import day_22

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input = "..#\n#..\n..."
  let inputs = [#(input, 5587)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_22.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  let inputs = []

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_22.part2(input)
    assert result == expected
  })
}
