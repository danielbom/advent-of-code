import gleam/bit_array
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import gleam/string_tree as st

import iv.{type Array}

import utils

const standard_list_size = 256

const standard_length_suffix = [17, 31, 73, 47, 23]

type Hash {
  Hash(items: Array(Int), size: Int)
}

/// Returns the value at the circular index.
fn hash_get(hash: Hash, index: Int) -> Int {
  let ix = index % hash.size
  let assert Ok(value) = iv.get(hash.items, ix)
  value
}

/// Returns a new hash with the value written at the circular index.
fn hash_put(hash: Hash, index: Int, value: Int) -> Hash {
  let ix = index % hash.size
  Hash(..hash, items: iv.try_set(hash.items, ix, value))
}

/// Reverses a range of values in the circular hash buffer.
fn reverse_range(hash: Hash, left: Int, right: Int) -> Hash {
  case left < right {
    True -> {
      let fst = hash_get(hash, left)
      let snd = hash_get(hash, right)
      case fst == snd {
        True -> hash
        False -> hash |> hash_put(left, snd) |> hash_put(right, fst)
      }
      |> reverse_range(left + 1, right - 1)
    }
    False -> hash
  }
}

/// Executes a single Knot Hash round.
fn run_hash_round(lengths: List(Int), hash: Hash, index: Int, skip: Int) {
  case lengths {
    [] -> #(hash, index, skip)
    [length, ..lengths] -> {
      let left = index
      let right = left + length - 1
      let hash = reverse_range(hash, left, right)
      let index = { index + length + skip } % hash.size
      run_hash_round(lengths, hash, index, skip + 1)
    }
  }
}

/// Executes multiple Knot Hash rounds.
fn run_hash_rounds(
  rounds: Int,
  lengths: List(Int),
  hash: Hash,
  index: Int,
  skip: Int,
) {
  case rounds, run_hash_round(lengths, hash, index, skip) {
    0, #(hash, _, _) -> hash
    _, #(hash, index, skip) ->
      run_hash_rounds(rounds - 1, lengths, hash, index, skip)
  }
}

/// Computes the Knot Hash sparse hash after the given number of rounds.
fn knot_hash(rounds: Int, lengths: List(Int), hash: Hash) {
  run_hash_rounds(rounds - 1, lengths, hash, 0, 0)
}

fn xor_range(hash: Hash, from: Int, to: Int, result: Int) {
  case from < to {
    True -> {
      let result = int.bitwise_exclusive_or(result, hash_get(hash, from))
      xor_range(hash, from + 1, to, result)
    }
    False -> result
  }
}

fn dense_hash_loop(hash: Hash, acc: List(Int), start: Int) -> Hash {
  case start < 256 {
    True -> {
      let item = xor_range(hash, start, start + 16, 0)
      dense_hash_loop(hash, [item, ..acc], start + 16)
    }
    False -> Hash(items: iv.from_reverse_list(acc), size: 16)
  }
}

/// Converts a 256-element sparse hash into a 16-element dense hash.
fn dense_hash(hash: Hash) -> Hash {
  case hash.size {
    256 -> dense_hash_loop(hash, [], 0)
    16 -> hash
    _ -> panic as "invalid hash size"
  }
}

/// Encodes the hash as a lowercase hexadecimal string.
fn to_hex(hash: Hash) -> String {
  hash.items
  |> iv.fold(st.new(), fn(sb, x) {
    let x = int.to_base16(x)
    let x = string.pad_start(x, to: 2, with: "0")
    st.append(sb, x)
  })
  |> st.to_string()
  |> string.lowercase()
}

/// Computes the Knot Hash hexadecimal digest for the input lengths.
fn knot_hash_hex(lengths: List(Int)) {
  let size = standard_list_size
  let rounds = 64
  let hash = Hash(items: iv.initialise(size, fn(x) { x }), size: size)
  let lengths = list.append(lengths, standard_length_suffix)
  let sparse_hash = knot_hash(rounds, lengths, hash)
  let dense_hash = dense_hash(sparse_hash)
  to_hex(dense_hash)
}

/// Converts the input string into UTF-8 codepoint integers.
fn parse_ascii_codes(s: String) -> List(Int) {
  string.to_utf_codepoints(s)
  |> list.map(string.utf_codepoint_to_int)
}

fn extract_bits_loop(array: BitArray, acc: List(Bool)) {
  case array {
    <<bit:size(1), rest:bits>> -> extract_bits_loop(rest, [bit == 1, ..acc])
    <<>> -> iv.from_reverse_list(acc)
    _ -> panic as "unreachable"
  }
}

fn extract_bits(array: BitArray) {
  extract_bits_loop(array, [])
}

type Matrix(a) =
  Array(Array(a))

fn build_matrix_loop(
  acc: List(Array(Bool)),
  s: String,
  index: Int,
) -> Matrix(Bool) {
  case index < 128 {
    True -> {
      let assert Ok(row) =
        string.append(s, "-")
        |> string.append(int.to_string(index))
        |> parse_ascii_codes()
        |> knot_hash_hex()
        |> bit_array.base16_decode()
      let row = extract_bits(row)
      build_matrix_loop([row, ..acc], s, index + 1)
    }
    False -> iv.from_reverse_list(acc)
  }
}

fn build_matrix(s: String) -> Matrix(Bool) {
  build_matrix_loop([], s, 0)
}

pub fn count_set_bits(acc: Int, array: Array(Bool)) {
  iv.fold(array, acc, fn(count, set) {
    case set {
      True -> count + 1
      False -> count
    }
  })
}

/// Counts the total number of set bits (True values) in the memory grid.
fn count_used_space(memory: Matrix(Bool)) {
  iv.fold(memory, 0, count_set_bits)
}

/// Counts the total number of set bits in the memory grid.
pub fn part1(s: String) {
  string.trim(s)
  |> build_matrix()
  |> count_used_space()
}

fn matrix_get_or(matrix: Matrix(a), y: Int, x: Int, default: a) -> a {
  iv.get(matrix, y)
  |> result.try(iv.get(_, x))
  |> result.unwrap(default)
}

fn matrix_set(matrix: Matrix(a), y: Int, x: Int, value: a) -> Matrix(a) {
  case iv.get(matrix, y) {
    Ok(row) -> iv.try_set(matrix, y, iv.try_set(row, x, value))
    Error(_) -> matrix
  }
}

/// Recursively flood-fills a connected region of set bits with a group ID.
fn fill_group(
  groups: Matrix(Int),
  memory: Matrix(Bool),
  y: Int,
  x: Int,
  group: Int,
) {
  case matrix_get_or(memory, y, x, False), matrix_get_or(groups, y, x, -1) {
    True, 0 -> {
      groups
      |> matrix_set(y, x, group)
      |> fill_group(memory, y - 1, x, group)
      |> fill_group(memory, y + 1, x, group)
      |> fill_group(memory, y, x - 1, group)
      |> fill_group(memory, y, x + 1, group)
    }
    _, _ -> groups
  }
}

/// Scans the memory grid to count distinct connected regions of set bits.
fn count_used_space_regions_loop(
  memory: Matrix(Bool),
  groups: Matrix(Int),
  y: Int,
  x: Int,
  count: Int,
) {
  let #(ny, nx) = case x < 127 {
    True -> #(y, x + 1)
    False -> #(y + 1, 0)
  }
  case matrix_get_or(memory, y, x, False), matrix_get_or(groups, y, x, -1) {
    _, -1 -> count
    True, 0 -> {
      let count = count + 1
      let groups = fill_group(groups, memory, y, x, count)
      count_used_space_regions_loop(memory, groups, ny, nx, count)
    }
    _, _ -> {
      count_used_space_regions_loop(memory, groups, ny, nx, count)
    }
  }
}

/// Counts the number of distinct connected regions in the memory grid.
fn count_used_space_regions(memory: Matrix(Bool)) {
  let n = iv.size(memory)
  let row = iv.initialise(n, fn(_) { 0 })
  let groups = iv.initialise(n, fn(_) { row })
  count_used_space_regions_loop(memory, groups, 0, 0, 0)
}

/// Counts the number of distinct connected regions of set bits in the grid.
pub fn part2(s: String) {
  string.trim(s)
  |> build_matrix()
  |> count_used_space_regions()
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-14.txt") |> string.trim()
  io.println("Day 14")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
