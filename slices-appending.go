//The purpose of this program is to show why we must append to a variable.
package main

import "fmt"

var runes []rune

func main(){

for _, r := range "hello"{
	runes = append(runes, r) // we must assign the append to the runes variable.
}

fmt.Printf("%q \n", runes)
//fmt.Println([]rune("hello"))
}

/*
Usually, we don't know whether a given call to append will cause a reallocation, so we can't
assume that the the original slice refers to the same array as the resulting slice, nor that it
refers to a different one. Similarly, we must not assume that operations on elements of the old slice 
will not be reflected in the new slice. As a result, it's common to assign the result of a call
to append to the same slice variable whose value we passed to append.
*/