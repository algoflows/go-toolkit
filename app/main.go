package main

import (
	"fmt"

	"github.com/algoflows/toolkit"
)

func main() {
	var tools toolkit.Tools // Instantiate the module

	s := tools.RandomString(10) // Call the RandomString method
	fmt.Println("Random string:", s)
}
