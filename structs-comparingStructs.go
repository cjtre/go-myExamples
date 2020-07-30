// the purpose of this program is to show that structs are comparable:
package main

import "fmt"

type Point struct{ X, Y int }

func main() {
	p := Point{1, 2}
	q := Point{2, 1}

	fmt.Println(p.X == q.X && p.Y == q.Y) // false
	fmt.Println(p == q)
}
