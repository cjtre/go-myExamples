/*
The purpose of this program is to show how we are able to nest structs
*/
package main

import "fmt"

type Salary struct {
	basic     int
	insurance int
	allowance int
}

type Employee struct {
	firstName, lastName string
	salary              Salary
	bool
}

func main() {
	ross := Employee{
		firstName: "Ross",
		lastName:  "Geller",
		bool:      true,
		salary:    Salary{1100, 50, 50},
	}

	fmt.Println(ross)
	fmt.Println(ross.salary.basic)
}
