package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(1, 2)
	want := 3
	if got != want {
		t.Errorf("add(1, 2) = %d, want %d", got, want)
	}
}

func TestSub(t *testing.T) {
	got := sub(4, 3)
	want := 1
	if got != want {
		t.Errorf("sub(4, 3) = %d, want %d", got, want)
	}
}

func TestDiv(t *testing.T) {
	got := div(9, 3)
	want := 3
	if got != want {
		t.Errorf("div(9, 3) = %d, want %d", got, want)
	}
}

func AnotehrDivTest(t *testing.T) {
	got := div(9, 4)
	want := 2
	if got != want {
		t.Errorf("div(9, 0) = %d, want %d", got, want)
	}
}
func TestMul(t *testing.T) {
	got := mul(4, 2)
	want := 8
	if got != want {
		t.Errorf("mul(4, 2) = %d, want %d", got, want)
	}
}
