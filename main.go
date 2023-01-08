package main

import (
	"fmt"
	"os"

	"github.com/tidwall/pretty"
)

func main() {
	// Read from standard input.
	arg := os.Args[1]
	fmt.Println(arg)
	jsonBytes := []byte(os.Args[1])
	// Pretty-print the JSON string.
	b := pretty.Pretty(jsonBytes)
	c := pretty.Color(b, nil)
	fmt.Println(string(c))
}
