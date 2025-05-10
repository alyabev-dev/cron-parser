package main

import (
	"reflect"
	"testing"
)

func TestParseWildcard(t *testing.T) {
	parser := &Parser{}
	field := parser.Parse("*", 0, 59, "minute")
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59}

	if !reflect.DeepEqual(field.Values, expected) {
		t.Errorf("expected %v, got %v", expected, field.Values)
	}
}

func TestParseStep(t *testing.T) {
	parser := &Parser{}
	field := parser.Parse("*/15", 0, 59, "minute")
	expected := []int{0, 15, 30, 45}

	if !reflect.DeepEqual(field.Values, expected) {
		t.Errorf("expected %v, got %v", expected, field.Values)
	}
}

func TestParseRange(t *testing.T) {
	parser := &Parser{}
	field := parser.Parse("1-5", 0, 59, "minute")
	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(field.Values, expected) {
		t.Errorf("expected %v, got %v", expected, field.Values)
	}
}

func TestParseList(t *testing.T) {
	parser := &Parser{}
	field := parser.Parse("1,2,3", 0, 59, "minute")
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(field.Values, expected) {
		t.Errorf("expected %v, got %v", expected, field.Values)
	}
}

func TestParseSingleValue(t *testing.T) {
	parser := &Parser{}
	field := parser.Parse("5", 0, 59, "minute")
	expected := []int{5}

	if !reflect.DeepEqual(field.Values, expected) {
		t.Errorf("expected %v, got %v", expected, field.Values)
	}
}

func TestInvalidInput(t *testing.T) {
	parser := &Parser{}

	field := parser.Parse("*/0", 0, 59, "minute")
	if len(field.Values) != 0 {
		t.Errorf("expected empty values for invalid step, got %v", field.Values)
	}

	field = parser.Parse("10-5", 0, 59, "minute")
	if len(field.Values) != 0 {
		t.Errorf("expected empty values for invalid range, got %v", field.Values)
	}

	field = parser.Parse("1,60", 0, 59, "minute")
	expected := []int{1}
	if !reflect.DeepEqual(field.Values, expected) {
		t.Errorf("expected %v, got %v", expected, field.Values)
	}

	field = parser.Parse("100", 0, 59, "minute")
	if len(field.Values) != 0 {
		t.Errorf("expected empty values for out-of-range single value, got %v", field.Values)
	}
}
