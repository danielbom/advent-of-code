import gleeunit

import gleam/list

import day_14

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  // Too timing consuming - uncomment to test it
  let inputs = [
    // #("flqrgnkx", 8108),
  ]

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_14.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  // Too timing consuming - uncomment to test it
  let inputs = [
    // #("flqrgnkx", 1242),
  ]

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_14.part2(input)
    assert result == expected
  })
}
