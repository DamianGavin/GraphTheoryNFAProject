//Damian Gavin: Graph theory project 2018
//Non-Finite-Automaton
//Adapted from https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
//https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d

//I have also used https://swtch.com/~rsc/regexp/regexp1.html in my research for this project

package main

import (
	"fmt"
)

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

//poregtonfa is post-fix regular expression to non deterministic
//finite automaton.Must return a pointer to one of the structs.
//An array of pointers to nfa's that are empty."[]*nfa"
func poregtonfa(pofix string) *nfa {
	var nfastack []*nfa
	//This is the algorithm. I loop through the postfix regular expression a character
	//or a rune at a time.
	for _, r := range pofix {
		switch r {
		/*case '.':
			e2 = pop();
			e1 = pop();
			patch(e1.out, e2.start);
			push(frag(e1.start, e2.out));
			break;*/
		case '.':                                 //concatenation
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

		case '*': //Clainey star
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
		default: //all I need is to push off the stack, create new accept state.
			accept := state{}
			//new initial state where r points at edge1 state just created.
			initial := state{symbol: r, edge1: &accept}
			//push to nfa stack as a fragment. initial is a pointer to initial
			//state just created and accept points to accept just created.
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	if len(nfastack) != 1{
		fmt.Println("Uh oh", len(nfastack), nfastack)
	}
	//the goal is only 1 item on nfa stack at the end
	return nfastack[0]
}
func addstate(l []*state, s *state, a *state) []*state{
	l = append(l, s)

	if s != a && s.symbol == 0{
		l = addstate(l, s.edge1, a)
		if s.edge2 != nil{
			l = addstate (l, s.edge2, a)
		}
	}
	return l
}

func pomatch(po string, s string) bool{
	ismatch := false
	ponfa := poregtonfa(po)

	var current []*state
	var next []*state

	current = addstate(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {
		for _, c := range current{
			if c.symbol == r{
				next = addstate(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*state{}

	}
	for _, c := range current{
		if c == ponfa.accept{
			ismatch = true
			break
		}
	}

	return ismatch
}


func main() {
	fmt.Println(pomatch("ab.c*|", "cccc"))
}