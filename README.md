# FlippyText [![Build Status](https://travis-ci.org/bengadbois/flippytext.svg?branch=master)](https://travis-ci.org/bengadbois/flippytext) [![Coverage Status](https://coveralls.io/repos/github/bengadbois/flippytext/badge.svg?branch=master)](https://coveralls.io/github/bengadbois/flippytext?branch=master) [![GoDoc](https://godoc.org/github.com/bengadbois/flippytext?status.svg)](https://godoc.org/github.com/bengadbois/flippytext)

FlippyText is a Go library for printing animated text one character at a time.

![Screencap](screencap.gif)

## Install

```
go get github.com/bengadbois/flippytext
```

## Example 

```go
import "github.com/bengadbois/flippytext"

func main() {
		flippytext.New().Write("hello world")
}
```
