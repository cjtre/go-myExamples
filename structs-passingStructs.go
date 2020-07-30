//The purpose of this program is to show how a struct can be passed into and return a struct
package main

import "fmt"

type Score struct {
	team1 int
}

func updateScore(sc Score, t1 int) int {
	sc.team1 = t1
	return sc.team1
}

func main() {

	s1 := Score{
		22,
	}

	fmt.Println(updateScore(Score{s1.team1}, 10))
	fmt.Println(s1.team1)
}
