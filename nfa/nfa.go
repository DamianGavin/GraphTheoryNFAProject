//Damian Gavin: Graph theory project 2018

//Non-Finite-Automaton
//Adapted from https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
//https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d

//I have also used https://swtch.com/~rsc/regexp/regexp1.html in my research for this project

package nfa

import "fmt"

//Struct called state with 2 edges.
type state struct {
	symbol rune //default is 0. Represents a state that has 2 arrows coming from it
	edge1  *state
	edge2  *state
}

//This struct will keep track of the initial state and
//the accept state of the fragment of the automaton.
type nfa struct {
	initial *state
	accept  *state
}

//infixToPostfix converts infix regex to postfix, shunting yard.
func infixToPostfix(infix string) string {
	//A map called specials which will map *, ., | to integer values
	specials := map[rune]int{'*': 10, '+': 9, '?': 8, '.': 7, '|': 6}
	//stack used to temporarily store operators as read from infix string
	var stack []rune
	//Built in rune datatype. A rune is a character as displayed on-screen(UTF-8)
	var postfix []rune

	//loops through infix and returns index of character.
	//converts strings to an array of runes.
	for _, r := range infix {
		switch {
		//1st case is open bracket "("
		//append is a built in function that just adds to the end.
		case r == '(':
			stack = append(stack, r)
		case r == ')':
			//closing bracket will pop off the stack until we find the open bracket.
			//ie. for the last element != "("
			for stack[len(stack)-1] != '(' {
				postfix, stack = append(postfix, stack[len(stack)-1]), stack[:len(stack)-1]
			}
			//if the character is a closing bracket keep popping off the stack.
			stack = stack[:len(stack)-1]

			//this will be true if r is in the special characters array
		case specials[r] > 0:
			//while stack still has elements and the precedence of current element
			//is <= the precedence of the top element of the stack. Take the element off the
			//stack and put it into postfix.
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
				postfix, stack = append(postfix, stack[len(stack)-1]), stack[:len(stack)-1]
			}
			stack = append(stack, r)

			//default is when r is neither a bracket nor a special character.
		default:
			postfix = append(postfix, r)
		}
	}
	//at the end of the process if there is anything left on the stack just
	//append it onto the output.
	for len(stack) > 0 {
		postfix, stack = append(postfix, stack[len(stack)-1]), stack[:len(stack)-1]
	}

	//cast postfix to a string
	return string(postfix)
}

//postRegExToNfa is post-fix regular expression to non deterministic
//finite automaton.Must return a pointer to one of the structs.
//An array of pointers to nfa's that are empty."[]*nfa"
func postRegExToNfa(postFix string) *nfa {
	var nfaStack []*nfa
	//This is the algorithm. I loop through the postfix regular expression a character
	//or a rune at a time.
	for _, r := range postFix {
		switch r {
		case '.': //concatenation
			frag2 := nfaStack[len(nfaStack)-1]    //pop something off nfa stack.
			nfaStack = nfaStack[:len(nfaStack)-1] //get rid of the last thing ":"=up to
			frag1 := nfaStack[len(nfaStack)-1]    //pop another thing off stack.
			nfaStack = nfaStack[:len(nfaStack)-1] //for concatenation.

			//join them together and push the concatenated fragment back to nfa stack
			//&nfa gives us a pointer to the address.
			frag1.accept.edge1 = frag2.initial
			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})

		case '|': //or is similar, but I need new accept and initial states.
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			//accept is just an ordinary state
			accept := state{}
			//new initial is a new state where edge1 points at frag1 initial
			//and edge2 points at frag2 initial
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			//push a new thing on the stack that I have created. The new initial state
			//of the fragment I was pushing
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		case '*': //Kleene star
			//pop off stack
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			//new fragment is the old fragment with 2 extra states
			accept := state{}
			//new initial state with edge1 as initial of fragment popped off
			//and edge2 needs to point at new accept state
			initial := state{edge1: frag.initial, edge2: &accept}
			//join accept state of the fragment.edge1 to the initial state off
			//the fragment just popped off.
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			//push new fragment to the stack
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		case '+':
			//take a single element off the stack
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			accept := state{}
			//create an edge pointing back at itself
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = &initial
			nfaStack = append(nfaStack, &nfa{initial: frag.initial, accept: &accept})

		case '?':
			//take a single element off the stack
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			// create a new state that points to the existing item and also the accept state
			initial := state{edge1: frag.initial, edge2: frag.accept}
			//push the new nfa onto the stack
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: frag.accept})

		default: //all I need is to push off the stack, create new accept state.
			accept := state{}
			//new initial state where r points at edge1 state just created.
			initial := state{symbol: r, edge1: &accept}
			//push to nfa stack as a fragment. initial is a pointer to initial
			//state just created and accept points to accept just created.
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}
	}

	//print warning if more than 1 item on stack at the end.
	if len(nfaStack) != 1 {
		fmt.Println("Uh oh", len(nfaStack), nfaStack)
	}
	//the goal is only 1 item on nfa stack at the end
	return nfaStack[0]
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
func poMatch(post string, test string) bool {
	isMatch := false
	poNfa := postRegExToNfa(post)

	//I need to keep track of where I am
	var current []*state
	var next []*state

	//traverse the linked list of poNfa by creating a function addState
	//every time I need to add a state this will be called. It will also look at the states
	//available to the current state.
	current = addState(current[:], poNfa.initial, poNfa.accept)

	//loop through test a character at a time. Everytime I read a character I llo
	//through current array "c".
	for _, r := range test {
		for _, c := range current {
			//if c is the same as the symbol I am reading from r
			if c.symbol == r {
				//follow the arrow.
				next = addState(next[:], c.edge1, poNfa.accept)
			}
		}
		//swap current and next, and forget previous next and make it current.
		current, next = next, []*state{}

	}
	//loop through the current array(the state I am in) and if the state I am
	//looping through=accept state of poNfa, match is true.
	for _, c := range current {
		if c == poNfa.accept {
			isMatch = true
			break
		}
	}

	return isMatch
}

func MatchString(infix, testString string) bool {
	postfixStr := infixToPostfix(infix)
	return poMatch(postfixStr, testString)
}
