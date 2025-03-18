package main

import (
	"fmt"
	"container/list"	
)
var p int = 0;
// var input int = 123;
var input string = "{{1}, 2, 3}";
var array []int;

var array_string [] string;

func main() {

	// JSON : (number, string, array[JSON])
	//test num checking
	// consume_num()
	// // fmt.Printf("p : %d", p);
	// for i := 0; i < len(array); i++ {
	// 	// fmt.Printf("hello \n");
	// 	fmt.Print(array[i], " \n");
	// }\
	// fmt.Print(is_string());
	// consume_string();
	// fmt.Print(string_length(0), "\n");
	// fmt.Print(num_length(0));
	// for i := 0; i < len(array_string); i++ {
	// 	fmt.Print(string_length(0);, "\n");
	// }
	// fmt.Print(array_string[0], "\n");
	fmt.Print(array_length(1), "\n");
}

// func consume_array() {
// 	if is_array() {
// 		consume_bracket()
// 		//require list of comma seperated tokens ignoring strings for 
// 		for is_json() {
// 			consume_json()
// 		}
// 		consume_bracket()
// 	}
// }


func is_array(start_position int) bool {
	//[(JSON ,)*]
	//starts with a bracket
	if(array[p] != '[') {
		return false;
	}
	//ends with a matching bracket
	stack := new(list.List);
	for i := start_position; i < len(input); i++ {
		if(input[i] == '"') {
			//skip string
			i += string_length(i);
		} else if(input[i] == '{') {
			stack.PushFront('{');
		} else if(input[i] == '}') {
			if(stack.Len() == 0) {
				return false;
				//too many right brackets
			}
			stack.Remove(stack.Front());
		}
		if(stack.Len() == 0) {
			//this is a valid paranthesis
			break;
		}
	}
	if(stack.Len() != 0) {
		//too many left brackets
		return false;
	}
	for is_json(){
		//need a function to skip based on the token amount

		if is_comma() {
			//then keep iterating
			start_position++;
		} else if is_end_bracket(){
			return true;
		} else{
			return false;
		}
	}
	return false;
	//has 0 - infinite valid json in between

}

func array_length (start_position int) int {
	stack := new(list.List);
	temp := start_position;
	for i := start_position; i < len(input); i++ {
		if(input[i] == '"') {
			//skip string
			i += string_length(i);
		} else if(input[i] == '{') {
			stack.PushFront('{');
		} else if(input[i] == '}') {
			if(stack.Len() == 0) {
				return -1;
				//too many right brackets
			}
			stack.Remove(stack.Front());
		}
		if(stack.Len() == 0) {
			//this is a valid paranthesis
			temp = i;
			break;
		}
	}
	return temp - start_position + 1;
}

func num_length(start_position int) int {
	temp := start_position
	if(is_0_9_digit(temp)) {
		temp++;
	}
	for is_1_9_digit(temp){
		temp++;
	}
	return temp - start_position;
}

func string_length(start_position int) int {
	temp := start_position
	if(input[temp] != '"') {
		return -1;
	}
	temp++
	for (!is_out_of_bounds(temp)) {
		// fmt.Printf("%d \n", temp);
		if is_char(temp) {
			temp++;
		} else if(is_escape(temp)) {
			if(is_out_of_bounds(temp + 1)) {
				return -1;
			}
			temp+=2;
		} else{
			break;
		}
	}

	if(input[temp] != '"') {
		return -1;
	}
	return temp - start_position + 1;
}

func is_json() bool {
  var result = is_num(p) || is_string() || is_array(p) 
	return result;
}

func is_comma() bool {
	if(input[p] == ',') {
		return true;
	}
	return false;
}

func is_end_bracket() bool {
	if(input[p] == '}') {
		return true;
	}
	return false;
}


func consume_string() {
	// string -> "(char)*"
	// char -> ([a-z] | [A-Z] | '[' | ']')
	if is_string() {
		consume_quote()
		for is_char(p) {
			consume_char()
		}
		consume_quote()
	}
}

func consume_quote() {
	array_string = append(array_string, "\"");
	p++;
}
func consume_char() {
	//consume 
	//if escape character then consume 2 characters 
	if is_escape(p) {
		array_string = append(array_string, string(input[p]));
		p++;
	}
	array_string = append(array_string, string(input[p]));
	p++;

}

func is_string() bool {
	//has a starting quote
	//has 0-infinite characters in between
	//has an ending quote that isnt before an escape character
	
	//use a temporary variable to avoid changing overall p state for consumption
	var temp int = p;
	if(input[temp] != '"') {
		return false
	}
	temp++
	for (!is_out_of_bounds(temp)) {
		// fmt.Printf("%d \n", temp);
		if is_char(temp) {
			temp++;
		} else if(is_escape(temp)) {
			if(is_out_of_bounds(temp + 1)) {
				return false;
			}
			temp+=2;
		} else{
			break;
		}
	}

	if(input[temp] != '"') {
		return false;
	}
	return true;
}

func is_escape(temp int) bool {
	if(input[temp] == 92) {
		if(is_out_of_bounds(temp)) {

		}
	}
	return false;
}

func is_char(temp int) bool {
	// char -> ([a-z] | [A-Z] | '[' | ']')
	if(input[temp] >= 65 && input[temp] <= 90) {
		return true;
	} else if(input[temp] >= 97 && input[temp] <= 122) {
		return true;
	} else if(input[temp] == 91 || input[temp] == 93) {
		return true;
	} else if(is_0_9_digit(temp)) {
		return true;
	}

	return false;
}



func consume_num() {
	
	if(is_num(p)) {
		// fmt.Printf("HELLO \n")
		//[1-9][0-9]*
		consume_1_9_digit();
		for {
			if !is_0_9_digit(p) {
				break;
			}
			consume_0_9_digit();
		}
	}	
	return;
}

func consume_1_9_digit() {
	array = append(array, int(input[p]) - 48);
	p++;
	return;
}
func consume_0_9_digit() {
	consume_1_9_digit()
}



func is_num(pos int) bool {
	//[1-9][0-9]*
	return is_1_9_digit(p);
}

func is_0_9_digit(pos int) bool {
	//make sure not oob

	if(is_out_of_bounds(pos)) {
		return false;
	}
	
	if(input[pos] <= 57 && input[pos] >= 48) {
		return true;
	}
	return false;
}

func is_1_9_digit(pos int) bool {

	if(is_out_of_bounds(pos)) {
		// fmt.Printf("HELLO out of bounds %d \n", p)
		return false;
	}

	if(input[pos] >= 49 && input[pos] <= 57) {
		return true;
	}
	return false;
}

func is_out_of_bounds(pos int) bool {

	if(pos >= len(input)){
		return true;
	}
	return false;

}

// func deserialize() {

// 	//consume number
// 	if num {
// 		return deserialize_num()

// 	} else if string {
// 		return deserialize_string()

// 	} else if array {
// 		return deserialize_array()
// 	}
// }