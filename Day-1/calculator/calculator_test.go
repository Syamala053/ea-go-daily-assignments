package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(6, 2)
	expected := 8

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := sub(6, 2)
	expected := 4

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestMul(t *testing.T) {
	result := mul(6, 2)
	expected := 12

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestDiv(t *testing.T) {
	result := div(6, 2)
	expected := 3

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}
