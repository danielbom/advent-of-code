import gleeunit

import gleam/list

import day_17

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let inputs = [
    //
    #("3", 638),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_17.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  let inputs = [
    //
    #("3", 1_222_153),
  ]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_17.part2(input)
    assert result == expected
  })
}
