[![Go Reference](https://pkg.go.dev/badge/github.com/txgruppi/parseargs-go.svg)](https://pkg.go.dev/github.com/txgruppi/parseargs-go)
[![Go Report Card](https://img.shields.io/badge/go_report-A+-brightgreen.svg?style=flat-square)](https://goreportcard.com/report/github.com/txgruppi/parseargs-go)

# `parseargs-go`

This is a port of the [parserargs.js](https://github.com/txgruppi/parseargs.js) project to [Go](https://golang.org).

What about parsing arguments allowing quotes in them? But beware that this library will not parse flags (-- and -), flags will be returned as simple strings.

## Installation

```sh
go get github.com/txgruppi/parseargs-go
```

## Example

```go
package main

import (
  "fmt"
  "log"

  "github.com/txgruppi/parseargs-go"
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

## Tests

```sh
go test ./...
```

## License

MIT