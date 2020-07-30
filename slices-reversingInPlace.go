//This program will demonstrate how to reverse a slice in-place
package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	myInts := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Println(myInts[len(myInts)-1])
	reverse(myInts[:])
	fmt.Println(myInts)
}
