## go-anydate

Attempts to detect the correct Go time layout format for parsing a given time string.
It analyzes the input time string and returns a corresponding time format layout that can be used with `time.Parse()` or `time.Format()`.

[![go-anydate Release](https://img.shields.io/github/release/nodivbyzero/go-anydate.svg)](https://github.com/nodivbyzero/go-anydate/releases)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)


## Install

```
go get -u github.com/nodivbyzero/go-anydate
```



## Usage

Calculating time based on current time

```go
import "github.com/nodivbyzero/go-anydate"

str := "2024-11-14 22:43:57"
got, err := DetectFormat(str) // got = 2006-01-02 15:04:05

date, err := time.Parse(got, str)

```

## License

Released under the [MIT License](http://www.opensource.org/licenses/MIT).
