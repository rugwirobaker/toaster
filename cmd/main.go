package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/rugwirobaker/toaster"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! welcome to toaster!\n",
		user.Username)
	fmt.Printf("Feel free to type in sql queries\n")
	fmt.Printf("To exit this prompt type in: '\\q'\n\n")

	toaster.Start(os.Stdin, os.Stdout)

	os.Exit(0)
}
