//Damian Gavin: Graph theory project 2018

//Non-Finite-Automaton
//Adapted from https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
//https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d

//I have also used https://swtch.com/~rsc/regexp/regexp1.html in my research for this project

package nfa

import "fmt"

//Struct called state with 2 edges.
type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

//This struct will keep track of the initial state and
//the accept state of the fragment of the automaton.
type nfa struct {
	initial *state
	accept  *state
}

func infixToPostfix(infix string) string {
	//A map called specials which will map *, ., | to integer values
	specials := map[rune]int{'*': 10, '+': 9, '?': 8, '.': 7, '|': 6}
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

//poregtonfa is post-fix regular expression to non deterministic
//finite automaton.Must return a pointer to one of the structs.
//An array of pointers to nfa's that are empty."[]*nfa"
func poregtonfa(pofix string) *nfa {
	var nfastack []*nfa
	//This is the algorithm. I loop through the postfix regular expression a character
	//or a rune at a time.
	for _, r := range pofix {
		switch r {
		case '.': //concatenation
			frag2 := nfastack[len(nfastack)-1]    //pop something off nfa stack.
			nfastack = nfastack[:len(nfastack)-1] //get rid of the last thing ":"=up to
			frag1 := nfastack[len(nfastack)-1]    //pop another thing off stack.
			nfastack = nfastack[:len(nfastack)-1] //for concatenation.

			//join them together and push the concatenated fragment back to nfa stack
			//&nfa gives us a pointer to the address.
			frag1.accept.edge1 = frag2.initial
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

		case '|': //or is similar, but I need new accept and initial states.
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			//accept is just an ordinary state
			accept := state{}
			//new initial is a new state where edge1 points at frag1 initial
			//and edge2 points at frag2 initial
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			//push a new thing on the stack that I have created. The new initial state
			//of the fragment I was pushing
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '*': //Kleene star
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			//new initial state with edge1 as initial of fragment popped off
			//and edge2 needs to point at new accept state
			initial := state{edge1: frag.initial, edge2: &accept}
			//join accept state of the fragment.edge1 to the initial state off
			//the fragment just popped off.
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '+':
			//take a single element off the stack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			accept := state{}
			//create an edge pointing back at itself
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = &initial
			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})

		case '?':
			//take a single element off the stack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			// create a new state that points to the existing item and also the accept state
			initial := state{edge1: frag.initial, edge2: frag.accept}
			//push the new nfa onto the stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})

		default: //all I need is to push off the stack, create new accept state.
			accept := state{}
			//new initial state where r points at edge1 state just created.
			initial := state{symbol: r, edge1: &accept}
			//push to nfa stack as a fragment. initial is a pointer to initial
			//state just created and accept points to accept just created.
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	if len(nfastack) != 1 {
		fmt.Println("Uh oh", len(nfastack), nfastack)
	}
	//the goal is only 1 item on nfa stack at the end
	return nfastack[0]
}

//this helper function is recursive.
func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	//any state that has the 0 value of runes it means there are e arrows
	//coming from it.
	if s != a && s.symbol == 0 {
		//if we get in here we're not in the accept state
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
}

//this function takes a postfix reg ex as 1st argument and string s.
func pomatch(po string, s string) bool {
	ismatch := false
	ponfa := poregtonfa(po)

	//I need to keep track of where I am
	var current []*state
	var next []*state

	//traverse the linked list of ponfa by creating a function addState
	//every time I need to add a state this will be called. It will also look at the states
	//available to the current state.
	current = addState(current[:], ponfa.initial, ponfa.accept)

	//loop through s a character at a time. Everytime I read a character I llo
	//through current array "c".
	for _, r := range s {
		for _, c := range current {
			//if c is the same as the symbol I am reading from r
			if c.symbol == r {
				//follow the arrow.
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		//swap current and next, and forget previous next and make it current.
		current, next = next, []*state{}

	}
	//loop through the current array(the state I am in) and if the state I am
	//looping through=accept state of ponfa, match is true.
	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}

func MatchString(infix, testString string) bool {
	postfixStr := infixToPostfix(infix)
	return pomatch(postfixStr, testString)
}
