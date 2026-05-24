import gleam/bool
import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/option.{type Option, None, Some}
import gleam/result
import gleam/string

import utils

fn abs(x: Int) -> Int {
  case x >= 0 {
    True -> x
    False -> -x
  }
}

fn manhattan_distance(x: Int, y: Int) -> Int {
  abs(x) + abs(y)
}

type Point =
  #(Int, Int)

fn point_scale(p: Point, scalar: Int) -> Point {
  let #(x, y) = p
  #(x * scalar, y * scalar)
}

fn point_add(lhs: Point, rhs: Point) -> Point {
  let #(x1, y1) = lhs
  let #(x2, y2) = rhs
  #(x1 + x2, y1 + y2)
}

type Dir {
  Down
  Right
  Top
  Left
}

fn dir_next(dir: Dir) -> Dir {
  case dir {
    Down -> Right
    Right -> Top
    Top -> Left
    Left -> Down
  }
}

fn dir_to_point(dir: Dir) -> Point {
  case dir {
    Right -> #(1, 0)
    Top -> #(0, -1)
    Left -> #(-1, 0)
    Down -> #(0, 1)
  }
}

type Spiral {
  Spiral(index: Int, point: Point, count: Int, dir: Dir)
}

fn spiral_from(s: Spiral, count: Int) -> Spiral {
  let dir = dir_next(s.dir)
  let next_point =
    dir_to_point(dir)
    |> point_scale(count)
    |> point_add(s.point)
  Spiral(s.index + count, next_point, count, dir)
}

fn spiral_next(s: Spiral) -> #(Spiral, Spiral) {
  let next_count = s.count + 1
  let s1 = spiral_from(s, next_count)
  let s2 = spiral_from(s1, next_count)
  #(s1, s2)
}

fn find_spiral_loop(ix: Int, s: Spiral) -> Spiral {
  use <- bool.guard(s.index >= ix, s)
  let #(s1, s2) = spiral_next(s)
  use <- bool.guard(s1.index >= ix, s1)
  find_spiral_loop(ix, s2)
}

fn find_spiral(ix: Int) -> Spiral {
  let start = Spiral(1, #(0, 0), 0, Down)
  find_spiral_loop(ix, start)
}

fn adjusted_steps_shortest_path(s: Spiral, index: Int) -> Int {
  let count = s.index - index
  let #(x, y) =
    dir_to_point(s.dir)
    |> point_scale(-count)
    |> point_add(s.point)
  manhattan_distance(x, y)
}

fn steps_shortest_path(index: Int) -> Int {
  let s = find_spiral(index)
  adjusted_steps_shortest_path(s, index)
}

pub fn part1(s: String) -> Int {
  case int.parse(s) {
    Ok(x) -> steps_shortest_path(x)
    Error(_) -> -1
  }
}

type Map =
  Dict(#(Int, Int), Int)

/// Computes the sum of all adjacent neighbor values around a point.
///
/// Missing neighbors contribute `0` to the total.
fn sum_of_computed_adjacent_neighbors(map: Map, p: Point) -> Int {
  [
    point_add(p, #(-1, -1)),
    point_add(p, #(0, -1)),
    point_add(p, #(1, -1)),
    point_add(p, #(-1, 0)),
    //point_add(p, #(0, 0)),
    point_add(p, #(1, 0)),
    point_add(p, #(-1, 1)),
    point_add(p, #(0, 1)),
    point_add(p, #(1, 1)),
  ]
  |> list.map(dict.get(map, _))
  |> result.values()
  |> int.sum()
}

fn fill_spiral_until_greater_loop(
  map: Map,
  x: Int,
  curr: Point,
  end: Point,
  inc: Point,
) -> #(Map, Option(Int)) {
  let value = sum_of_computed_adjacent_neighbors(map, curr)
  let new_map = dict.insert(map, curr, value)
  use <- bool.lazy_guard(value > x, fn() { #(new_map, Some(value)) })
  use <- bool.lazy_guard(curr == end, fn() { #(new_map, None) })
  fill_spiral_until_greater_loop(new_map, x, point_add(curr, inc), end, inc)
}

/// Updates spiral positions from `begin` to `end`,
/// storing computed values in the map until a value
/// greater than `x` is found.
///
/// Returns the updated map and the matching value, if any.
fn fill_spiral_until_greater(
  map: Map,
  x: Int,
  begin: Spiral,
  end: Spiral,
) -> #(Map, Option(Int)) {
  let dir = dir_to_point(end.dir)
  fill_spiral_until_greater_loop(
    map,
    x,
    point_add(begin.point, dir),
    end.point,
    dir,
  )
}

/// Iterates through spiral positions and returns
/// the first computed value greater than `x`.
///
/// Each position value is derived from its adjacent neighbors
/// and stored in the map during traversal.
///
/// Returns `Some(value)` if found, otherwise `None`.
fn find_first_value_greater_than_loop(
  x: Int,
  s: Spiral,
  map: Map,
) -> Option(Int) {
  let #(s1, s2) = spiral_next(s)
  let #(m1, end1) = fill_spiral_until_greater(map, x, s, s1)
  use <- option.lazy_or(end1)
  let #(m2, end2) = fill_spiral_until_greater(m1, x, s1, s2)
  use <- option.lazy_or(end2)
  find_first_value_greater_than_loop(x, s2, m2)
}

fn find_first_value_greater_than(x: Int) -> Int {
  use <- bool.guard(x < 1, 1)
  let start = Spiral(1, #(0, 0), 0, Down)
  let map = dict.from_list([#(start.point, 1)])
  let assert Some(result) = find_first_value_greater_than_loop(x, start, map)
  result
}

pub fn part2(s: String) -> Int {
  case int.parse(s) {
    Ok(x) -> find_first_value_greater_than(x)
    Error(_) -> -1
  }
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-03.txt") |> string.trim()
  io.println("Day 03")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
