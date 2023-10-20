package main

import "testing"

func TestCountWordsLongString(t *testing.T) {
	input := "This is a test"
	expected := 4

	actual := countWordsInString(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestCountWordsEmptyString(t *testing.T) {
	input := ""
	expected := 0

	actual := countWordsInString(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
