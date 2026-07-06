import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/set.{type Set}
import gleam/string

import utils

/// A coordinate in the simulation grid.
type Coord =
  #(Int, Int)

/// Coordinates currently marked as infected.
type InfectedNodes =
  Set(Coord)

/// State of a node in the evolved virus simulation.
type NodeState {
  Weakened
  Infected
  Flagged
}

/// Sparse grid of node states for the evolved simulation.
type NodeStates =
  Dict(Coord, NodeState)

/// Cardinal movement direction of the virus carrier.
type Direction {
  Up
  Down
  Left
  Right
}

fn turn_right(direction: Direction) -> Direction {
  case direction {
    Up -> Right
    Down -> Left
    Right -> Down
    Left -> Up
  }
}

fn turn_left(direction: Direction) -> Direction {
  case direction {
    Up -> Left
    Down -> Right
    Right -> Up
    Left -> Down
  }
}

fn reverse_direction(direction: Direction) -> Direction {
  case direction {
    Up -> Down
    Down -> Up
    Right -> Left
    Left -> Right
  }
}

fn move_forward(coord, direction) {
  let #(row, col) = coord
  case direction {
    Up -> #(row - 1, col)
    Down -> #(row + 1, col)
    Left -> #(row, col - 1)
    Right -> #(row, col + 1)
  }
}

const node_is_infected = set.contains

const mark_infected = set.insert

const clean_node = set.delete

/// Executes one burst in the basic virus simulation.
///
/// Returns whether a new infection occurred and the updated:
/// - direction
/// - carrier position
/// - infected node set
fn simulate_burst(
  direction: Direction,
  coord: Coord,
  infected_nodes: InfectedNodes,
) -> #(Bool, #(Direction, Coord, InfectedNodes)) {
  case node_is_infected(infected_nodes, coord) {
    True -> {
      let direction = turn_right(direction)
      let infected_nodes = clean_node(infected_nodes, coord)
      let coord = move_forward(coord, direction)
      #(False, #(direction, coord, infected_nodes))
    }
    False -> {
      let direction = turn_left(direction)
      let infected_nodes = mark_infected(infected_nodes, coord)
      let coord = move_forward(coord, direction)
      #(True, #(direction, coord, infected_nodes))
    }
  }
}

/// Executes one burst in the evolved four-state simulation.
///
/// Node transitions:
/// Clean -> Weakened
/// Weakened -> Infected
/// Infected -> Flagged
/// Flagged -> Clean
///
/// Returns whether a node became infected during this burst,
/// along with the updated simulation state.
fn simulate_evolved_burst(
  direction: Direction,
  coord: Coord,
  node_states: NodeStates,
) {
  case dict.get(node_states, coord) {
    Error(_) -> {
      // Clean -> Weakended
      let direction = turn_left(direction)
      let node_states = dict.insert(node_states, coord, Weakened)
      let coord = move_forward(coord, direction)
      #(False, #(direction, coord, node_states))
    }
    Ok(Weakened) -> {
      // Weakened -> Infected
      let node_states = dict.insert(node_states, coord, Infected)
      let coord = move_forward(coord, direction)
      #(True, #(direction, coord, node_states))
    }
    Ok(Infected) -> {
      // Infected -> Flagged
      let direction = turn_right(direction)
      let node_states = dict.insert(node_states, coord, Flagged)
      let coord = move_forward(coord, direction)
      #(False, #(direction, coord, node_states))
    }
    Ok(Flagged) -> {
      // Flagged -> Clean
      let direction = reverse_direction(direction)
      let node_states = dict.delete(node_states, coord)
      let coord = move_forward(coord, direction)
      #(False, #(direction, coord, node_states))
    }
  }
}

fn parse(input: String) -> #(Int, Int, InfectedNodes) {
  let lines = string.split(input, on: "\n")
  let grid =
    lines
    |> list.index_fold(set.new(), fn(grid, line, row) {
      string.to_graphemes(line)
      |> list.index_fold(grid, fn(grid, cell, col) {
        case cell {
          "#" -> set.insert(grid, #(row, col))
          _ -> grid
        }
      })
    })
  let rows = list.length(lines)
  let cols = case lines {
    [head, ..] -> string.length(head)
    _ -> 0
  }
  #(rows, cols, grid)
}

fn run_burst(
  iterations: Int,
  count: Int,
  direction: Direction,
  coord: Coord,
  node_states: InfectedNodes,
) {
  case iterations > 0 {
    False -> count
    True -> {
      let #(caused_infection, #(direction, coord, node_states)) =
        simulate_burst(direction, coord, node_states)
      let count = case caused_infection {
        True -> count + 1
        False -> count
      }
      run_burst(iterations - 1, count, direction, coord, node_states)
    }
  }
}

fn run_evolved_burst(
  iterations: Int,
  count: Int,
  direction: Direction,
  coord: Coord,
  node_states: NodeStates,
) {
  case iterations > 0 {
    False -> count
    True -> {
      let #(caused_infection, #(direction, coord, node_states)) =
        simulate_evolved_burst(direction, coord, node_states)
      let count = case caused_infection {
        True -> count + 1
        False -> count
      }
      run_evolved_burst(iterations - 1, count, direction, coord, node_states)
    }
  }
}

pub fn part1(s: String) {
  let #(rows, cols, node_states) = parse(s)
  let coord = #(rows / 2, cols / 2)
  let direction = Up
  run_burst(10_000, 0, direction, coord, node_states)
}

pub fn part2(s: String) {
  let #(rows, cols, node_states) = parse(s)
  let coord = #(rows / 2, cols / 2)
  let direction = Up
  let node_states =
    set.fold(node_states, dict.new(), fn(grid, coord) {
      dict.insert(grid, coord, Infected)
    })
  run_evolved_burst(10_000_000, 0, direction, coord, node_states)
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-22.txt") |> string.trim()
  io.println("Day 22")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
