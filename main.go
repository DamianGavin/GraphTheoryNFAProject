//Damian Gavin: Graph theory project 2018
//Non-Finite-Automaton
//Adapted from https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
//https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d

//I have also used https://swtch.com/~rsc/regexp/regexp1.html in my research for this project

package main

import (
	"./nfa"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter regular expression: ")
	scanner.Scan()
	regex := scanner.Text()
	fmt.Print("Enter test string: ")
	scanner.Scan()
	testString := scanner.Text()
	//fmt.Println(nfa.MatchString("a.b|c*", "ccccc"))
	if nfa.MatchString(regex, testString) {
		fmt.Println("The regular expression", regex, "matched the string", testString)
	} else {
		fmt.Println("The regular expression", regex, "did not match the string", testString)
	}
}
