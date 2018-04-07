## GraphTheoryNFAProject

## Introduction

My name is Damian Gavin

This repository is my regular expression engine created by modelling an (NFA) Non-Deterministic Finite automaton in the
Go programming language.

This project is part of my Graph Theory module in my 3rd year Software Development course at GMIT.

## Setup Instructions

You need to have [Go](https://golang.org/dl/) installed to run this project. You will also need to have your `GOPATH` set up.

if you're on a Ubuntu based system, you can run the commands

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
[This page](https://github.com/golang/go/wiki/SettingGOPATH) will provide instructions on how to set up everything on your
machine.
If you have git installed you can clone the repository here:
```bash
git clone https://github.com/DamianGavin/GraphTheoryNFAProject
```
You can also click on the "Clone or Download" button to download the project in a zip file.

Next you will need to navigate to the new directory
```bash
cd GraphTheoryNFAProject
```
and run the program.
```bash
go run main.go
```

##The code.
I have adapted my program from Two main sources, the first on being the videos supplied to me by my module lecturer
 [Ian Mcloughlin](https://github.com/ianmcloughlin)
 and the second being a paper written by Russ Cox.[this paper](https://swtch.com/~rsc/regexp/regexp1.html)

 The program consists of two algorithms that I have combined into one program. These are called the Shunting yard algorithm
 and Thompsons algorithm.
 # ShuntingYard
 Regular Expression Algorithm

 In computer science, the shunting-yard algorithm is a method for parsing mathematical expressions specified in infix notation.
 It can produce either a postfix notation string, also known as Reverse Polish notation (RPN), or an abstract syntax tree (AST).
 The algorithm was invented by Edsger Dijkstra and named the "shunting yard" algorithm because its operation resembles that of
 a railroad shunting yard. Dijkstra first described the Shunting Yard Algorithm in the Mathematisch Centrum report MR 34/61.

 Like the evaluation of RPN, the shunting yard algorithm is stack-based. Infix expressions are the form of mathematical notation
 most people are used to, for instance "3 + 4" or "3 + 4 × (2 − 1)". For the conversion there are two text variables (strings),
 the input and the output. There is also a stack that holds operators not yet added to the output queue. To convert, the program
 reads each symbol in order and does something based on that symbol. The result for the above examples would be "3 4 +" or
 "3 4 2 1 − × +".

 The shunting-yard algorithm was later generalized into operator-precedence parsing.

 #Thomson's construction algorithm.

 In Computer science this algorithm is for transforming a regular expression into an equivalent nondeterministic finite automaton(NFA).
 This can then be used to match strings against the regular expression.

 The algorithm works recursively by splitting an expression into its constituent subexpressions, from which the NFA will be constructed
 using a set of rules.

 My Program consists of Two Go files, main.go and nfa.go. main.go is the runner class and calls the MatchString function from nfa.go.
 The nfa.go file starts with 2 structs, one keeps track of the edges and the other keeps track of the initial state and the accept state.
 My 1st function is called poregtonfa. This function loops through the postfix regex one character at a time, It consists of a switch statement
  that sets out the rules for the special characters '.', '|', '+', '?', and '*'. My code is exhaustively commented so I refer you to the code
  for an in-depth explanation of each.
  I also have functions called addState, pomatch, infixToPostfix, and MatchString which is called from main.go.
