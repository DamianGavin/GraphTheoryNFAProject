package main

import (
	"fmt"
)

func intopost(infix string) string {
	specials := map[rune]int{'"': 10, '.': 9, '|': 8}
	var s []rune
	var pofix []rune

	for _, r := range infix {
		switch {
		case r == '(':
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '(' {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		default:
			pofix = append(pofix, r)
		}
	}
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix)
}
func main() {
	//answer should be ab.c*
	fmt.Print("Infix:   ", "a.b.c*")
	fmt.Print("Postfix  ", intopost("a.b.c*"))

	//answer should be abd|.*
	fmt.Print("Infix:   ", "a.(b|d))*")
	fmt.Print("Postfix  ", intopost("(a.(b|d))*"))

	//answer should be abd|.c*.
	fmt.Print("Infix:   ", "a.(b|d).c*")
	fmt.Print("Postfix: ", intopost("a.(b|d).c*"))

	//answer should be abb.+.c.
	fmt.Print("Infix    ", "a.(b.b)+.c")
	fmt.Print("Postfix: ", intopost("a.(b.b)+.c"))

}
