package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	sourceString := "Hello, OTUS!"
	outputString := stringutil.Reverse(sourceString)
	fmt.Println(outputString)
}
