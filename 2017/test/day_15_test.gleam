import gleeunit

import gleam/list

import day_15

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  // Too timing consuming - uncomment to test it
  let inputs = [
    //
  // #("a with 65\nb with 8921", 588),
  ]

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_15.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  // Too timing consuming - uncomment to test it
  let inputs = [
    //
  // #("a with 65\nb with 8921", 309),
  ]

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_15.part2(input)
    assert result == expected
  })
}
