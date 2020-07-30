//The purpose of this programme is to demonstrate how we can achieve deletion of elements in slices
package main

import (
	"fmt"
	"time"
)

//deletion by copying last element into element need deleting. Does not preserve order
func del1(num int, sl []string) []string {
	//Copy the last element into sl[i]
	sl[num] = sl[len(sl)-1]
	//remove last element
	sl[len(sl)-1] = ""
	return sl[:len(sl)-1] //truncate slice
}
// using Copy - preserving order
func del2(num int, sl []string) []string {
	copy(sl[num:], sl[num+1:])
	return sl[:len(sl)-1]
}

func del3(num int, sl []string) []string{
	sl[num] = sl[len(sl)-1] // take the last element of the slice and insert in sl[num]
	return sl[:len(sl)] // returns all elements of slice - this excludes last element.
	//This is pretty fast,however, the slice needs to be truncated.
}

func main() {

	start := time.Now()
	sli := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("Removing last elements method: \n%q", del1(1, sli))
	fmt.Println(time.Since(start))

	fmt.Println()

	start = time.Now()
	sli1 := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("Using copy method: \n%q ", del2(1, sli1))
	fmt.Println(time.Since(start))

	fmt.Println()

	start = time.Now()
	sli2 := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("Maintaining order \n%q ", del3(1, sli2))
	fmt.Println(time.Since(start))


}
