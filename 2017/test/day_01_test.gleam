import gleeunit

import gleam/list

import day_01

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let inputs = [
    //
    #("1122", 3),
    #("1111", 4),
    #("1234", 0),
    #("91212129", 9),
  ]

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_01.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  let inputs = [
    #("1212", 6),
    #("1221", 0),
    #("123425", 4),
    #("123123", 12),
    #("12131415", 4),
  ]

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_01.part2(input)
    assert result == expected
  })
}
