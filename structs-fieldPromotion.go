/*
The purpose of this program is to show fields can be promoted in an anonymous nested struct. That is
to say that the nested struct fields are automatically available on the parent struct.
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
	Salary
}

func main() {
	ross := Employee{
		firstName: "Ross",
		lastName:  "Geller",
		Salary:    Salary{1100, 50, 50},
	} // notice - there is no shorthand for the struct literal syntax so neither of the following will compile:

	/*
		ross := Employee {"Ross", "Geller", 1100, 50, 50}
		or
		ross := Employee {firstName: "Ross", lastName: "Geller", Salary: {1100, 50, 50}}
	*/

	ross.basic = 1200
	ross.insurance = 0
	ross.allowance = 0
	fmt.Println("Ross' basic salary: ", ross.basic) // the basic field has been promoted
	fmt.Println("Ross is", ross)
}
