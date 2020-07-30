//This programme shows how we can sort maps
package main

import(
	"fmt"
	"sort"
)

var ages = map[string] int {
	"Luke": 23,
	"Simon": 22,
	"Theresa": 35,
}

func main(){

names := []string{}

for name := range ages {
	names = append(names, name)
}

sort.Strings(names)


for _, name := range names {
	fmt.Printf("%s\t%d\n", name, ages[name])
	}
}