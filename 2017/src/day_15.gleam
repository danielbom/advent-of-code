import gleam/int
import gleam/io
import gleam/string

import utils

const max = 2_147_483_647

/// Lower 16 bits mask: (1 << 16) - 1
const mask = 65_535

fn parse(s: String) -> #(Int, Int) {
  let assert Ok(#(a, b)) = string.split_once(string.trim(s), on: "\n")
  let assert Ok(#(_, a)) = string.split_once(a, on: "with ")
  let assert Ok(#(_, b)) = string.split_once(b, on: "with ")
  let assert Ok(a) = int.parse(a)
  let assert Ok(b) = int.parse(b)
  #(a, b)
}

fn increment_if_match(count: Int, a: Int, b: Int) {
  case int.bitwise_and(a, mask) == int.bitwise_and(b, mask) {
    True -> count + 1
    False -> count
  }
}

fn count_matching_pairs(a: Int, b: Int, samples: Int, count: Int) -> Int {
  case samples > 0 {
    False -> count
    True -> {
      let count = increment_if_match(count, a, b)
      let a = { a * 16_807 } % max
      let b = { b * 48_271 } % max
      count_matching_pairs(a, b, samples - 1, count)
    }
  }
}

fn next_multiple_of(value: Int, factor: Int, multiple_of: Int) {
  let value = { value * factor } % max
  case value % multiple_of == 0 {
    True -> value
    False -> next_multiple_of(value, factor, multiple_of)
  }
}

fn count_matching_pairs2(a: Int, b: Int, samples: Int, count: Int) -> Int {
  case samples > 0 {
    False -> count
    True -> {
      let count = increment_if_match(count, a, b)
      let a = next_multiple_of(a, 16_807, 4)
      let b = next_multiple_of(b, 48_271, 8)
      count_matching_pairs2(a, b, samples - 1, count)
    }
  }
}

pub fn part1(s: String) {
  let #(seed_a, seed_b) = parse(s)
  count_matching_pairs(seed_a, seed_b, 40_000_000, 0)
}

pub fn part2(s: String) {
  let #(seed_a, seed_b) = parse(s)
  count_matching_pairs2(seed_a, seed_b, 5_000_000, 0)
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-15.txt") |> string.trim()
  io.println("Day 15")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
