import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

pub type Tile {
  Path
  Letter(String)
}

type Position =
  #(Int, Int)

pub type Grid {
  Grid(tiles: Dict(Position, Tile), rows: Int, cols: Int)
}

fn parse_loop(chars, row, col, tiles) {
  case chars {
    [] -> tiles
    [tile, ..chars] -> {
      case tile {
        "\n" -> parse_loop(chars, row + 1, 0, tiles)
        " " -> parse_loop(chars, row, col + 1, tiles)
        "|" | "+" | "-" -> {
          let tiles = dict.insert(tiles, #(row, col), Path)
          parse_loop(chars, row, col + 1, tiles)
        }
        ch -> {
          let tiles = dict.insert(tiles, #(row, col), Letter(ch))
          parse_loop(chars, row, col + 1, tiles)
        }
      }
    }
  }
}

fn parse(s: String) -> Grid {
  let chars = string.to_graphemes(s)
  let tiles = parse_loop(chars, 0, 0, dict.new())
  let #(rows, cols) =
    dict.fold(tiles, #(0, 0), fn(bounds, key, _) {
      let #(row, col) = key
      let #(rows, cols) = bounds
      #(int.max(row, rows), int.max(col, cols))
    })
  Grid(tiles, rows + 1, cols + 1)
}

fn grid_debug_loop(grid: Grid, row: Int, col: Int, acc) -> String {
  case row >= 0, col >= 0 {
    True, True -> {
      let tile = case dict.get(grid.tiles, #(row, col)) {
        Ok(Path) -> "~"
        Ok(Letter(key)) -> key
        Error(_) -> " "
      }
      grid_debug_loop(grid, row, col - 1, [tile, ..acc])
    }
    True, False if row > 0 ->
      grid_debug_loop(grid, row - 1, grid.cols - 1, ["\n", ..acc])
    _, _ -> string.concat(acc)
  }
}

/// Debug helper to print the parsed grid.
pub fn grid_debug(grid: Grid) {
  grid_debug_loop(grid, grid.rows - 1, grid.cols - 1, [])
}

fn grid_start_loop(grid: Grid, col: Int) {
  case col < grid.cols {
    False -> Error(Nil)
    True -> {
      case dict.get(grid.tiles, #(0, col)) {
        Ok(_) -> Ok(#(0, col))
        Error(_) -> grid_start_loop(grid, col + 1)
      }
    }
  }
}

/// Finds the starting position on the first row.
fn grid_start(grid: Grid) {
  grid_start_loop(grid, 0)
}

/// Applies a directional increment to a position.
fn step(pos: Position, direction: Position) {
  let #(row, col) = pos
  let #(row_inc, col_inc) = direction
  #(row + row_inc, col + col_inc)
}

/// Determines the next valid direction at intersections.
fn find_turn_direction(grid: Grid, pos: Position, direction: Position) {
  case direction {
    #(0, 0) -> [#(0, -1), #(0, 1), #(1, 0), #(-1, 0)]
    #(1, 0) | #(-1, 0) -> [#(0, -1), #(0, 1)]
    #(0, 1) | #(0, -1) -> [#(1, 0), #(-1, 0)]
    _ -> []
  }
  |> list.find_map(fn(next_inc) {
    let next_pos = step(pos, next_inc)
    let #(next_row, next_col) = next_pos
    case
      0 <= next_row
      && next_row < grid.rows
      && 0 <= next_col
      && next_col < grid.cols
    {
      False -> Error(Nil)
      True -> {
        case dict.get(grid.tiles, next_pos) {
          Ok(_) -> Ok(next_inc)
          Error(_) -> Error(Nil)
        }
      }
    }
  })
}

fn walk_path_loop(grid: Grid, pos, direction, letters, count) {
  let next_pos = step(pos, direction)
  case dict.get(grid.tiles, next_pos) {
    Ok(Path) -> walk_path_loop(grid, next_pos, direction, letters, count + 1)
    Ok(Letter(letter)) ->
      walk_path_loop(grid, next_pos, direction, [letter, ..letters], count + 1)
    Error(_) -> {
      case find_turn_direction(grid, pos, direction) {
        Ok(next_inc) -> walk_path_loop(grid, pos, next_inc, letters, count)
        Error(_) -> {
          let letters =
            letters
            |> list.reverse()
            |> string.concat()
          #(letters, count + 1)
        }
      }
    }
  }
}

/// Traverses the path collecting letters and counting steps.
fn walk_path(grid: Grid) {
  let assert Ok(start) = grid_start(grid)
  let assert Ok(dir) = find_turn_direction(grid, start, #(0, 0))
  walk_path_loop(grid, start, dir, [], 0)
}

fn collect_letters(grid: Grid) {
  let #(letters, _) = walk_path(grid)
  letters
}

fn count_steps(grid: Grid) {
  let #(_, count) = walk_path(grid)
  count
}

pub fn part1(s: String) {
  parse(s)
  |> collect_letters()
}

pub fn part2(s: String) {
  parse(s)
  |> count_steps()
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-19.txt")
  io.println("Day 19")
  utils.time_it("Part 1", fn() { part1(input) })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
