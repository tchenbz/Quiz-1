package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	got := area()
	expected := 25

	if got != expected {
		t.Error("got %q expected %q", got, expected)
	}
}

func TestPerimeter(t *testing.T) {
	got := perimeter()
	expected := 30

	if got != expected {
		t.Error("got %q expected %q", got, expected)
	}
}
