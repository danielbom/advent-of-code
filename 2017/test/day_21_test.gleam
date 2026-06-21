import gleeunit

import gleam/list

import day_21

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let input =
    "../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#"
  let inputs = [#(input, 12)]

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_21.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  let inputs = []

  list.map(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_21.part2(input)
    assert result == expected
  })
}
