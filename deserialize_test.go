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

func TestIsJsonSimple(t *testing.T) {
	input := "{\"hello\":[12, \"123\"]}";
	want := true
	if is_json(input, 0) != want {
		t.Errorf("is_json(), wanted true for input %q", input)
	}
}

func TestKeyValueLength(t *testing.T) {
	input := "\"hello\":[12, \"123\"]";
	want := 19
	if val := k_v_pair_length(input, 0); val != want {
		t.Errorf("k_v_pair_length(), wanted 18 for input %q, got %d instead", input, val)
	}
}


func TestIsKeyValuePairFailingMissingColon(t *testing.T) {
	input := "\"hello\"[12, \"123\"]";
	want := false
	if is_k_v_pair(input, 0) != want {
		t.Errorf("is_k_v_pair(), want false for input %q", input)
	}
}

func TestIsKeyValuePairArrayValue(t *testing.T) {
	input := "\"hello\":[12, \"123\"]";
	want := true
	if is_k_v_pair(input, 0) != want {
		t.Errorf("is_k_v_pair(), want true for input %q", input)
	}
}


func TestIsKeyValuePairSimple(t *testing.T) {
	input := "\"hello\":\"12\"";
	want := true
	if is_k_v_pair(input, 0) != want {
		t.Errorf("is_k_v_pair(), want true for input %q", input)
	}
}

func TestValueLength(t *testing.T) {
	input:="[12, \"3\"]"
	want:= 9
	if val := value_length(input, 0); val != want {
		t.Errorf("is_k_v_pair(), wanted %d for input %q, got %d instead", want, input, val)

	}
}



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