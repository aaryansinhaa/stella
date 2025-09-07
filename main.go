package main

import (
	"fmt"
	"os"
	"os/user"
	"stella/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Stella 🐶!\n",
		user.Username)
	fmt.Printf("An interpreted programming language named after my beloved dog!\n")
	repl.Start(os.Stdin, os.Stdout)
}
