package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jkulton/mkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s. This is the Monkey programming language.\n", user.Username)
	fmt.Printf("Feel free to type commands in this REPL\n")
	repl.Start(os.Stdin, os.Stdout)
}
