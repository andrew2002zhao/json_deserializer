package main

import (
	"fmt"
	"container/list"	
	"strings"
)
//TODO : change from static variable to local variable for consumption
var p int = 0;
// var input int = 123;
// var input string = "{{1}, 2, 3}";
var array []int;

var array_string [] string;


func main() {

	// value : (number, string, array[value])
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
	fmt.Print(is_array("[[1]]", 0), "\n"); 
}



func is_json (input string, pos int) bool {
	//validate that starting brace has a matching ending brace
	input = eliminate_whitespace(input);
	start_position := pos
	stack := new(list.List);
	for i := start_position; i < len(input); i++ {
		if(input[i] == '"') {
			//skip string
			i += string_length(input, i) - 1;
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

	if(!is_start_json_bracket(input, pos)) {
		return false;
	}
	pos++
	//0 - infinite key value pairs
	for(is_k_v_pair(input, pos)) {
		pos += k_v_pair_length(input, pos);
		if(is_comma(input, pos)) {
			pos++;
		}else {
			break;
		}
	}
	
	if(!is_end_json_bracket(input, pos)){
		return false;
	}
	return true;

}

func is_k_v_pair(input string, pos int) bool {
	if(!is_key(input, pos)){
		return false;
	}
	pos += key_length(input, pos)
	if(!is_colon(input, pos)){
		return false;
	}
	pos++
	if(!is_value(input, pos)) {
		return false;
	}
	return true;
}

func json_length(input string, pos int) int {
	input = eliminate_whitespace(input);
	start_position := pos
	temp := start_position
	stack := new(list.List);
	for i := start_position; i < len(input); i++ {
		if(input[i] == '"') {
			//skip string
			i += string_length(input, i) - 1;
		} else if(input[i] == '{') {
			stack.PushFront('}');
		} else if(input[i] == '}') {
			if(stack.Len() == 0) {
				return -1;
				//too many right brackets
			}
			stack.Remove(stack.Front());
		}
		if(stack.Len() == 0) {
			//this is a valid paranthesis
			temp = i
			break;
		}
	}
	if(stack.Len() != 0) {
		//too many left brackets
		return -1;
	}
	return temp - pos + 1;

}

func is_colon(input string, pos int) bool {
	return input[pos] == ':';
}

func k_v_pair_length(input string, pos int) int {
	key_length := key_length(input, pos);
	value_length := value_length(input, pos + key_length + 1);
	return key_length + value_length + 1
}

func is_key(input string, pos int) bool {
	return is_string(input, pos);
}

func key_length(input string, pos int) int {
	return string_length(input, pos);
}



func is_array(input string, start_position int) bool {
	//[(value ,)*]
	//starts with a bracket
	pos := start_position
	if(input[pos] != '[') {
		return false;
	}
	//ends with a matching bracket
	stack := new(list.List);
	for i := start_position; i < len(input); i++ {
		if(input[i] == '"') {
			//skip string
			i += string_length(input, i) - 1;
		} else if(input[i] == '[') {
			stack.PushFront('[');
		} else if(input[i] == ']') {
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
	pos++;
	for is_value(input, pos){

		//need a function to skip based on the token amount
		length := value_length(input, pos)
		pos += length
		if is_comma(input, pos) {
			//then keep iterating
			// start_position++;
			pos++;
		} else if is_end_array_bracket(input, pos){
			return true;
		} else{
			
			return false;
		}
	}
	
	return true;
	//has 0 - infinite valid value in between

}

func array_length (input string, start_position int) int {
	stack := new(list.List);
	temp := start_position;
	for i := start_position; i < len(input); i++ {
		if(input[i] == '"') {
			//skip string
			i += string_length(input, i) - 1;
		} else if(input[i] == '[') {
			stack.PushFront('{');
		} else if(input[i] == ']') {
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

func num_length(input string, start_position int) int {
	temp := start_position
	if(is_0_9_digit(input, temp)) {
		temp++;
	}
	for is_1_9_digit(input, temp){
		temp++;
	}
	return temp - start_position;
}

func string_length(input string, start_position int) int {
	temp := start_position
	if(input[temp] != '"') {
		return -1;
	}
	temp++
	for (!is_out_of_bounds(input, temp)) {
		// fmt.Printf("%d \n", temp);
		if is_char(input, temp) {
			temp++;
		} else if(is_escape(input, temp)) {
			if(is_out_of_bounds(input, temp + 1)) {
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

func value_length(input string, pos int) int {
	var result int = 0;
	if(is_string(input, pos)){
		result = string_length(input, pos);
	} else if (is_array(input, pos)) {
		result = array_length(input, pos);
	} else if (is_num(input, pos)) {
		result = num_length(input, pos);
	} else if (is_json(input, pos)){
		result = json_length(input, pos);
	}
	return result;
}

func is_value(input string, pos int) bool {
  var result = is_num(input, pos)|| is_string(input, pos) || is_array(input, pos) || is_json(input, pos) 
	return result;
}
func is_comma(input string, pos int) bool {
	return input[pos] == ','
}

func is_start_json_bracket(input string, pos int) bool {
	return input[pos] == '{'
}

func is_end_json_bracket(input string, pos int) bool {
	return input[pos] == '}'
}

func is_start_array_bracket(input string, pos int) bool {
	return input[pos] == '['
}

func is_end_array_bracket(input string, pos int) bool {
	return input[pos] == ']'
}


func consume_string(input string) {
	// string -> "(char)*"
	// char -> ([a-z] | [A-Z] | '[' | ']')
	if is_string(input, p) {
		consume_quote(input)
		for is_char(input, p) {
			consume_char(input)
		}
		consume_quote(input)
	}
}

func consume_quote(input string) {
	array_string = append(array_string, "\"");
	p++;
}
func consume_char(input string) {
	//consume 
	//if escape character then consume 2 characters 
	if is_escape(input, p) {
		array_string = append(array_string, string(input[p]));
		p++;
	}
	array_string = append(array_string, string(input[p]));
	p++;

}

func is_string(input string, pos int) bool {
	//has a starting quote
	//has 0-infinite characters in between
	//has an ending quote that isnt before an escape character
	
	//use a temporary variable to avoid changing overall p state for consumption

	if(input[pos] != '"') {
		return false
	}
	pos++
	for (!is_out_of_bounds(input, pos)) {
		// fmt.Printf("%d \n", temp);
		if is_char(input, pos) {
			pos++;
		} else if(is_escape(input, pos)) {
			if(is_out_of_bounds(input, pos + 1)) {
				return false;
			}
			pos+=2;
		} else{
			break;
		}
	}

	if(input[pos] != '"') {
		return false;
	}
	return true;
}

func is_escape(input string, temp int) bool {
	if(input[temp] == 92) {
		if(is_out_of_bounds(input, temp)) {

		}
	}
	return false;
}

func is_char(input string, temp int) bool {
	// char -> ([a-z] | [A-Z] | '[' | ']')
	if(input[temp] >= 65 && input[temp] <= 90) {
		return true;
	} else if(input[temp] >= 97 && input[temp] <= 122) {
		return true;
	} else if(input[temp] == 91 || input[temp] == 93) {
		return true;
	} else if(is_0_9_digit(input, temp)) {
		return true;
	}

	return false;
}



func consume_num(input string) {
	
	if(is_num(input, p)) {
		// fmt.Printf("HELLO \n")
		//[1-9][0-9]*
		consume_1_9_digit(input);
		for {
			if !is_0_9_digit(input, p) {
				break;
			}
			consume_0_9_digit(input);
		}
	}	
	return;
}

func consume_1_9_digit(input string) {
	array = append(array, int(input[p]) - 48);
	p++;
	return;
}
func consume_0_9_digit(input string) {
	consume_1_9_digit(input)
}



func is_num(input string, pos int) bool {
	//[1-9][0-9]*
	if(!is_1_9_digit(input, pos)) {
		return false;
	}
	pos++;
	for{
		if(is_0_9_digit(input, pos)){
			pos++
		} else if(is_char(input, pos)){
			return false;
		} else {
			break;
		}
	}
	return true;
}

func is_0_9_digit(input string, pos int) bool {
	//make sure not oob

	if(is_out_of_bounds(input, pos)) {
		return false;
	}
	
	if(input[pos] <= 57 && input[pos] >= 48) {
		return true;
	}
	return false;
}

func is_1_9_digit(input string, pos int) bool {

	if(is_out_of_bounds(input, pos)) {
		// fmt.Printf("HELLO out of bounds %d \n", p)
		return false;
	}

	if(input[pos] >= 49 && input[pos] <= 57) {
		return true;
	}
	return false;
}

func is_out_of_bounds(input string, pos int) bool {

	if(pos >= len(input)){
		return true;
	}
	return false;

}

func eliminate_whitespace(input string) string {
	return strings.ReplaceAll(input, " ", "");
}
