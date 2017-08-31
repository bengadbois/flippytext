package flippytext

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

type invalidWriter struct {
}

func (i invalidWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("invalid")
}

func TestNew(t *testing.T) {
	ft := New()
	want := FlippyText{
		TickerTime:  defaultTickerTime,
		TickerCount: defaultTickerCount,
		RandomChars: defaultRandomChars,
		Output:      os.Stdout,
	}
	if !reflect.DeepEqual(*ft, want) {
		t.Errorf("New() == %+v wanted %+v", *ft, want)
	}
}

func TestWrite(t *testing.T) {
	cases := []struct {
		ft      *FlippyText
		word    string
		wantErr bool
	}{
		{New(), "", false},
		{New(), "hello world", false},
		{New(), "hello\r\r\rworld", false},
		{New(), "\r", false},
		{New(), "\r\r\n\r\n\r\r", false},
		{New(), "\n", false},
		{
			&FlippyText{
				TickerTime:  time.Second * 0,
				TickerCount: defaultTickerCount,
				RandomChars: defaultRandomChars,
				Output:      os.Stdout,
			}, "hello world", false,
		},
		{
			&FlippyText{
				TickerTime:  time.Second * 0,
				TickerCount: defaultTickerCount,
				RandomChars: defaultRandomChars,
				Output:      os.Stderr,
			}, "hello world", false,
		},
		{
			&FlippyText{
				TickerTime:  time.Second * 0,
				TickerCount: defaultTickerCount,
				RandomChars: defaultRandomChars,
				Output:      ioutil.Discard,
			}, "hello world", false,
		},
		{
			&FlippyText{
				TickerTime:  defaultTickerTime,
				TickerCount: defaultTickerCount,
				RandomChars: "",
				Output:      os.Stdout,
			}, "hello world", true,
		},
		{
			&FlippyText{
				TickerTime:  defaultTickerTime,
				TickerCount: 0,
				RandomChars: defaultRandomChars,
				Output:      nil,
			}, "hello world", true,
		},
		{
			&FlippyText{
				TickerTime:  time.Second * 0,
				TickerCount: defaultTickerCount,
				RandomChars: defaultRandomChars,
				Output:      invalidWriter{},
			}, "hello world", true,
		},
	}
	for _, c := range cases {
		err := c.ft.Write(c.word)
		if (err != nil) != c.wantErr {
			t.Errorf("%+v.Write(%s) err: %t wanted %t", c.ft, c.word, (err != nil), c.wantErr)
		}
	}
}
