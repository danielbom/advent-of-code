import gleam/bool
import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
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

fn find_spiral_loop(x: Int, s: Spiral) -> Spiral {
  use <- bool.guard(s.index >= x, s)
  let #(s1, s2) = spiral_next(s)
  use <- bool.guard(s1.index >= x, s1)
  find_spiral_loop(x, s2)
}

fn find_spiral(x: Int) -> Spiral {
  let start = Spiral(1, #(0, 0), 0, Down)
  find_spiral_loop(x, start)
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

fn squares_update_list(
  m: Map,
  ix: Int,
  curr: Point,
  end: Point,
  inc: Point,
) -> #(Map, List(Int)) {
  let value =
    [
      dict.get(m, point_add(curr, #(-1, -1))),
      dict.get(m, point_add(curr, #(0, -1))),
      dict.get(m, point_add(curr, #(1, -1))),
      dict.get(m, point_add(curr, #(-1, 0))),
      //dict.get(m, point_add(curr, #(0, 0))),
      dict.get(m, point_add(curr, #(1, 0))),
      dict.get(m, point_add(curr, #(-1, 1))),
      dict.get(m, point_add(curr, #(0, 1))),
      dict.get(m, point_add(curr, #(1, 1))),
    ]
    |> result.values()
    |> int.sum()
  let nm = dict.insert(m, curr, value)
  use <- bool.lazy_guard(value > ix, fn() { #(nm, [value]) })
  use <- bool.lazy_guard(curr == end, fn() { #(nm, []) })
  squares_update_list(nm, ix, point_add(curr, inc), end, inc)
}

fn squares_update(
  m: Map,
  ix: Int,
  begin: Spiral,
  end: Spiral,
) -> #(Map, List(Int)) {
  let dir = dir_to_point(end.dir)
  squares_update_list(m, ix, point_add(begin.point, dir), end.point, dir)
}

fn find_first_largest_value(ix: Int, s0: Spiral, m0: Map) -> List(Int) {
  let #(s1, s2) = spiral_next(s0)
  let #(m1, end1) = squares_update(m0, ix, s0, s1)
  use <- bool.guard(!list.is_empty(end1), end1)
  let #(m2, end2) = squares_update(m1, ix, s1, s2)
  use <- bool.guard(!list.is_empty(end2), end2)
  find_first_largest_value(ix, s2, m2)
}

fn first_largest_value(index_squared: Int) -> Int {
  use <- bool.guard(index_squared < 1, 1)
  let start = Spiral(1, #(0, 0), 0, Down)
  let map: Map = dict.from_list([#(start.point, 1)])
  let assert [result] = find_first_largest_value(index_squared, start, map)
  result
}

pub fn part2(s: String) -> Int {
  case int.parse(s) {
    Ok(x) -> first_largest_value(x)
    Error(_) -> -1
  }
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-03.txt") |> string.trim()
  io.println("Day 03")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
