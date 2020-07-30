//The purpose of this programme is to show how we can use the new() function
package main

import (
	"fmt"
)

func main(){

p:= new(int) // 
fmt.Println(*p)

*p = 2
fmt.Println(*p)
}

/*
The expression new(T) creates an unnamed variable of type T, initializes it to the zero value of T, and 
returns its address, which is of type *T.

Why is the new function rarely used?
Because most common unnamed variables are of struct types, for which the struct literal syntax
is more flexible.
*/