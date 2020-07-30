//The purpose of this programme is to show how we can compare two slices
package main

import "fmt"

func equals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i:= range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main(){

	sl1 := []string{"one", "two", "three"}
	sl2 := []string{"one", "two", "four"}

	fmt.Println(equals(sl1, sl2))
}