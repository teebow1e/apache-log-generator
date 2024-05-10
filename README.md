# apache-log-generator
Golang application to generate fake Apache HTTPD/NGINX log entry for analytics purpose.

## How to use
```
Usage of ./apache-log-generator:
  -num int
        number of lines to generate (default 100)
  -output string
        the output file name (default "access_log_11_05_24_01h56.log")
  -size int
        size of the output file in MB
  -sleep int
        number of seconds between two log entries (default 5)
```

## Todo
- Better flag handling
- Handle probability of random choices (for example, 200 or GET should be the most popular)
- Refactor code
- add support for ipv6?