package flippytext

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	ft := New()
	want := FlippyText{
		TickerTime:  defaultTickerTime,
		TickerCount: defaultTickerCount,
		RandomChars: defaultRandomChars,
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
			}, "hello world", false,
		},
		{
			&FlippyText{
				TickerTime:  defaultTickerTime,
				TickerCount: defaultTickerCount,
				RandomChars: "",
			}, "hello world", true,
		},
		{
			&FlippyText{
				TickerTime:  defaultTickerTime,
				TickerCount: 0,
				RandomChars: defaultRandomChars,
			}, "hello world", false,
		},
	}
	for _, c := range cases {
		err := c.ft.Write(c.word)
		if (err != nil) != c.wantErr {
			t.Errorf("%+v.Write(%s) err: %t wanted %t", c.ft, c.word, (err != nil), c.wantErr)
		}
	}
}
