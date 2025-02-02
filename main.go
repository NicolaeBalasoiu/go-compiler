package main

import (
	"comp/repl"
	"fmt"
	"os"
	"os/user"
)

func main () {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the best programming language! \n", user.Username)
	fmt.Printf("Type in commands \n")
	repl.Start(os.Stdin, os.Stdout)
}