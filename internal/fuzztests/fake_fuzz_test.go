package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzFails(f *testing.F) {
	f.Add(0, 0)
	f.Add(2, 4)

	f.Fuzz(func(t *testing.T, a, b int) {
		expect := a + b
		actual := myAdd(a, b)

		require.Equal(t, expect, actual, "Expect %d + %d = %d", a, b, expect)
	})
}

func myAdd(a, b int) int {
	// Force myAdd to return the wrong sum if a is an odd number.
	if a%2 != 0 {
		return a + b + 1
	}
	return a + b
}
