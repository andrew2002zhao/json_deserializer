package main

import (
	"testing"
)

func TestIsJsonFailingMisingBrackets(t *testing.T) {
	input:="12, 3"
	want:=false
	if is_json(input, 0) != want {
		t.Errorf("is_json(), want false for input %q", input)
	}
}



func TestIsJsonFailingMisingEndBracket(t *testing.T) {
	input:="[12, 3"
	want:=false
	if is_json(input, 0) != want {
		t.Errorf("is_json(), want false for input %q", input)
	}
}

func TestIsJsonQuoteBracket(t *testing.T) {
		input:="[12, \"]\", \"1\"]"
		want:=true
		if is_json(input, 0) != want {
			t.Errorf("is_json(), want true for input %q", input)
		}
}


func TestIsJsonSimple(t *testing.T) {
	input:="[12, \"1\"]"
	want:=true
	if is_json(input, 0) != want {
		t.Errorf("is_json(), want true for input %q", input)
	}
}
 

func TestIsArrayNestedArrays2(t *testing.T) {
	array:="[[123]]"
	want:=true
	if is_array(array, 0) != want {
		t.Errorf("is_array(), want array for input %q", array)
	}
}

func TestIsArrayNestedArrays1(t *testing.T) {
	array:="[123,[12]]"
	want:=true
	if is_array(array, 0) != want {
		t.Errorf("is_array(), want array for input %q", array)
	}
}


func TestIsArraySimpleWorking(t *testing.T) {
	array:="[123,12]"
	want:=true
	if is_array(array, 0) != want {
		t.Errorf("is_array(), want array for input %q", array)
	}
}
func TestIsStringSimple(t *testing.T) {
	input:="\"abc\""
	want:=true
	if is_string(input, 0) != want {
		t.Errorf("is_string(), want true for %q", input)
	}
}

func TestIsStringEscapeCharacter(t *testing.T) {
	input := "\"\"abc\""
	want:=true
	if is_string(input, 0) != want {
		t.Errorf("is_string(), want true for %q", input)
	}
}

func TestArrayLengthNestedArray(t *testing.T) {
	array:="[[123],12]"
	want:=10
	if array_length(array, 0) != want {
		t.Errorf("array_length(0), want array for input %q", array)
	}
}

func TestArrayLengthTwoElements(t *testing.T) {
	array:="[123,12]"
	want:=8
	if array_length(array, 0) != want {
		t.Errorf("array_length(0), want array for input %q", array)
	}
}

func TestArrayLengthOneElement(t *testing.T) {
	array:="[123]"
	want:=5
	if array_length(array, 0) != want {
		t.Errorf("array_length(0), want array for input %q", array)
	}
}