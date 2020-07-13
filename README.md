# gongram ![CI](https://github.com/rafatbiin/gongram/workflows/CI/badge.svg)
Ngram generator in Go that just works.

## Overview
gongram provides a Go package which generates Ngrams of given
gram size from given text. It has support for unicode language as well. 
It'll remove all the punctuations but keep the cases of the letters untouched.

In *addition* this library provides a [CLI](#cli) to perform the above operation via console.

## Installing

To start using gongram, install Go and run `go get`:

```sh
$ go get github.com/rafatbiin/gongram/...
```

This will retrieve the library and install the `gongram` command line utility into
your `$GOBIN` path.

## Generating Ngram

```go
package main

import (
	"fmt"
	"github.com/rafatbiin/gongram"
)

func main()  {
	text := "It’s not about ideas. It’s about making ideas happen!"
	ngrams, err:= gongram.Generate(text, 4) // ["Its not about ideas","not about ideas Its","about ideas Its about","ideas Its about making","Its about making ideas","about making ideas happen"]
	if err != nil {
		// Handle error
	}
	...
}
```

## CLI
```sh
$ gongram ngram -g 4 It’s not about ideas. It’s about making ideas happen!
["Its not about ideas","not about ideas Its","about ideas Its about","ideas Its about making","Its about making ideas","about making ideas happen"]
```