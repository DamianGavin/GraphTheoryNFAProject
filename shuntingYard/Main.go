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
	for _, r := range infix {
		switch {
		case r == '(':
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '(' {
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		default:
			postfix = append(postfix, r)
		}
	}
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
