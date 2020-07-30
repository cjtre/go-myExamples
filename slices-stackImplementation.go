//The purpose of this programme is to show how we can use a slice to implement a stack using append method
package main

import "fmt"

func main(){

	stack := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	val := "eleven"

	stack = append(stack, val) //push val onto the stack
	fmt.Println(stack)

	// The top of the stack is the last element.

	top := stack[len(stack)-1]
	fmt.Println(top)

	// Shrink the stack by popping the last element.
	stack = stack[:len(stack)-1] //pop's last element from stack

	fmt.Println(stack)



}