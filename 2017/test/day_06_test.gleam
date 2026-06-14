import gleeunit

import gleam/list

import day_06

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let inputs = [
    //
    #("0\t2\t7\t0", 5),
  ]

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_06.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  let inputs = [
    //
    #("0\t2\t7\t0", 4),
  ]

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_06.part2(input)
    assert result == expected
  })
}
