# go-md2

Package **md2** implements the MD2 cryptographic hash algorithm as per [IETF RFC-1319](https://datatracker.ietf.org/doc/html/rfc1319), for the Go programming language.

Note that MD2 is no longer considered _secure_.
However, it may still be useful for some non-secure (and historical) use-cases.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-md2

[![GoDoc](https://godoc.org/github.com/reiver/go-md2?status.svg)](https://godoc.org/github.com/reiver/go-md2)

## Examples

Here is a simple example of calculating the MD2 digest of data contained in a `[]byte`:

```go
import "github.com/reiver/go-md2"

// ...

var data []byte = []byte("Hello world!")

digest := md2.Sum(data)
```

Here is a more complex example of putting data into the an MD2 hasher in multiple pieces.

```go
import "github.com/reiver/go-md2"

// ...

hasher := md2.New()

// This is equivalent to writing:
//
//	[]byte("once+twice-thrice_fource")
hasher.Write([]byte("once"))
hasher.Write([]byte("+"))
hasher.Write([]byte("twice"))
hasher.Write([]byte("-"))
hasher.Write([]byte("thrice"))
hasher.Write([]byte("_"))
hasher.Write([]byte("fource"))

digest := hasher.Sum(nil)
```

## Import

To import package **md2** use `import` code like the following:
```
import "github.com/reiver/go-md2"
```

## Installation

To install package **md2** do the following:
```
GOPROXY=direct go get github.com/reiver/go-md2
```

## Author

Package **md2** was written by [Charles Iliya Krempeaux](http://reiver.link)
