import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

const use_formula = True

type Scanner {
  Scanner(depth: Int, range: Int)
}

// Parse firewall scanner definitions
fn parse(s: String) -> List(Scanner) {
  string.trim(s)
  |> string.split(on: "\n")
  |> list.map(fn(line) {
    let assert Ok(#(depth, range)) = string.split_once(line, on: ": ")
    let assert Ok(depth) = int.parse(depth)
    let assert Ok(range) = int.parse(range)
    Scanner(depth:, range:)
  })
}

fn simulate_scanner(range: Int, position: Int, steps: Int, forward: Bool) {
  case steps > 0, forward {
    False, _ -> position
    _, True -> {
      let advance = int.min(steps, range)
      let position = position + advance
      let forward = position < range - 1
      simulate_scanner(range, position, steps - advance, forward)
    }
    _, False -> {
      let advance = int.min(steps, position)
      let position = position - advance
      let forward = position == 0
      simulate_scanner(range, position, steps - advance, forward)
    }
  }
}

/// Compute the scanner position by simulating its movement.
///
/// The scanner's path repeats every (range - 1) * 2 steps,
/// so only the position within the current cycle needs to
/// be simulated.
fn scanner_position_simulated(scanner: Scanner, delay: Int) {
  let cycle_length = { scanner.range - 1 } * 2
  let elapsed = scanner.depth + delay
  let cycle_step = elapsed % cycle_length
  simulate_scanner(scanner.range, 0, cycle_step, True)
}

/// Compute the scanner position using a triangular-wave formula.
///
/// The scanner moves between 0 and range - 1, repeating every
/// (range - 1) * 2 steps. The position is calculated directly
/// without simulating the movement.
fn scanner_position_formula(scanner: Scanner, delay: Int) {
  let n = scanner.range - 1
  let i = { scanner.depth + delay } % { n * 2 }
  n - int.absolute_value(i - n)
}

/// Return the scanner position when the packet reaches
/// the scanner after waiting `delay` picoseconds.
fn scanner_position(scanner: Scanner, delay: Int) {
  case use_formula {
    True -> scanner_position_formula(scanner, delay)
    False -> scanner_position_simulated(scanner, delay)
  }
}

/// Visit each scanner and accumulate a result.
///
/// The traversal can terminate early by returning `Error`.
fn traverse(
  scanners: List(Scanner),
  delay: Int,
  initial acc: a,
  visit visit: fn(a, Scanner, Int) -> Result(a, a),
) -> a {
  case scanners {
    [] -> acc
    [scanner, ..scanners] -> {
      let position = scanner_position(scanner, delay)
      case visit(acc, scanner, position) {
        Ok(curr) -> traverse(scanners, delay, curr, visit)
        Error(result) -> result
      }
    }
  }
}

// Compute the total severity of traversing the firewall
// without any delay.
fn trip_severity(scanners: List(Scanner)) -> Int {
  traverse(scanners, 0, 0, fn(severity, scanner, position) {
    case position == 0 {
      True -> Ok(severity + scanner.range * scanner.depth)
      False -> Ok(severity)
    }
  })
}

// Return True if traversing the firewall with the given
// delay causes the packet to be caught by any scanner.
fn is_caught(scanners: List(Scanner), delay: Int) -> Bool {
  traverse(scanners, delay, False, fn(_, _, position) {
    case position == 0 {
      True -> Error(True)
      False -> Ok(False)
    }
  })
}

pub fn part1(s: String) {
  trip_severity(parse(s))
}

fn find_safe_delay_loop(scanners: List(Scanner), delay: Int) {
  case is_caught(scanners, delay) {
    False -> delay
    True -> find_safe_delay_loop(scanners, delay + 1)
  }
}

/// Find the smallest delay that allows the packet to pass
/// through the firewall without being caught.
fn find_safe_delay(scanners: List(Scanner)) {
  find_safe_delay_loop(scanners, 0)
}

pub fn part2(s: String) {
  find_safe_delay(parse(s))
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-13.txt") |> string.trim()
  io.println("Day 13")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
