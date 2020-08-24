package main

import (
	"errors"
	"fmt"
	"container/list"
)


func isBalance(value string) (bool, error) {

	var stack list.List

	for _, element := range value {

		switch element {
			case '(', '{', '[': 
				stack.PushBack(element)	
				continue;
			case ')': 
				if stack.Len() != 0 || stack.Back().Value == '(' { 
					stack.Remove(stack.Back())
					continue;
				}
			case '}':
				
				if stack.Len() != 0 || stack.Back().Value== '{' { 
					stack.Remove(stack.Back())
					continue;
				}
			case ']':
				if stack.Len() != 0 || stack.Back().Value== '[' { 
					stack.Remove(stack.Back())
					continue;
				}

		default:
			return false, fmt.Errorf("Invalid Character '%c'", element)
		}		
	}

	if stack.Len() != 0 {
		return false, errors.New("Is not balance")
	}

	return true, nil
}

func main() {

	fmt.Println(isBalance("{|}"))
}