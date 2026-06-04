import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

fn parse(s: String) -> Int {
  let assert Ok(value) = int.parse(s)
  value
}

fn naive_loop(
  spin: Dict(Int, Int),
  index: Int,
  step: Int,
  count: Int,
  end: Int,
) {
  case count <= end {
    False -> {
      let assert Ok(value) = dict.get(spin, { index + 1 } % count)
      value
    }
    True -> {
      let index = { index + step } % count
      let index = { index + 1 } % { count + 1 }
      let spin =
        dict.fold(spin, [], fn(acc, key, value) {
          case key < index {
            True -> [#(key, value), ..acc]
            False -> [#(key + 1, value), ..acc]
          }
        })
        |> dict.from_list()
        |> dict.insert(index, count)
      naive_loop(spin, index, step, count + 1, end)
    }
  }
}

fn naive(step: Int, end: Int) {
  naive_loop(dict.new(), 0, step, 0, end)
}

pub fn part1(s: String) {
  let step = parse(s)
  naive(step, 2017)
}

fn zipper_step(left: List(a), right: List(a), count: Int) {
  case count > 0, left, right {
    _, [], [] -> #(left, right)
    _, [], _ -> zipper_step(list.reverse(right), [], count)
    False, _, _ -> #(left, right)
    True, [head, ..tail], _ -> zipper_step(tail, [head, ..right], count - 1)
  }
}

fn zipper_find_loop(left: List(a), right: List(a), zipper: List(a), value: a) {
  case left, right {
    [], [] -> #(list.reverse(zipper), [])
    [head, ..tail], _ -> {
      case head == value {
        True -> #(left, right)
        False -> zipper_find_loop(tail, right, [head, ..zipper], value)
      }
    }
    _, _ -> zipper_find_loop(list.reverse(right), [], zipper, value)
  }
}

fn zipper_find(left: List(a), right: List(a), value: a) {
  zipper_find_loop(left, right, [], value)
}

fn optimized_loop(left, right, step: Int, count: Int, end: Int) {
  case count <= end {
    False -> {
      let #(left, right) = zipper_find(left, right, 0)
      let #(left, _right) = zipper_step(left, right, 1)
      case list.first(left) {
        Ok(next) -> next
        Error(_) -> -1
      }
    }
    True -> {
      let #(left, right) = zipper_step(left, right, step + 1)
      let left = [count, ..left]
      optimized_loop(left, right, step, count + 1, end)
    }
  }
}

fn optimized(step: Int, end: Int) {
  optimized_loop([], [], step, 0, end)
}

pub fn part2(s: String) {
  let step = parse(s)
  optimized(step, 50_000_000)
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-17.txt") |> string.trim()
  io.println("Day 17")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
