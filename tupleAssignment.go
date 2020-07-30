//The purpose of this programme is to show how we can implement tuple assignment in Go.
package main

import "fmt"

func main(){

	i := 1
	j := 2

	i, j = j, i

	fmt.Println(i)
	fmt.Println(j)


	fmt.Println("-----------------------------------------------")

	i, j = j, i+1
	fmt.Println(i)
	fmt.Println(j)
}