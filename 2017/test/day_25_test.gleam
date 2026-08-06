import gleeunit

import gleam/list

import day_25

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input =
    "Begin in state A.
  Perform a diagnostic checksum after 6 steps.

  In state A:
    If the current value is 0:
      - Write the value 1.
      - Move one slot to the right.
      - Continue with state B.
    If the current value is 1:
      - Write the value 0.
      - Move one slot to the left.
      - Continue with state B.

  In state B:
    If the current value is 0:
      - Write the value 1.
      - Move one slot to the left.
      - Continue with state A.
    If the current value is 1:
      - Write the value 1.
      - Move one slot to the right.
      - Continue with state A."
  let inputs = [#(input, 3)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_25.part1(input)
    assert result == expected
  })
}
