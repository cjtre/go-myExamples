package main

import "fmt"

var x int // &x - read as the address of x - yields a pointer to an integer variable - *int 

func main(){

x := 1
p := &x //p, of type *int (pointer to int), points to x
fmt.Println(*p) // 1

*p = 2 //equivalent to x
fmt.Println(x) // 2

}