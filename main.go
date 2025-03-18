package main

import (
	"fmt"
)
var p int = 0;
var input string = "\"321\"";
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
	// fmt.Print(check_if_string());
	consume_string();
	for i := 0; i < len(array_string); i++ {
		fmt.Print(array_string[i], "\n");
	}
	// fmt.Print(array_string[0], "\n");
}



func consume_string() {
	// string -> "(char)*"
	// char -> ([a-z] | [A-Z] | '[' | ']')
	if check_if_string() {

		consume_quote()
		for check_if_char(p) {
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
	if check_if_escape(p) {
		array_string = append(array_string, string(input[p]));
		p++;
	}
	array_string = append(array_string, string(input[p]));
	p++;

}

func check_if_string() bool {
	//has a starting quote
	//has 0-infinite characters in between
	//has an ending quote that isnt before an escape character
	
	//use a temporary variable to avoid changing overall p state for consumption
	var temp int = p;
	if(input[temp] != '"') {
		return false
	}
	temp++
	for (!check_if_out_of_bounds(temp)) {
		// fmt.Printf("%d \n", temp);
		if check_if_char(temp) {
			temp++;
		} else if(check_if_escape(temp)) {
			if(check_if_out_of_bounds(temp + 1)) {
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

func check_if_escape(temp int) bool {
	if(input[temp] == 92) {
		if(check_if_out_of_bounds(temp)) {

		}
	}
	return false;
}

func check_if_char(temp int) bool {
	// char -> ([a-z] | [A-Z] | '[' | ']')
	if(input[temp] >= 65 && input[temp] <= 90) {
		return true;
	} else if(input[temp] >= 97 && input[temp] <= 122) {
		return true;
	} else if(input[temp] == 91 || input[temp] == 93) {
		return true;
	} else if(check_if_0_9_digit(temp)) {
		return true;
	}

	return false;
}



func consume_num() {
	
	if(check_if_num()) {
		// fmt.Printf("HELLO \n")
		//[1-9][0-9]*
		consume_1_9_digit();
		for {
			if !check_if_0_9_digit(p) {
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



func check_if_num() bool {
	//[1-9][0-9]*
	return check_if_1_9_digit(p);
}

func check_if_0_9_digit(temp int) bool {
	//make sure not oob

	if(check_if_out_of_bounds(temp)) {
		return false;
	}
	
	if(input[temp] <= 57 && input[temp] >= 48) {
		return true;
	}
	return false;
}

func check_if_1_9_digit(temp int) bool {

	if(check_if_out_of_bounds(temp)) {
		// fmt.Printf("HELLO out of bounds %d \n", p)
		return false;
	}

	if(input[temp] >= 49 && input[temp] <= 57) {
		return true;
	}
	return false;
}

func check_if_out_of_bounds(temp int) bool {

	if(temp >= len(input)){
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