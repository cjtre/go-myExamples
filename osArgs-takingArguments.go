//The purpose of this program is to show how we can use os.Args to take arguments
package main

import(
	"fmt"
	"os"
)

func main(){
	
	s, sep := "", ""
	// os.Args[0] - the first element is the name of the command itself so we start from index 1.
	//Go uses half-open intervals- includes the first element, excludes last
	for _, args := range os.Args[1:] {
		s += sep + args // 
		sep = ""
	}
	fmt.Println(s)
}
	// line 15 inefficient as assignment += receives completely new contents with previous contents
	// garbage collected. If the amount of data is large, then this could be costly.
	// more efficient: fmt.Println(strings.Join(os.Args[1:], " "))