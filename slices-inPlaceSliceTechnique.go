package main

import "fmt"

func nonEmpty(strings []string)[]string{
	i:= 0
	for _, s := range strings{
		if s!="" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main(){

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonEmpty(data)) // ["one", "three"]
	fmt.Printf("%q\n", data)

}