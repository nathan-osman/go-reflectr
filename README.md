## go-reflectr

[![Build Status](https://travis-ci.org/nathan-osman/go-reflectr.svg?branch=master)](https://travis-ci.org/nathan-osman/go-reflectr)
[![Coverage Status](https://coveralls.io/repos/github/nathan-osman/go-reflectr/badge.svg?branch=master)](https://coveralls.io/github/nathan-osman/go-reflectr?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/nathan-osman/go-reflectr)](https://goreportcard.com/report/github.com/nathan-osman/go-reflectr)
[![GoDoc](https://godoc.org/github.com/nathan-osman/go-reflectr?status.svg)](https://godoc.org/github.com/nathan-osman/go-reflectr)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This package simplifies the task of working with reflection in Go.

### Usage

To begin using go-reflect, import the `reflectr` package:

```go
import "github.com/nathan-osman/go-reflectr"
```

Consider the following (simplified) struct and method:

```go
type File struct {}

func (f *File) Read(n int) ([]byte, error) { ... }
```

The following would be used to verify that File contains a method matching the signature above:

```go
f := NewFile()

if err := reflectr.
    Struct(f).
    Method("Read").
    Params(0).
    Returns([]byte(nil), reflectr.ErrorType).
    Error(); err != nil {
    // handle error
}
```

The `Params` method will confirm that the `Read` method accepts a single `int` as a parameter and returns a `[]byte` and `error`.
