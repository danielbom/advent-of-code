import gleeunit

import gleam/list
import gleam/string

import day_20

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn part1_test() {
  let inputs = [
    #(
      [
        "p=<-3,0,0>, v=<-2,0,0>, a=<1,0,0>",
        "p=<-4,0,0>, v=< 0,0,0>, a=<2,0,0>",
      ]
        |> string.join("\n"),
      0,
    ),
    #(
      [
        "p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>",
        "p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>",
      ]
        |> string.join("\n"),
      0,
    ),
    #(
      [
        "p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>",
        "p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>",
      ]
        |> string.join("\n"),
      1,
    ),
  ]

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_20.part1(input)
    assert result == expected
  })
}

pub fn part2_test() {
  let inputs = []

  list.each(inputs, fn(p) {
    let #(input, expected) = p
    let result = day_20.part2(input)
    assert result == expected
  })
}
