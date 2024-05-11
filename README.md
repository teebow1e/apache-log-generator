# apache-log-generator
Golang application to generate fake Apache HTTPD/NGINX log entry for analytics purpose.

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

## Todo
- [x] Better flag handling
- [x] Handle probability of random choices (for example, 200 or GET should be the most popular)
- [ ] Refactor code
- [ ] add support for ipv6?