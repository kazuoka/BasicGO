package main

import "fmt"

// Scope
// Global variable
var gVariable int = 500

func main() {
	lVariable := 40
	fmt.Println("Gobal=", gVariable)
	fmt.Println("Local=", lVariable)
	anotherFunction()
}
func anotherFunction() {
	lVariable := 20
	fmt.Println("Gobal=", gVariable)
	fmt.Println("Local=", lVariable)
}
