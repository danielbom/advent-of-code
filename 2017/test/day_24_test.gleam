import gleeunit

import gleam/list

import day_24

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input = "0/2\n2/2\n2/3\n3/4\n3/5\n0/1\n10/1\n9/10"
  let inputs = [#(input, 31)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_24.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  let input = "0/2\n2/2\n2/3\n3/4\n3/5\n0/1\n10/1\n9/10"
  let inputs = [#(input, 19)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_24.part2(input)
    assert result == expected
  })
}
