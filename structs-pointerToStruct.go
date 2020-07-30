//The purpose of this program is to show how we can create a pointer that points to the value of a struct (in just one statement). This is an alternative to creating a struct.
package main

import "fmt"

//the syntax to create a pointer to a struct

//s:= &StructType{...}

type Employee struct {
	firstName, lastName string
	salary              int
	fulltime            bool
}

func main() {
	ross := &Employee{
		firstName: "ross",
		lastName:  "Bing",
		salary:    1200,
		fulltime:  true,
	}

	fmt.Println("firstname", (*ross).firstName)
}

//ross is a pointer, so we need to use *ross dereferencing syntax to get the actual value the struct is pointing to. Now we can use (*ross).firstName to access the firstName variable.

//it is also worth noting that there is a shortcut: we can access the fields of a struct without dereferencing it first.

/*
ross := &Employee {
    firstName: "ross",
    lastName:  "Bing",
    salary:    1200,
    fullTime:  true,
}
fmt.Println("firstName", ross.firstName) // ross is a pointer
*/
