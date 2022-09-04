/* package main

import "fmt"

func main() {
	type f1 func(int, int) int
	var demo f1
	demo = sum
	r := demo(10, 10)
	fmt.Printf("r: %v\n", r)
	demo = max
	m := demo(1000, 10)
	fmt.Printf("m: %v\n", m)
}

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
*/