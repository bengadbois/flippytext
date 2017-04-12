package flippytext

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const defaultTickerTime = time.Millisecond * 10
const defaultTickerCount = 10
const defaultRandomChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type FlippyText struct {
	TickerTime  time.Duration
	TickerCount int
	RandomChars string
}

func (t *FlippyText) Write(word string) error {
	if word == "" {
		return nil
	}
	if t.RandomChars == "" {
		return errors.New("random char is empty")
	}
	cleaned := strings.Replace(word, "\r", "", -1) //strip out "\r"s
	parts := strings.Split(cleaned, "\n")
	for _, part := range parts {
		for c := 0; c < len(part); c++ {
			for i := 0; i < t.TickerCount; i++ {
				time.Sleep(t.TickerTime)
				r := rand.Intn(len(t.RandomChars) - 1)
				fmt.Printf("\r%s%s", part[:c], t.RandomChars[r:r+1])
			}
		}
		fmt.Print("\r" + part + "\n")
	}
	return nil
}

func New() *FlippyText {
	return &FlippyText{
		TickerTime:  defaultTickerTime,
		TickerCount: defaultTickerCount,
		RandomChars: defaultRandomChars,
	}
}
