//The purpose of this application is to show how we can eliminate adjacent duplicates
package main

import "fmt"

func removeAdjDup(s []string) []string {
	
	for i:=0; i < len(s)-1; {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}

//Using for-loops to remove duplicates from a slice
var state bool

//checkDups checks to see if an element is present in a slice
func checkDups(el int, s []int) bool{
	for _, val := range s {
		if val == el {
			 return true
		}
	}
	return false
}

//removeDup iterates through a slice adding elements to a second slice only if checkDups returns false.
func removeDup(s []int) []int{
	for _, val := range seq{
		if !checkDups(val, newSeq){
			newSeq = append(newSeq, val)
		}
	}
	return newSeq
}

var seq = []int{1, 2, 3, 4, 4, 5, 6, 7, 1, 2, 3, 4, 8, 9, 10}
var newSeq = []int{}

func main(){

	//sl := []string{"one", "two", "three", "three", "five"}

	// fmt.Println(removeAdjDup(sl))

	fmt.Println(removeDup(seq))

}