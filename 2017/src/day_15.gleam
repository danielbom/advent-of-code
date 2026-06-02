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

type Gen {
  Gen(factor: Int, value: Int, next: fn(Gen) -> Gen)
}

fn generator_next(gen: Gen) -> Gen {
  let value = { gen.value * gen.factor } % max
  Gen(..gen, value:)
}

fn generators_match(a: Gen, b: Gen) {
  int.bitwise_and(a.value, mask) == int.bitwise_and(b.value, mask)
}

fn count_matching_pairs_loop(a: Gen, b: Gen, samples: Int, count: Int) -> Int {
  case samples > 0 {
    False -> count
    True -> {
      let count = case generators_match(a, b) {
        True -> count + 1
        False -> count
      }
      let a = a.next(a)
      let b = b.next(b)
      count_matching_pairs_loop(a, b, samples - 1, count)
    }
  }
}

fn count_matching_pairs(a: Gen, b: Gen, samples: Int) -> Int {
  count_matching_pairs_loop(a, b, samples, 0)
}

pub fn part1(s: String) {
  let #(seed_a, seed_b) = parse(s)
  count_matching_pairs(
    Gen(16_807, seed_a, generator_next),
    Gen(48_271, seed_b, generator_next),
    40_000_000,
  )
}

fn generator_next_multiple_of(gen: Gen, multiple_of: Int) {
  let g = generator_next(gen)
  case g.value % multiple_of == 0 {
    True -> g
    False -> generator_next_multiple_of(g, multiple_of)
  }
}

pub fn part2(s: String) {
  let #(seed_a, seed_b) = parse(s)
  count_matching_pairs(
    Gen(16_807, seed_a, fn(gen) { generator_next_multiple_of(gen, 4) }),
    Gen(48_271, seed_b, fn(gen) { generator_next_multiple_of(gen, 8) }),
    5_000_000,
  )
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-15.txt") |> string.trim()
  io.println("Day 15")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
