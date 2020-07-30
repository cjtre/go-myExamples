//The purpose of this programme is to demonstrate how two maps can be compared.
package main

import "fmt"

func compareMap(a, b map[string] int) bool{

	if len(a) != len(b){
		return false
	}

	for k, v := range a {
		if b[k] != v{
			return false
		}
	}
	return true
}

func main(){

	mp := map[string] int {
		"Thomas": 1,
		"Lucy": 2,
	}

	mp1 := map[string] int {
		"Thomas": 1,
		"Lilly": 2,
	}

	fmt.Println(compareMap(mp, mp1))
}

/*
As with slices, maps cannot be compared to each other.
*/