package main

import "fmt"

 var students = map[string] int{
	"Luke": 23,
	"Simon": 22,
	"Theresa": 35,
}

func main(){

fmt.Println(students)

//map has a built-in delete function
delete(students, "Theresa")

fmt.Println(students)
}