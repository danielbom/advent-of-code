import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string

import utils

type Banks =
  Dict(Int, Int)

fn parse(s: String) -> Banks {
  string.split(s, on: "\t")
  |> list.map(int.parse)
  |> result.all()
  |> result.unwrap([])
  |> list.index_map(fn(x, i) { #(i, x) })
  |> dict.from_list()
}

fn redistribute_blocks(banks: Banks, max: Int, index: Int) -> Banks {
  case max > 0, dict.get(banks, index) {
    False, _ -> banks
    True, Error(_) -> redistribute_blocks(banks, max, 0)
    True, Ok(blocks) -> {
      let result = dict.insert(banks, index, blocks + 1)
      redistribute_blocks(result, max - 1, index + 1)
    }
  }
}

fn find_largest_bank(banks: Banks) -> #(Int, Int) {
  dict.fold(banks, #(0, -1), fn(state, index, value) {
    let #(max, _) = state
    case max < value {
      True -> #(value, index)
      False -> state
    }
  })
}

fn redistribute_once(banks: Banks) -> Banks {
  let #(value, index) = find_largest_bank(banks)
  redistribute_blocks(dict.insert(banks, index, 0), value, index + 1)
}

/// Counts redistribution cycles until a bank configuration repeats.
///
/// Returns the repeated bank configuration and the number
/// of cycles executed before the repetition was detected.
fn find_repeated_configuration(
  seen: Dict(Banks, Bool),
  banks: Banks,
  count: Int,
) -> #(Banks, Int) {
  case dict.get(seen, banks) {
    Ok(_) -> #(banks, count)
    Error(_) -> {
      let seen = dict.insert(seen, banks, True)
      let banks = redistribute_once(banks)
      find_repeated_configuration(seen, banks, count + 1)
    }
  }
}

pub fn part1(s: String) -> Int {
  let banks = parse(s)
  let #(_, count) = find_repeated_configuration(dict.new(), banks, 0)
  count
}

pub fn part2(s: String) -> Int {
  let banks = parse(s)
  let #(banks, _) = find_repeated_configuration(dict.new(), banks, 0)
  let #(_, count) = find_repeated_configuration(dict.new(), banks, 0)
  count
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-06.txt") |> string.trim()
  io.println("Day 06")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
