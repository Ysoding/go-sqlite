package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Ysoding/go-sqlite/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is tiny sqlite!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
