//The purpose of the this program is to demonstrate how structs can be used.
package main

import "fmt"

type fruit struct {
	fruitName string
	shape     string
	quantity  int
}

func main() {

	apple := fruit{
		"apple",
		"round",
		3,
	}

	pair := fruit{
		"pair",
		"oval",
		2,
	}

	listOfFruit := []fruit{apple, pair}

	for f := range listOfFruit {
		fmt.Println(listOfFruit[f].fruitName)
	}

}
