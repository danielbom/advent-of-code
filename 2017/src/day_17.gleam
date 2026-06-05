import gleam/int
import gleam/io
import gleam/list
import gleam/string

import iv

import utils

fn parse(s: String) -> Int {
  let assert Ok(value) = int.parse(s)
  value
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

pub fn spinlock_zip_loop(left, right, step: Int, count: Int, end: Int) {
  case count <= end {
    False -> {
      let #(left, right) = zipper_find(left, right, end)
      let #(left, _right) = zipper_step(left, right, 1)
      case list.first(left) {
        Ok(next) -> next
        Error(_) -> -1
      }
    }
    True -> {
      let #(left, right) = zipper_step(left, right, step + 1)
      let left = [count, ..left]
      spinlock_zip_loop(left, right, step, count + 1, end)
    }
  }
}

fn next_index(current: Int, step: Int, size: Int) -> Int {
  { { current + step } % size } + 1
}

fn insert_at(buffer: List(Int), index: Int, value: Int) -> List(Int) {
  let #(left, right) = list.split(buffer, index)
  list.append(left, [value, ..right])
}

fn find_next_loop(list: List(a), first: Result(a, Nil), value: a) {
  case list {
    [] -> Error(Nil)
    [head, next, ..] if head == value -> Ok(next)
    [head, ..] if head == value -> first
    [_, ..tail] -> find_next_loop(tail, first, value)
  }
}

fn find_next(list: List(a), value: a) {
  case list {
    [] -> Error(Nil)
    [first, ..] -> find_next_loop(list, Ok(first), value)
  }
}

pub fn spinlock_list_loop(
  buffer: List(Int),
  index: Int,
  step: Int,
  value: Int,
  end: Int,
) -> Int {
  case value > end {
    True -> {
      let assert Ok(result) = find_next(buffer, end)
      result
    }

    False -> {
      let index = next_index(index, step, value)
      let buffer = insert_at(buffer, index, value)
      spinlock_list_loop(buffer, index, step, value + 1, end)
    }
  }
}

pub fn spinlock_array_loop(
  spin: iv.Array(Int),
  index: Int,
  step: Int,
  count: Int,
  end: Int,
) {
  case count <= end {
    False -> iv.get_or_default(spin, { index + 1 } % count, 0)
    True -> {
      let index = { { index + step } % count } + 1
      let spin = iv.insert_clamped(spin, index, count)
      spinlock_array_loop(spin, index, step, count + 1, end)
    }
  }
}

pub fn part1(input: String) -> Int {
  let step = parse(input)

  spinlock_zip_loop([], [], step, 0, 2017)
  // spinlock_list_loop([], 0, step, 0, 2017)
  // spinlock_array_loop(iv.new(), 0, step, 0, 2017)
  // spinlock_array_loop(iv.new(), 0, step, 0, 1_000_000) // 12602.0608ms
  // spinlock_zip_loop([], [], step, 0, 1_000_000) // 6996.7872ms
  // spinlock_list_loop([], 0, 3, 0, 100_000) // 85677.2608ms (8603)
}

fn value_after_zero_loop(
  index: Int,
  step: Int,
  value: Int,
  end: Int,
  result: Int,
) -> Int {
  case value > end {
    True -> result

    False -> {
      let index = next_index(index, step, value)
      let result = case index == 1 {
        True -> value
        False -> result
      }
      value_after_zero_loop(index, step, value + 1, end, result)
    }
  }
}

pub fn part2(input: String) -> Int {
  let step = parse(input)
  value_after_zero_loop(0, step, 1, 50_000_000, 0)
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-17.txt") |> string.trim()
  io.println("Day 17")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
