package main

import (
	"testing"
)

// TODO: debug
// currently this is a feature since the way my code works is that it tests valid input from start character to whenever it exits

// this fails because it assumes there are no extra characters after
// even if we have matching left bracket sets
// that doesnt gurantee we have matching right bracket sets

// what my mind is thinking rn is that after check for validity once, we dont do a rigorous check after
// like if we have an array, there should be nothing after unless its nested arrays or commas
// the problem is with comma and ending brace consumption
// i think solution should be to error on double commas
// so when we value check, commas, braces, brackets and commas are only valid within their context
// one solution is stateful
// if we have a boolean to indicate that we are currently within a brace then any comma is valid,
// but wait wouldnt

// func TestIsValueFailingInvalidInputImproperCommaBehind(t *testing.T) {
// 	input:="1,"
// 	want:=false
// 	if is_value(input, 0) != want {
// 		t.Errorf("is_value(), want false for input %q", input)
// 	}
// }

// func TestJsonFailingExtraRightBracket(t *testing.T) {
// 	input := "{\"hello\":{\"hello2\":123}}}"
// 	want := false
// 	if is_json(input, 0) != want {
// 		t.Errorf("is_json(), wanted false for input %q", input)
// 	}
// }

// func TestIsArrayFailingExtraRightBracket(t *testing.T) {
// 	input := "[12, 1]]"
// 	want:=false
// 	if is_array(input, 0) != want {
// 		t.Errorf("is_array(), want false for input %q", input)
// 	}
// }

func TestIsJsonMixedNestedArraysAndMultipleJSON(t *testing.T) {
	input := " {\"hello\":123,\"hello2\":[1,2,[\"a\"]]}"
	want := true
	if is_json(input, 0) != want {
		t.Errorf("is_json(), want true for input %q", input)
	}
}

func TestIsJsonMultipleJSON(t *testing.T) {
	input := " {\"hello\":123,\"hello2\":21}"
	want := true
	if is_json(input, 0) != want {
		t.Errorf("is_json(), want true for input %q", input)
	}
}

func TestIsJsonLeadingWhiteSpace(t *testing.T) {
	input := "         {\"hello\":{\"hello2\":123}}"
	want := true
	if is_json(input, 0) != want {
		t.Errorf("is_json(), want true for input %q", input)
	}
}

func TestIsJsonTrailingWhiteSpace(t *testing.T) {
	input := "{\"hello\":{\"hello2\":123}}      "
	want := true
	if is_json(input, 0) != want {
		t.Errorf("is_json(), want true for input %q", input)
	}
}

func TestIsJsonRandomWhiteSpace(t *testing.T) {
	input := "{  \"hello\":{  \"hello2\" :123}  }"
	want := true
	if is_json(input, 0) != want {
		t.Errorf("is_json(), want true for input %q", input)
	}
}

func TestIsArrayFailingExtraLeftBracket(t *testing.T) {
	input := "[[12, 1]"
	want := false
	if is_array(input, 0) != want {
		t.Errorf("is_array(), want false for input %q", input)
	}
}

func TestIsValueFailingInvalidInputMixedValues(t *testing.T) {
	input := "1a2"
	want := false
	if is_value(input, 0) != want {
		t.Errorf("is_value(), want false for input %q", input)
	}
}

func TestIsValueFailingInvalidInputImproperCommaFront(t *testing.T) {
	input := ",1"
	want := false
	if is_value(input, 0) != want {
		t.Errorf("is_value(), want false for input %q", input)
	}
}

func TestJsonFailingExtraLeftBracket(t *testing.T) {
	input := "{{\"hello\":{\"hello2\":123}}"
	want := false
	if is_json(input, 0) != want {
		t.Errorf("is_json(), wanted false for input %q", input)
	}
}

func TestJsonLengthNested(t *testing.T) {
	input := "{\"hello\":{\"hello2\":123}}"
	want := 24
	if val := json_length(input, 0); val != want {
		t.Errorf("json_length(), wanted %d instead got %d for input %q", want, val, input)
	}
}

func TestJsonLengthSimple(t *testing.T) {
	input := "{\"hello\":[12,\"123\"]}"
	want := 20
	if val := json_length(input, 0); val != want {
		t.Errorf("json_length(), wanted %d instead got %d for input %q", want, val, input)
	}
}

func TestIsJsonNested(t *testing.T) {
	input := "{\"hello\":{\"hello2\":123}}"
	want := true
	if is_json(input, 0) != want {
		t.Errorf("is_json(), want true for input %q", input)
	}
}

func TestIsJsonSimple(t *testing.T) {
	input := "{\"hello\":[12,\"123\"]}"
	want := true
	if is_json(input, 0) != want {
		t.Errorf("is_json(), wanted true for input %q", input)
	}
}

func TestKeyValueLength(t *testing.T) {
	input := "\"hello\":[12, \"123\"]"
	want := 18
	if val := k_v_pair_length(input, 0); val != want {
		t.Errorf("k_v_pair_length(), wanted 18 for input %q, got %d instead", input, val)
	}
}

func TestIsKeyValuePairFailingMissingColon(t *testing.T) {
	input := "\"hello\"[12, \"123\"]"
	want := false
	if is_k_v_pair(input, 0) != want {
		t.Errorf("is_k_v_pair(), want false for input %q", input)
	}
}

func TestIsKeyValuePairArrayValue(t *testing.T) {
	input := "\"hello\":[12, \"123\"]"
	want := true
	if is_k_v_pair(input, 0) != want {
		t.Errorf("is_k_v_pair(), want true for input %q", input)
	}
}

func TestIsKeyValuePairSimple(t *testing.T) {
	input := "\"hello\":\"12\""
	want := true
	if is_k_v_pair(input, 0) != want {
		t.Errorf("is_k_v_pair(), want true for input %q", input)
	}
}

func TestValueLength(t *testing.T) {
	input := "[12, \"3\"]"
	want := 9
	if val := value_length(input, 0); val != want {
		t.Errorf("is_k_v_pair(), wanted %d for input %q, got %d instead", want, input, val)

	}
}

func TestIsValueFailingMisingEndBracket(t *testing.T) {
	input := "[12, 3"
	want := false
	if is_value(input, 0) != want {
		t.Errorf("is_value(), want false for input %q", input)
	}
}

func TestIsValueQuoteBracket(t *testing.T) {
	input := "[12, \"]\", \"1\"]"
	want := true
	if is_value(input, 0) != want {
		t.Errorf("is_value(), want true for input %q", input)
	}
}

func TestIsValueSimple(t *testing.T) {
	input := "[12, \"1\"]"
	want := true
	if is_value(input, 0) != want {
		t.Errorf("is_value(), want true for input %q", input)
	}
}

func TestIsArrayNestedArrays2(t *testing.T) {
	array := "[[123]]"
	want := true
	if is_array(array, 0) != want {
		t.Errorf("is_array(), want array for input %q", array)
	}
}

func TestIsArrayNestedArrays1(t *testing.T) {
	array := "[123,[12]]"
	want := true
	if is_array(array, 0) != want {
		t.Errorf("is_array(), want array for input %q", array)
	}
}

func TestIsArraySimpleWorking(t *testing.T) {
	array := "[123,12]"
	want := true
	if is_array(array, 0) != want {
		t.Errorf("is_array(), want array for input %q", array)
	}
}
func TestIsStringSimple(t *testing.T) {
	input := "\"abc\""
	want := true
	if is_string(input, 0) != want {
		t.Errorf("is_string(), want true for %q", input)
	}
}

func TestIsStringEscapeCharacter(t *testing.T) {
	input := "\"\"abc\""
	want := true
	if is_string(input, 0) != want {
		t.Errorf("is_string(), want true for %q", input)
	}
}

func TestArrayLengthNestedArray(t *testing.T) {
	array := "[[123],12]"
	want := 10
	if array_length(array, 0) != want {
		t.Errorf("array_length(0), want array for input %q", array)
	}
}

func TestArrayLengthTwoElements(t *testing.T) {
	array := "[123,12]"
	want := 8
	if array_length(array, 0) != want {
		t.Errorf("array_length(0), want array for input %q", array)
	}
}

func TestArrayLengthOneElement(t *testing.T) {
	array := "[123]"
	want := 5
	if array_length(array, 0) != want {
		t.Errorf("array_length(0), want array for input %q", array)
	}
}
