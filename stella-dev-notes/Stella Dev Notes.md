---
title: Stella Dev Notes
author: Aaryan Kumar Sinha(aaryansinhaiiit@gmail.com)
date: 2025-09-05T00:10:00
Version: v0.1
---
---

Stella will be interpreted by a "Tree Walking" interpreter.
>In simple words,  the interpreter parses the source code, builds an abstract syntax tree (AST) out of it and then evaluates this tree.

Stella will have it's own lexer, own parser, it's own tree representation and evaluator.

---
The  following is the feature list of this interpreter:
- C-like syntax  
- variable bindings  
- integers and booleans  
- arithmetic expressions  
- built-in functions  
- first-class and higher-order functions  
- closures  
- a string data structure  
- an array data structure  
- a hash data structure

## Variable usage:
- `set a = 1;`
- `set name = "Stella";`
- `set output = 10* 10/(5/10);`
## Arrays:
- `set myArray = [1,2,3,4];`
- `set associativeArray = {"name": "Aaryan", "language": "Stella"}`
- `myArray[0]; // => 1`
- `associativeArray["name"]; // => "Aaryan"`
## Functions:
**Defining function**
- `set add  = fn(a,b) {return a+b;}`
**Calling Function:**
- `add(2,3);`
**More Complex Function:** ```
``` c
set fibonacci = fn(x) {
	if (x == 0) {
		0
	} 
	else {
		if (x == 1) {
			1
		} else {
			fibonacci(x - 1) + fibonacci(x - 2);
		}
	}
};
```
**Higher Order Functions**:
*These are functions that take other functions as arguments*.
```c
set twice = fn(f, x) {  
	return f(f(x));  
};  
set addTwo = fn(x) {  
	return x + 2;  
};  
twice(addTwo, 2); // => 6
```
==Note: Functions in Stella are just values, like integers or strings.==

## Printing Something:
- `say("Hello, World!");`
## Taking inputs:
- `set input = ask("What is your name?");`

# Lexer:
## Purpose:
The lexer's job is to perform **lexical analysis**, which is the process of taking the raw text of our Monkey code and turning it into a stream of structured data known as **tokens**.

---

## Token:
```go
type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}
```
> We defined the `TokenType` type to be a string. That allows us to use many different values as `TokenTypes`, which in turn allows us to distinguish between different types of tokens.
- `Type` stores the type of the keyword.
- `Literal` is the sequence of characters for the token.
---

In the lexer, `position` is like your finger on the **current character** you're examining, while `readPosition` is the finger on the **next character**.
```go
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}
```
- `readPosition` is what helps us know when we've reached the end of the input in `readChar()`.
-  So, `position` is for the character we've already decided to process, and `readPosition` is for looking ahead.
- This "looking ahead" ability becomes really important when the lexer encounters characters that could be part of a two-character operator.

**==Note:==** The Lexer only supports ASCII characters instead of full UNICODE range to focus only on the core implementation of the interpreter.
> If you wish to support full UNICODE range, change `l.ch` from a `byte` to `rune` and change the way we read the next characters, since they could be multiple `bytes` wide now. Using `l.input[l.readPosition]` wouldn’t work anymore. And then we’d also need to change a few other methods and functions.
---

**==TODO: Implement support for UNICODE==**.

---
