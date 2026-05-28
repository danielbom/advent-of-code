import gleam/int
import gleam/io
import gleam/list
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
fn get(hash: Hash, index: Int) -> Int {
  let ix = index % hash.size
  let assert Ok(value) = iv.get(hash.items, ix)
  value
}

/// Returns a new hash with the value written at the circular index.
fn put(hash: Hash, index: Int, value: Int) -> Hash {
  let ix = index % hash.size
  Hash(..hash, items: iv.try_set(hash.items, ix, value))
}

/// Reverses a range of values in the circular hash buffer.
fn reverse_range(hash: Hash, left: Int, right: Int) -> Hash {
  case left < right {
    True -> {
      let fst = get(hash, left)
      let snd = get(hash, right)
      case fst == snd {
        True -> hash
        False -> hash |> put(left, snd) |> put(right, fst)
      }
      |> reverse_range(left + 1, right - 1)
    }
    False -> hash
  }
}

/// Parses a comma-separated list of integers.
fn parse_lengths(s: String) -> List(Int) {
  string.trim(s)
  |> string.split(on: ",")
  |> list.map(fn(value) {
    let assert Ok(value) = int.parse(value)
    value
  })
}

/// Converts the input string into UTF-8 codepoint integers.
fn parse_ascii_codes(s: String) -> List(Int) {
  string.to_utf_codepoints(s)
  |> list.map(string.utf_codepoint_to_int)
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

/// Computes the product of the first two hash values after one round.
pub fn knot_hash_check(lengths: List(Int), size: Int) {
  let hash = Hash(items: iv.initialise(size, fn(x) { x }), size: size)
  let result = knot_hash(1, lengths, hash)
  let fisrt = get(result, 0)
  let second = get(result, 1)
  fisrt * second
}

pub fn part1(s: String) {
  let lengths = parse_lengths(s)
  knot_hash_check(lengths, standard_list_size)
}

fn xor_range(h: Hash, from: Int, to: Int, result: Int) {
  case from < to {
    True -> {
      let result = int.bitwise_exclusive_or(result, get(h, from))
      xor_range(h, from + 1, to, result)
    }
    False -> result
  }
}

fn dense_hash_loop(h: Hash, acc: List(Int), start: Int) -> Hash {
  case start < 256 {
    True -> {
      let item = xor_range(h, start, start + 16, 0)
      dense_hash_loop(h, [item, ..acc], start + 16)
    }
    False -> Hash(items: iv.from_reverse_list(acc), size: 16)
  }
}

/// Converts a 256-element sparse hash into a 16-element dense hash.
fn dense_hash(h: Hash) -> Hash {
  case h.size {
    256 -> dense_hash_loop(h, [], 0)
    16 -> h
    _ -> panic as "invalid hash size"
  }
}

/// Encodes the hash as a lowercase hexadecimal string.
fn to_hex(h: Hash) -> String {
  h.items
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

pub fn part2(s: String) {
  parse_ascii_codes(s)
  |> knot_hash_hex()
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-10.txt") |> string.trim()
  io.println("Day 10")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) })
}
