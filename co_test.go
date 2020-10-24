package main

import (
	"fmt"
	"testing"
)

func TestObfuscate(t *testing.T) {
	check := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got: '%q' want: '%q'", got, want)
		}
	}

	t.Run("obfuscate", func(t *testing.T) {
		c := &Confuse{"", true, 3, 8}
		got := c.coalgo3("h")
		want := "-0_9"
		check(t, got, want)
	})

	t.Run("obfuscate", func(t *testing.T) {
		c := &Confuse{"", true, 3, 8}
		got := c.coalgo3("")
		want := ""
		check(t, got, want)
	})
}

func ExampleObfuscate() {
	c := &Confuse{"", true, 3, 8}
	ret := c.coalgo3("h")
	fmt.Println(ret)
	// Output: -0_9
}
