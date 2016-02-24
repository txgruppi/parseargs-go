[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/nproc/parseargs-go)
[![Codeship](https://img.shields.io/codeship/173b62f0-bcc9-0133-0239-6e8926ac3d5c/master.svg?style=flat-square)](https://codeship.com/projects/136367)
[![Codecov](https://img.shields.io/codecov/c/github/nproc/parseargs-go/master.svg?style=flat-square)](https://codecov.io/github/nproc/parseargs-go)
[![Go Report Card](https://img.shields.io/badge/go_report-A+-brightgreen.svg?style=flat-square)](https://goreportcard.com/report/github.com/nproc/parseargs-go)

# `parseargs-go`

This is a port of the [parserargs.js](https://github.com/txgruppi/parseargs.js) project to [Go](https://golang.org).

What about parsing arguments allowing quotes in them? But beware that this library will not parse flags (-- and -), flags will be returned as simple strings.

## Installation

`go get -u github.com/nproc/parseargs-go`

## Example

```go
package main

import (
  "fmt"
  "log"

  "github.com/nproc/parseargs-go"
)

func main() {
  setInRedis := `set name "Put your name here"`
  parsed, err := parseargs.Parse(setInRedis)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%#v\n", parsed) // []string{"set", "name", "Put your name here"}
}
```

## License

MIT
