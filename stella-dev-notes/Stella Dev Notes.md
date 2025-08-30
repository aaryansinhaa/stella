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
