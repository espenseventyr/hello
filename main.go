// hello/main.go
// Just for fun
package main

import (
	"fmt"
	"github.com/espenseventyr/stringutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("No argument! Or too many!", stringutil.Reverse("!sdrawkcab si sihT"))
		fmt.Println(stringutil.Reverse("1 2 3 4 5 6 7 8 9 0"))
		os.Exit(1) // Exit with status code 1 (no argument, or too many).
	}
	// It is idomatic in Go to omit the else when the if statement's body has
	// a break, continue, return or, goto.

	fmt.Printf("Great! You wrote %s!\n", os.Args[1])
}
