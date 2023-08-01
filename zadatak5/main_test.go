package main

import (
	"reflect"
	"testing"
)

func TestExtractNumbers(t *testing.T) {
	input := "[ 3, 1, 2, 3, 4, 2, 5 ]"
	expected := []int{3,1,2,3,4,2,5}

	result := extractNumbers(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Greška u TestExtractNumbers. Očekivano: %v, Dobijeno: %v", expected, result)
	}
}

func TestDeduplicate(t *testing.T) {
	input := []int{3, 1, 2, 3, 4, 2, 5}
	expected := []int{3,1,2,4,5}

	result := deduplicate(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Greška u TestDeduplicate. Očekivano: %v, Dobijeno: %v", expected, result)
	}
}
