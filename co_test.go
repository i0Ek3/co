package main

import (
	"fmt"
	"testing"
)

func TestObfuscate(t *testing.T) {
	c := &Confuse{"", true, 3, 8}
	check := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	}

	t.Run("obfuscate", func(t *testing.T) {
		got := c.coalgo3("h")
		want := "-0_9"
		check(t, got, want)
	})

	t.Run("obfuscate", func(t *testing.T) {
		got := c.coalgo3("")
		want := ""
		check(t, got, want)
	})
}

func BenchmarkObfuscate(b *testing.B) {
	c := &Confuse{"", true, 3, 8}
	for i := 0; i < b.N; i++ {
		c.Obfuscate("hello", c.algoid)
	}
}

func BenchmarkDeobfuscate(b *testing.B) {
	c := &Confuse{"", true, 2, 8}
	for i := 0; i < b.N; i++ {
		c.Deobfuscate("-0_9-0_6", c.algoid)
	}
}

func ExampleObfuscate() {
	c := &Confuse{"", true, 3, 8}
	ret := c.coalgo3("h")
	fmt.Println(ret)
	// Output: -0_9
}
