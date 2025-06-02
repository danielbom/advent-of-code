import gleam/int
import gleam/io
import gleam/list
import gleam/set
import gleam/string

import utils

pub fn passphrase_uniq_words(pass: String) -> Bool {
  let words = string.split(pass, on: " ")
  let uniqs = set.from_list(words)
  list.length(words) == set.size(uniqs)
}

pub fn part1(s: String) -> Int {
  string.split(s, on: "\n")
  |> list.filter(passphrase_uniq_words)
  |> list.length()
}

fn passphrase_no_anagrams_loop(counters: List(List(String))) -> Bool {
  case counters {
    [] -> True
    [_] -> True
    [head, ..tail] -> {
      let other = list.find(tail, fn(other) { other == head })
      case other {
        Ok(_) -> False
        Error(_) -> passphrase_no_anagrams_loop(tail)
      }
    }
  }
}

fn passphrase_no_anagrams(pass: String) -> Bool {
  string.split(pass, on: " ")
  |> list.map(fn(word: String) {
    string.to_graphemes(word) |> list.sort(string.compare)
  })
  |> passphrase_no_anagrams_loop()
}

pub fn part2(s: String) -> Int {
  string.split(s, on: "\n")
  |> list.filter(passphrase_no_anagrams)
  |> list.length()
}

pub fn solve() {
  let input = utils.read_all_file("inputs/day-04.txt") |> string.trim()
  io.println("Day 04")
  utils.time_it("Part 1", fn() { part1(input) |> int.to_string() })
  utils.time_it("Part 2", fn() { part2(input) |> int.to_string() })
}
