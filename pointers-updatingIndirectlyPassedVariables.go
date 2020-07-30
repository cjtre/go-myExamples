package main

import "fmt"

func main(){

	v := 1
	incr(&v)
	fmt.Println(v)
}

func incr(p *int) int{
	*p++ // increments what p points to; does not change p itself
	return *p
}

/*
Because a pointer contains the address of a variable, passing a pointer argument to a function
makes it possible for the function to update the variable that was indirectly passed.
*/

