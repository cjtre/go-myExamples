//The purpose of this program is to demonstrate how fields of a struct can be modified using pointers.
package main

import "fmt"

//define struct
type Employee struct {
	firstName string
	lastName  string
	age       int
	position  string
}

//create a modify function that takes in a pointer to a type
func modifyEmployee(emp *Employee, str string) {
	//first it is necessary to check for a nill
	if emp == nil {
		fmt.Println("emp is nil")
		return
	}

	emp.position = str

}

func main() {

	Katie := Employee{
		"Katie",
		"Colwoth",
		35,
		"General Specialist",
	}
	//call function
	modifyEmployee(&Katie, "General Manager")
	fmt.Println(Katie.position)

}
