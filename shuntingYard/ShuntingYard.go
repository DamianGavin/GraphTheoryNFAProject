//Damian Gavin. Graph theory project 2018
//shunting yard algorithm executable.

package main

import (
	"fmt"
)

//function to convert infix reg expressions to post-fix.
func infixToPostfix(infix string) string {
	//A map called specials which will map *, ., | to integer values
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	//stack s
	var s []rune
	//Built in rune datatype. A rune is a character as displayed on-screen(UTF-8)
	var postfix []rune

	//this is the algorithm.
	//loops through infix and returns index of character.
	//converts strings to an array of runes.
	for _, r := range infix {
		switch {
		//1st case is open bracket "("
		//append is a built in function that just adds to the end.
		case r == '(':
			s = append(s, r)
		case r == ')':
			//closing bracket will pop off the stack until we find the open bracket.
			//ie. for the last element != "("
			for s[len(s)-1] != '(' {
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			//if the character is a closing bracket keep popping off the stack.
			s = s[:len(s)-1]

			//this will be true if r is in the special characters array
		case specials[r] > 0:
			//while stack still has elements and the precedence of current element
			//is <= the precedence of the top element of the stack. Take the element off the
			//stack and put it into postfix.
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)

			//default is when r is neither a bracket nor a special character.
		default:
			postfix = append(postfix, r)
		}
	}
	//at the end of the process if there is anything left on the stack just
	//append it onto the output.
	for len(s) > 0 {
		postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
	}
	//cast postfix to a string
	return string(postfix)
}
func main() {
	//answer should be ab.c*
	fmt.Print("Infix:   ", "a.b.c*")
	fmt.Print("Postfix  ", infixToPostfix("a.b.c*"))

	//answer should be abd|.*
	fmt.Print("Infix:   ", "a.(b|d))*")
	fmt.Print("Postfix  ", infixToPostfix("(a.(b|d))*"))

	//answer should be abd|.c*.
	fmt.Print("Infix:   ", "a.(b|d).c*")
	fmt.Print("Postfix: ", infixToPostfix("a.(b|d).c*"))

	//answer should be abb.+.c.
	fmt.Print("Infix    ", "a.(b.b)+.c")
	fmt.Print("Postfix: ", infixToPostfix("a.(b.b)+.c"))

}
