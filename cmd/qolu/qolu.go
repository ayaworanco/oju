package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Qolu CLI")
	// rules
	// 1. the first argument always will be a query
	// 2. queries needs to be in a simple quote -> ''
	query := os.Args[1]

	character := "'"

	if !strings.HasPrefix(string(query[0]), character) || !strings.HasSuffix(string(query[len(query)-1]), character) {
		fmt.Println("query is invalid")
	}

	fmt.Println(query)
	// Step 1.
	// get the query and break down to a valid query
}
