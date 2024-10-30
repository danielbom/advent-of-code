# My AOC Repository

* [Advent of Code 2016](https://adventofcode.com/2016)

# Build, Run, Test

## Only one

```bash
go run cmd/main.go <day>
```

## All (Powershell)

```powershell
1..25 | % { go run cmd/main.go $_; }
```

## All (Bash)

```bash
for i in $(seq 1 25); do go run cmd/main.go $i; done;
```

## Test

```bash
go test internal/day01/ -v --run TestPart1
```
