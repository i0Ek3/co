package main

import (
	"testing"

	_ "github.com/stretchr/testify/suite"
)

func TestObfuscate(t *testing.T) {

	t.Run("", func(t *testing.T) {
		c := &Confuse{"", true, 3, 8}
		got := c.coalgo3("h")
		want := "-0_9"

		if got != want {
			t.Errorf("got: '%q' want: '%q'", got, want)
		}
	})

	t.Run("", func(t *testing.T) {
		c := &Confuse{"", true, 3, 8}
		got := c.coalgo3("")
		want := ""

		if got != want {
			t.Errorf("got: '%q' want: '%q'", got, want)
		}
	})
	
	t.Run("", func(t *testing.T) {
		c := &Confuse{"", true, 3, 8}
		got := c.dealgo3("-0_9")
		want := "h"

		if got != want {
			t.Errorf("got: '%q' want: '%q'", got, want)
		}
	})
	
	t.Run("", func(t *testing.T) {
		c := &Confuse{"", true, 3, 8}
		got := c.dealgo3("h")
		want := "-0_9"

		if got != want {
			t.Errorf("got: '%q' want: '%q'", got, want)
		}
	})
}
