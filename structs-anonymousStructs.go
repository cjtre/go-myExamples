//The purpose of this program is to show how we can create and use anonymous sctructs
package main

import "fmt"

func main() {

	// Here, we are creating a struct names without defining a derived
	// struct type. This is useful when you don't want to re-use a struct
	// type.
	Chris := struct {
		Name string
		Age  int
	}{
		Name: "Chris",
		Age:  35,
	}

	fmt.Println(Chris)
}

// It is not clear when using an anonymous function what type Chris is.
// To find the type, we can use the fmt.Printf function and the %T format syntax

// use a derived type when wanting to re-use structs
