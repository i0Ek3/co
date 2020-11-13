package main

import (
	"fmt"
	"testing"
)

func TestObfuscate(t *testing.T) {
	c := &Confuse{cobit: 8, debug: "false"}
	check := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	}

	t.Run("obfuscate", func(t *testing.T) {
		got := c.coalgo4("hello", c.debug)
		want := "-0_7-0_4-1_3-1_3-1_6"
		check(t, got, want)
	})

	t.Run("obfuscate", func(t *testing.T) {
		got := c.coalgo4("", c.debug)
		want := ""
		check(t, got, want)
	})
}

func BenchmarkObfuscate(b *testing.B) {
	c := &Confuse{"", true, 4, 8, "false"}
	for i := 0; i < b.N; i++ {
		c.Obfuscate("hello", c.debug, c.algoid)
	}
}

func TestDebfuscate(t *testing.T) {
	c := &Confuse{cobit: 8, debug: "false"}
	check := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	}

	t.Run("deobfuscate", func(t *testing.T) {
		got := c.dealgo3("-0_7-0_4-1_3-1_3-1_6", c.debug)
		want := "hello"
		check(t, got, want)
	})

	t.Run("deobfuscate", func(t *testing.T) {
		got := c.dealgo3("", c.debug)
		want := ""
		check(t, got, want)
	})
}

func BenchmarkDeobfuscate(b *testing.B) {
	c := &Confuse{"", true, 3, 8, "false"}
	for i := 0; i < b.N; i++ {
		c.Deobfuscate("-1_0-1_0-1_0", c.debug, c.algoid)
	}
}

func ExampleObfuscate() {
	c := &Confuse{"", true, 3, 8, "false"}
	ret := c.coalgo4("h", c.debug)
	fmt.Println(ret)
	// Output: -0_7
}
