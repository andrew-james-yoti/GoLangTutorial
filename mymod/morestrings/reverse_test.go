package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	reverseStr := ReverseRunes("Hello, Go!")
	if reverseStr == "Hello, Go!" {
		t.Errorf("String not reversed")
	}
}
