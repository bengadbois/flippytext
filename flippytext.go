package flippytext

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

const defaultTickerTime = time.Millisecond * 10
const defaultTickerCount = 10
const defaultRandomChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type FlippyText struct {
	// How long to pause between each characer flipping
	TickerTime time.Duration
	// How many times each character should flip through before resolving
	TickerCount int
	// The list of characters to use while flipping
	RandomChars string
	// Where to write the output
	Output io.Writer
}

// Flip through the characters of word, printing to stdout
func (t *FlippyText) Write(word string) error {
	if word == "" {
		return nil
	}
	if t.RandomChars == "" {
		return errors.New("random char is empty")
	}
	if t.Output == nil {
		return errors.New("nil output for writing")
	}
	_, err := fmt.Fprint(t.Output, "")
	if err != nil {
		return errors.New("unable to write to output:" + err.Error())
	}
	cleaned := strings.Replace(word, "\r", "", -1) //strip out "\r"s
	parts := strings.Split(cleaned, "\n")
	for _, part := range parts {
		for c := 0; c < len(part); c++ {
			for i := 0; i < t.TickerCount; i++ {
				time.Sleep(t.TickerTime)
				r := rand.Intn(len(t.RandomChars) - 1)
				fmt.Fprintf(t.Output, "\r%s%s", part[:c], t.RandomChars[r:r+1])
			}
		}
		fmt.Fprintf(t.Output, "\r"+part+"\n")
	}
	return nil
}

// Create a new FlippyText
func New() *FlippyText {
	return &FlippyText{
		TickerTime:  defaultTickerTime,
		TickerCount: defaultTickerCount,
		RandomChars: defaultRandomChars,
		Output:      os.Stdout,
	}
}
