package main

import "testing"

func TestHello2(t *testing.T) {
	expected := "hello world!"

	actual := Hello()

	if expected != actual {
		t.Errorf("got %q want %q", actual,expected)
	}
}