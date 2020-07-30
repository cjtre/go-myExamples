//The purpose of this package is to demonstrate how elements can be read into a map and similar elements counted
package main 

import(
	"fmt"
	"bufio"
)

func main(){

	counts := make(map[string]int) //map stores key/value pairs
	input := bufio.NewScanner(os.Stdin)

	for input.Scan(){
		counts[input.Text()]++
	}

	for line, n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/*
1. The Scanner reads from the program's standard input.
2. Each call to input.Scan() reads the next line and removes the new line character from the end.
3. The result can be retrieved by calling input.Text()

The Scan function returns true if there is a line and false when there is no more input. 

A Map is a reference to the data structure created by make/

counts[input.Text()]++
Each time the program reads a line of input, the line is used as a key into the map and the corresponding
value is incremented. 
*/

