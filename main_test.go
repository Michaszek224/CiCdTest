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

func TestMul(t *testing.T) {
	got := mul(4, 2)
	want := 8
	if got != want {
		t.Errorf("mul(4, 2) = %d, want %d", got, want)
	}
}

func TestDiv(t *testing.T) {
	got, err := div(9, 3)

	if err != nil {
		t.Errorf("div(9, 3) error = %v", err)
	}

	want := 3
	if got != want {
		t.Errorf("div(9, 3) = %d, want %d", got, want)
	}
}

func TestDivByZero(t *testing.T) {
	_, err := div(9, 0)

	if err == nil {
		t.Fatal("expected error while dividing by zero")
	}

	expected := "division by zero is not allowed"
	if err.Error() != expected {
		t.Errorf("div(9, 0) error = %v, want %v", err, expected)
	}
}
