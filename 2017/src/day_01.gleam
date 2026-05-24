import gleam/int
import gleam/io
import gleam/list
import gleam/pair
import gleam/result
import gleam/string

import utils

fn sum_digits_windowed(s: String, gap: Int) -> Int {
  let graphemes = string.to_graphemes(s)
  graphemes
  |> list.zip(graphemes |> list.drop(gap) |> list.append(graphemes))
  |> list.filter(fn(p) { pair.first(p) == pair.second(p) })
  |> list.map(fn(p) { pair.first(p) |> int.parse })
  |> result.all()
  |> result.map(int.sum)
  |> result.unwrap(-1)
}

/// Computes the sum of all digits that match the next digit
/// in a circular sequence.
///
/// Each digit is compared with the following digit,
/// where the last digit wraps around to the first.
pub fn part1(s: String) -> Int {
  sum_digits_windowed(s, 1)
}

/// Computes the sum of all digits that match the digit
/// halfway around the circular sequence.
///
/// Each digit is compared with the digit located
/// `len / 2` positions ahead, wrapping around the string. 
pub fn part2(s: String) -> Int {
  sum_digits_windowed(s, string.length(s) / 2)
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-01.txt") |> string.trim()
  io.println("Day 01")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
