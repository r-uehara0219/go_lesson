package main

import (
	"strconv"
	"testing"
)

func TestIsEven(t *testing.T) {
	cases := []struct {
		input    int
		expected bool
	}{
		{input: 1, expected: false},
		{input: 2, expected: true},
	}
	for _, c := range cases {
		//shadowing variables so that jobs running in parallel are not affected
		c := c
		t.Run(strconv.Itoa(c.input), func(t *testing.T) {
			if actual := IsEven(c.input); c.expected != actual {
				t.Errorf("want IsEven(%d) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}
}
