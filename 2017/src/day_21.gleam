import gleam/int
import gleam/io
import gleam/list
import gleam/set.{type Set}
import gleam/string
import gleam/string_tree as st

import utils

fn fold_range(start, end, acc, fun) {
  case start < end {
    True -> fold_range(start + 1, end, fun(acc, start), fun)
    False -> acc
  }
}

/// Sparse grid representation storing only active (`#`) pixels.
/// Coordinates are encoded as #(row, column).
pub type PixelGrid {
  PixelGrid(items: Set(#(Int, Int)), size: Int)
}

pub fn grid_to_string(grid: PixelGrid) {
  fold_range(0, grid.size, st.new(), fn(sb, y) {
    let sb = case y > 0 {
      True -> st.append(sb, "\n")
      False -> sb
    }
    fold_range(0, grid.size, sb, fn(sb, x) {
      case set.contains(grid.items, #(y, x)) {
        True -> st.append(sb, "#")
        False -> st.append(sb, ".")
      }
    })
  })
  |> st.to_string()
}

/// Parses slash-separated grid patterns such as:
///
/// .#./..#/###
///
/// Into a sparse pixel grid.
fn parse_grid(input: String) -> PixelGrid {
  let rows = string.split(input, on: "/")
  let items =
    list.index_fold(rows, set.new(), fn(items, pixels, row) {
      pixels
      |> string.to_graphemes()
      |> list.index_fold(items, fn(items, pixel, col) {
        case pixel {
          "#" -> set.insert(items, #(row, col))
          "." -> items
          _ -> panic as string.append("unexpected pixel: ", pixel)
        }
      })
    })
  let assert [head, ..] = rows
  let size = string.length(head)
  PixelGrid(items:, size:)
}

type Rule {
  Rule(when: PixelGrid, then: PixelGrid)
}

type Rules {
  Rules(size2: List(Rule), size3: List(Rule))
}

fn parse_rules(input: String) -> Rules {
  let #(size2, size3) =
    input
    |> string.split(on: "\n")
    |> list.map(fn(rule) {
      let assert Ok(#(when, then)) = string.split_once(rule, on: " => ")
      let when = parse_grid(when)
      let then = parse_grid(then)
      // Enhancement rules always increase the subgrid dimension by one:
      // * 2×2 -> 3×3
      // * 3×3 -> 4×4
      //
      // Any other combination indicates invalid puzzle input.
      case when.size, then.size {
        2, 3 -> Nil
        3, 4 -> Nil
        _, _ -> panic as "unexpected enhancement rule"
      }
      Rule(when:, then:)
    })
    |> list.partition(fn(rule) { rule.when.size == 2 })
  Rules(size2:, size3:)
}

fn match_pattern_at(when: PixelGrid, size: Int, row: Int, col: Int, pixel_at) {
  case row < size, col < size {
    False, _ -> True
    True, False -> match_pattern_at(when, size, row + 1, 0, pixel_at)
    _, _ -> {
      let equals = set.contains(when.items, #(row, col)) == pixel_at(row, col)
      equals && match_pattern_at(when, size, row, col + 1, pixel_at)
    }
  }
}

/// Checks whether a rule matches a subgrid at the given offset.
///
/// Rules may match under any rotation or reflection symmetry.
fn rule_matches_at(rule: Rule, grid: PixelGrid, offy: Int, offx: Int) {
  let size = rule.when.size
  // All eight square symmetries:
  //
  // * identity
  // * horizontal flip
  // * vertical flip
  // * 180° rotation
  // * transpose
  // * transpose + flip variants
  [
    fn(y, x) { #(offy + y, offx + x) },
    fn(y, x) { #(offy + y, offx + { size - x - 1 }) },
    fn(y, x) { #(offy + { size - y - 1 }, offx + x) },
    fn(y, x) { #(offy + { size - y - 1 }, offx + { size - x - 1 }) },
    fn(y, x) { #(offy + x, offx + y) },
    fn(y, x) { #(offy + x, offx + { size - y - 1 }) },
    fn(y, x) { #(offy + { size - x - 1 }, offx + y) },
    fn(y, x) { #(offy + { size - x - 1 }, offx + { size - y - 1 }) },
  ]
  |> list.any(fn(getter) {
    match_pattern_at(rule.when, size, 0, 0, fn(y, x) {
      set.contains(grid.items, getter(y, x))
    })
  })
}

fn find_rule_for_subgrid(
  rules: List(Rule),
  grid: PixelGrid,
  offy: Int,
  offx: Int,
) {
  list.find(rules, fn(rule) { rule_matches_at(rule, grid, offy, offx) })
}

/// Traverses equally sized subgrids from left-to-right, top-to-bottom,
/// threading an accumulator through each chunk.
fn fold_grid_chunks(
  grid: PixelGrid,
  row: Int,
  col: Int,
  increment: Int,
  acc: a,
  fold: fn(a, Int, Int, Int) -> Result(a, Nil),
) {
  case row < grid.size, col < grid.size {
    False, _ -> Ok(acc)
    _, False -> {
      fold_grid_chunks(grid, row + increment, 0, increment, acc, fold)
    }
    _, _ -> {
      case fold(acc, increment, row, col) {
        Ok(next) ->
          fold_grid_chunks(grid, row, col + increment, increment, next, fold)
        Error(_) -> Error(Nil)
      }
    }
  }
}

/// Splits the grid into:
///
/// * 2×2 chunks when divisible by 2
/// * 3×3 chunks when divisible by 3
///
/// As defined by the enhancement rules.
fn fold_subgrids(
  grid: PixelGrid,
  acc: a,
  fold: fn(a, Int, Int, Int) -> Result(a, Nil),
) -> Result(a, Nil) {
  case grid.size % 2, grid.size % 3 {
    0, _ -> fold_grid_chunks(grid, 0, 0, 2, acc, fold)
    _, 0 -> fold_grid_chunks(grid, 0, 0, 3, acc, fold)
    _, _ -> panic as "unexpected grid size"
  }
}

/// Applies enhancement rules to every subgrid and assembles
/// the resulting larger grid.
fn enhance_once(rules: Rules, grid: PixelGrid) {
  let items =
    fold_subgrids(grid, set.new(), fn(items, size, offy, offx) {
      let rule = case size {
        2 -> find_rule_for_subgrid(rules.size2, grid, offy, offx)
        3 -> find_rule_for_subgrid(rules.size3, grid, offy, offx)
        _ -> panic as "unreachable"
      }
      case rule {
        Ok(rule) -> {
          // Map the source subgrid position into its destination position.
          // Enhanced subgrids grow from N×N to (N+1)×(N+1).
          let relative_row = { offy / size } * { size + 1 }
          let relative_col = { offx / size } * { size + 1 }
          set.fold(rule.then.items, items, fn(items, key) {
            let #(y, x) = key
            set.insert(items, #(relative_row + y, relative_col + x))
          })
          |> Ok()
        }
        Error(_) -> Error(Nil)
      }
    })
  case items {
    Ok(items) -> {
      let size = case grid.size % 2, grid.size % 3 {
        0, _ -> { grid.size / 2 } * 3
        _, 0 -> { grid.size / 3 } * 4
        _, _ -> panic as "unreachable"
      }
      Ok(PixelGrid(items:, size:))
    }
    Error(Nil) -> Error(Nil)
  }
}

/// Repeatedly enhances the grid for the requested number of iterations.
fn enhance_iterations(iterations: Int, rules: Rules, grid: PixelGrid) {
  case iterations > 0 {
    True -> {
      case enhance_once(rules, grid) {
        Error(Nil) -> grid
        Ok(next) -> {
          enhance_iterations(iterations - 1, rules, next)
        }
      }
    }
    False -> grid
  }
}

fn count_active_pixels(grid: PixelGrid) {
  set.size(grid.items)
}

// Initial grid state defined by the puzzle.
const initial_program = ".#./..#/###"

pub fn part1(s: String) {
  let rules = parse_rules(s)
  let grid = parse_grid(initial_program)
  enhance_iterations(5, rules, grid)
  |> count_active_pixels()
}

pub fn part2(s: String) {
  let rules = parse_rules(s)
  let grid = parse_grid(initial_program)
  enhance_iterations(18, rules, grid)
  |> count_active_pixels()
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-21.txt") |> string.trim()
  io.println("Day 21")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
