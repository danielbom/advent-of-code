import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

fn strict_int_parse(x: String) -> Int {
  case int.parse(x) {
    Ok(value) -> value
    Error(_) -> panic as string.append("invalid int: ", x)
  }
}

fn parse(s: String) -> Dict(Int, Int) {
  string.split(s, on: "\n")
  |> list.index_map(fn(x, i) { #(i, strict_int_parse(x)) })
  |> dict.from_list()
}

fn count_steps_1(steps: Int, index: Int, view: Dict(Int, Int)) -> Int {
  case dict.get(view, index) {
    Ok(count) -> {
      let next_count = count + 1
      count_steps_1(
        steps + 1,
        index + count,
        dict.insert(view, index, next_count),
      )
    }
    Error(_) -> steps
  }
}

pub fn part1(s: String) -> Int {
  count_steps_1(0, 0, parse(s))
}

fn count_steps_2(steps: Int, index: Int, view: Dict(Int, Int)) -> Int {
  case dict.get(view, index) {
    Ok(count) -> {
      let next_count = case count >= 3 {
        True -> count - 1
        False -> count + 1
      }
      count_steps_2(
        steps + 1,
        index + count,
        dict.insert(view, index, next_count),
      )
    }
    Error(_) -> steps
  }
}

pub fn part2(s: String) -> Int {
  count_steps_2(0, 0, parse(s))
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-05.txt") |> string.trim()
  io.println("Day 05")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
