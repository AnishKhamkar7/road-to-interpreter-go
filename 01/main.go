package main

import (
	"fmt"
	"go-int/src/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)

	}

	fmt.Printf("Hello %s! This is Monkey Programming Language!\n", user.Username)

	fmt.Printf("Feel Free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)

}
