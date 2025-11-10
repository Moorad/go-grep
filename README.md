# go-grep
A simple clone of [grep](https://en.wikipedia.org/wiki/Grep) written in go

## Benchmarks
Benchmarks were performed using [hyperfine](https://github.com/sharkdp/hyperfine). The benchmark script can be found in the root `benchmark.sh` file.

There are 4 main benchmark tests. Each try to benchmark a different aspect of the program.
1. General: Realistic use case. Single small file with a pretty simple pattern
2. Large File: A single 1GB file of random characters.
3. Many Files: 500 different small files.
4. Regex: A pretty complex regex pattern.

| | go-grep | GNU grep | ripgrep |
|---|---|---|---|
| General | 2.1ms ± 0.1 | 1.2ms ± 0.1 | 1.7ms ± 0.1 |
| Large File |  7681.7ms ± 81 | 1.2ms ± 0.1 | 120.9ms ± 3.2ms |
| Many Files | NA | NA | NA |
| Regex | NA | NA | NA |

## Tasks
- [x] Initial project setup and test setup
- [x] Search a single file for exact match
- [x] Setup pipeline and auto-release
- [x] Search multiple files for exact match
- [x] Case insensitive matching (-i)
- [x] Count number of matches (-c)
- [ ] Match whole words (-w)
- [ ] Display only the matched patterns (-o)
- [ ] Regular expressions
- [ ] Glob search
- [ ] Inverting the pattern (-v)
- [ ] Search recursively in a directory