# apache-log-generator
Golang application to generate fake Apache HTTPD/NGINX log entry for analytics purpose.

## Demo
![demo](assets/demo.gif)

## Installation
You can download the tool directly from the Release section or build from source.

> You need to have Go installed on your machine to build.

Simply run `go build` on the root directory of the repository.

## How to use
```
Usage of ./apache-log-generator:
  -ipv6
        enable ipv6 in log
  -num int
        number of lines to generate (default 100)
  -output string
        the output file name (default "access_log_11_05_24_09h18.log")
  -size int
        size of the output file in MB
  -sleep int
        number of seconds between two log entries
```

> Please note that you need to for boolean flag like -ipv6, you need to specify it like this: `-ipv6=true`

## Benchmark
> This program is being tested on WSL2 with Intel Core i7-1165G7, 16gb of RAM, Golang version 1.22.0.
1. Writing by lines
- 1000 lines in ~0.04 secs
- 10000 lines in ~0.2 secs

2. Writing by size
- 100 MB of logs (490k lines) in ~10s
- 500 MB of logs (2.4m lines) in ~50s
- 2 GB of logs (9.9m lines) in ~200s

## Todo
- [x] Better flag handling
- [x] Handle probability of random choices (for example, 200 or GET should be the most popular)
- [ ] Refactor code
- [ ] add support for ipv6?

## Credit
- This project was inspired by [kiritbasu/Fake-Apache-Log-Generator](kiritbasu/Fake-Apache-Log-Generator).  
- All the wordlist in `data/` is taken from [SecLists](https://github.com/danielmiessler/SecLists) by Daniel Miessler.
- The `weightedRandom` function was inspired by Jason McVetta's [randutil](https://github.com/jmcvetta/randutil) tool.
