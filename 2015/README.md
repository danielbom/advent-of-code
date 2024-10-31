# My AOC Repository

* [Advent of Code 2015](https://adventofcode.com/2015)
* [Libs](https://stackoverflow.com/questions/57756927/rust-modules-confusion-when-there-is-main-rs-and-lib-rs)

# Build, Run, Test

## Run

```bash
cargo run -- <day>  // Single day
cargo run -- 0      // All days
```

## Test

```bash
cargo test --package aoc2015 --bin aoc2015 -- day19::tests --show-output --nocapture
```

## Format

```bash
cargo fmt
```

## Lint

```bash
cargo clippy
```
