package main

import (
	"testing"
)


// TODO: debug
// func TestIsvalueFailingMisingBrackets(t *testing.T) {
// 	input:="12, 3"
// 	want:=false
// 	if is_value(input, 0) != want {
// 		t.Errorf("is_value(), want false for input %q", input)
// 	}
// }



func TestIsValueFailingMisingEndBracket(t *testing.T) {
	input:="[12, 3"
	want:=false
	if is_value(input, 0) != want {
		t.Errorf("is_value(), want false for input %q", input)
	}
}

func TestIsValueQuoteBracket(t *testing.T) {
		input:="[12, \"]\", \"1\"]"
		want:=true
		if is_value(input, 0) != want {
			t.Errorf("is_value(), want true for input %q", input)
		}
}


func TestIsValueSimple(t *testing.T) {
	input:="[12, \"1\"]"
	want:=true
	if is_value(input, 0) != want {
		t.Errorf("is_value(), want true for input %q", input)
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