# My AOC Repository

* [Advent of Code 2015](https://adventofcode.com/2015)
* [Libs](https://stackoverflow.com/questions/57756927/rust-modules-confusion-when-there-is-main-rs-and-lib-rs)

# Running

## Only one

```bash
cargo run -- <day>
```

## All (Powershell)

```powershell
cargo build
1..25 | % { ./target/debug/aoc2015.exe $_; }
```

## All (Bash)

```bash
cargo build
for i in $(seq 1 25); do ./target/debug/aoc2015.exe $i; done;
```
