//The purpose of this programme is to show how we can reverse a slice of ints in place
package main

import "fmt"

func reverse(s []int){
	for i, j := 0, len(s)-1; i<j; i,j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)  
}

func main(){

	sl := []int {10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	reverse(sl)

}
