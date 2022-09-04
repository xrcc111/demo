/* package main

import "fmt"

func main() {
	r := f2(10)
	fmt.Printf("r: %v\n", r)
}

func f1() {
	var s int = 1
	for i := 1; i <= 5; i++ {
		s *= i
	}
	fmt.Printf("s: %v\n", s)
}

func f2(a int) int {
	if a == 1 {
		return a
	} else {
		return a * f2(a-1)
	}
}
*/