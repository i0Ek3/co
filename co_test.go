package main

import (
	"fmt"
	"testing"
)

func TestObfuscate(t *testing.T) {
	c := &Confuse{cobit: 8}
	check := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	}

	t.Run("obfuscate", func(t *testing.T) {
		got := c.coalgo4String("h")
		want := "-1_0"
		check(t, got, want)
	})

	t.Run("obfuscate", func(t *testing.T) {
		got := c.coalgo4String("")
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
		c.Deobfuscate("-1_0-1_0-1_0", c.algoid)
	}
}

func ExampleObfuscate() {
	c := &Confuse{"", true, 3, 8}
	ret := c.coalgo4String("h")
	fmt.Println(ret)
	// Output: -1_0
}
