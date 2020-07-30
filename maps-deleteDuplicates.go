//The purpose of this program is to show how we can eliminate duplicates using a map.
package main

import "fmt"



func writeToMap(s []int) {
	newSeq := make(map[int]int)
	for k, v := range s {
		newSeq[v] = k
	}

	for key, val := range newSeq {
		fmt.Printf("%d : %d\n", key, val)
	}
	
}
// func printMap(map [int]int) int {
// 	for _, val := range newSeq {
// 		fmt.Println(k)
// 	}
// }

func main(){

	seq := []int{20, 19, 12 , 1, 20}

	writeToMap(seq)


}