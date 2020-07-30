package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FirstName string
	LastName  string
}

func main() {

	
	emp1 := Employee{
		FirstName: "George",
		LastName:  "Cunningham",
	}

	x, err := json.Marshal(emp1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(x))
}
