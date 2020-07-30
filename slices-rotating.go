package main

import "fmt"
import "time"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, el int) []int{
	reverse(s[:el])
	reverse(s[el:])
	reverse(s)
	return s
}

func rotateMod(s []int, el int) []int {
	num := el % len(s)
	double:= append(s, s[:num]...) 
	copy(s, double[num: num+len(s)])
	return s
}

func main(){

	s := []int {0, 1, 2, 3, 4, 5}
	//rotate left by 2 positions

	start := time.Now()
	fmt.Println(rotate(s, 2))
	fmt.Println(time.Since(start))

	fmt.Println()

	start = time.Now()
	fmt.Println(rotateMod(s, 2))
	fmt.Println(time.Since(start))

}

/*
{1, 0, 2, 3, 4, 5}
{1, 0, 5, 4, 3, 2}
{2, 3, 4, 5, 0, 1}
*/