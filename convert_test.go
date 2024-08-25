package main

import (
	"testing"
)

func TestSimpleConvert(t *testing.T) {
	word := "David"
	want := "Dav"
	abbr := SimpleConvert(word)

	if want != abbr {
		t.Fatalf(`SimpleConvert("%s") = %s, want %s`, word, abbr, want)
	}
}

func TestFirstEndConvert(t *testing.T) {
	word := "David"
	want := "Dd"
	abbr := FirstEndConvert(word)

	if want != abbr {
		t.Fatalf(`FirstEndConvert("%s") = %s, want %s`, word, abbr, want)
	}
}
