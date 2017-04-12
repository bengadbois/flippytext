# FlippyText

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
