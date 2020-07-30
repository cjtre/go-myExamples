package main

import (
	"fmt"
	"strconv"
)

func main(){

	var num int = 10
	var str string = "10"

	conv := strconv.Itoa(num)
	fmt.Println(conv)

	val, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}
		
	fmt.Println(val)

}