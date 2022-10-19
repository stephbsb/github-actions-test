package main

import "testing"

func TestSoma(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Error("Invalid result. Result: %d, Expected: %d", total, 30)
	}
}
