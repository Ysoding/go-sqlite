package main

import (
	"bufio"
	"fmt"
	"os"
)

func printPrompt() {
	fmt.Printf("db > ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		scanner.Scan()
		line := scanner.Text()

		if line == ".exit" {
			os.Exit(0)
		} else {
			fmt.Println("Unrecognized command", line)
		}
	}
}
