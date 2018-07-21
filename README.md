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

#### Structs

Consider the following (simplified) struct and method:

```go
type File struct {}

func (f *File) Read(n, timeout int) ([]byte, error) { ... }

file := &File{}
```

To begin introspection, use the `Struct()` function:

```go
reflectr.Struct(file)
```

This function returns a special object with a number of methods that can be used for introspection. Most of these methods return the object, allowing method calls to be chained.

#### Methods

To confirm that a method exists, pass its name to `Method`:

```go
reflectr.Struct(file).Method("Read")
```

To confirm the types for a method, use `Param` (for a single parameter) or `Params` (for all parameters):

```go
reflectr.Struct(file).Method("Read").Param(0, int(0))
reflectr.Struct(file).Method("Read").Params(int(0), int(0))
```

The first line ensures that the method's first parameter is an `int`. The second line ensures that the method takes two `int` parameters.

To verify return types, use the `Return` and `Returns` methods.

#### Interfaces

In order to use an interface type for introspection, the `Interface` helper function may be used:

```go
reflectr.Interface((*io.Reader)(nil))
```

This function returns a special value that can be passed to methods (like `Param`). An `error` can be specified using `reflectr.ErrorType`:

```go
reflectr.Struct(file).Method("Read").Return(1, reflectr.ErrorType)
```

#### Errors

If an error occurs in a method, it is stored internally and can be accessed via the `Error()` method. As soon as an error occurs, remaining method calls will terminate immediately to prevent unnecessary processing.

For example:

```go
reflectr.Struct(file).Method("Read").Param(0, int(0)).Error()
```

This will return `nil` if the method exists and accepts an `int` as its first parameter; it will return an error if it does not.
