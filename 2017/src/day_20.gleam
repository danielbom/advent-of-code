import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/string

import utils

type V3 {
  V3(x: Int, y: Int, z: Int)
}

fn add_v3(a: V3, b: V3) -> V3 {
  V3(a.x + b.x, a.y + b.y, a.z + b.z)
}

fn parse_v3(input: String) -> V3 {
  let assert Ok([x, y, z]) =
    input
    |> string.drop_start(3)
    |> string.drop_end(1)
    |> string.split(on: ",")
    |> list.try_map(fn(x) { int.parse(string.trim(x)) })
  V3(x, y, z)
}

/// Computes the Manhattan distance from the origin.
fn manhattan_distance(position: V3) -> Int {
  int.absolute_value(position.x)
  + int.absolute_value(position.y)
  + int.absolute_value(position.z)
}

type Particle {
  Particle(position: V3, velocity: V3, acceleration: V3)
}

fn parse_particles(input: String) -> List(Particle) {
  input
  |> string.split(on: "\n")
  |> list.map(fn(line) {
    let assert [position, velocity, acceleration] =
      line
      |> string.split(on: ", ")
      |> list.map(parse_v3)
    Particle(position:, velocity:, acceleration:)
  })
}

/// Advances a particle by one simulation step.
fn particle_step(particle: Particle) -> Particle {
  let Particle(position, velocity, acceleration) = particle
  let velocity = add_v3(velocity, acceleration)
  let position = add_v3(position, velocity)
  Particle(position:, velocity:, acceleration:)
}

/// Advances a particle forward by `steps` simulation ticks.
fn particle_step_times(particle: Particle, steps: Int) -> Particle {
  case steps > 0 {
    True -> particle_step_times(particle_step(particle), steps - 1)
    False -> particle
  }
}

/// Heuristically estimates how many simulation steps are
/// sufficient before particle ordering stabilizes.
fn estimate_steps_loop(
  particle: Particle,
  best_distance: Int,
  steps: Int,
) -> Int {
  let next_particle = particle_step(particle)
  let next_distance = manhattan_distance(next_particle.position)
  case
    // acceleration can still reduce distance to the origin
    { particle.position.x > 0 && particle.acceleration.x < 0 }
    || { particle.position.y > 0 && particle.acceleration.y < 0 }
    || { particle.position.z > 0 && particle.acceleration.z < 0 }
    || { particle.position.x < 0 && particle.acceleration.x > 0 }
    || { particle.position.y < 0 && particle.acceleration.y > 0 }
    || { particle.position.z < 0 && particle.acceleration.z > 0 }
    // particle is still approaching the origin
    || next_distance < best_distance
  {
    False -> steps
    True -> {
      let min_distance = int.min(next_distance, best_distance)
      estimate_steps_loop(next_particle, min_distance, steps + 1)
    }
  }
}

/// Estimates the number of steps before the particle stops
/// approaching the origin.
fn estimate_steps(particle: Particle) -> Int {
  estimate_steps_loop(particle, manhattan_distance(particle.position), 0)
}

/// Returns the largest estimated convergence duration
/// among all particles.
fn max_estimated_steps(particles: List(Particle)) -> Int {
  list.fold(particles, 0, fn(max_steps, particle) {
    let steps = estimate_steps(particle)
    int.max(max_steps, steps)
  })
}

/// Returns the index of the particle expected to remain
/// closest to the origin in the long term.
fn closest_particle_in_long_term(particles: List(Particle)) -> Int {
  let max_steps = max_estimated_steps(particles)
  let #(index, _) =
    list.index_fold(particles, #(-1, -1), fn(current_best, particle, index) {
      let #(_, min_distance) = current_best
      let particle = particle_step_times(particle, max_steps)
      let distance = manhattan_distance(particle.position)
      case min_distance == -1 || min_distance > distance {
        True -> #(index, distance)
        False -> current_best
      }
    })
  index
}

pub fn part1(s: String) -> Int {
  parse_particles(s)
  |> closest_particle_in_long_term()
}

/// Removes all particles that share the same position.
fn filter_collided_particles(particles: List(Particle)) -> List(Particle) {
  list.fold(particles, dict.new(), fn(acc, particle) {
    case dict.get(acc, particle.position) {
      Ok([]) -> acc
      Ok(_) -> dict.insert(acc, particle.position, [])
      Error(_) -> dict.insert(acc, particle.position, [particle])
    }
  })
  |> dict.fold([], fn(acc, _, values) {
    case values {
      [particle] -> [particle, ..acc]
      _ -> acc
    }
  })
}

/// Simulates particle movement and collision removal
/// for a fixed number of steps.
fn simulate_particles(particles: List(Particle), steps: Int) -> List(Particle) {
  case steps > 0 {
    True ->
      particles
      |> list.map(particle_step)
      |> filter_collided_particles()
      |> simulate_particles(steps - 1)
    False -> particles
  }
}

/// Estimates the number of particles remaining after
/// all relevant collisions have occurred.
pub fn part2(s: String) -> Int {
  let particles = parse_particles(s)
  let max_steps = max_estimated_steps(particles)
  simulate_particles(particles, max_steps)
  |> list.length()
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-20.txt") |> string.trim()
  io.println("Day 20")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
