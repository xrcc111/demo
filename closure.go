/* package main

import "fmt"

func main() {
	// t := closure()
	// fmt.Printf("t(10): %v\n", t(10))
	// fmt.Printf("t(20): %v\n", t(20))
	// fmt.Printf("t(30): %v\n", t(30))

	f1, f2 := calc(9)
	fmt.Println(f1(10), f2(10))
	fmt.Println(f1(40), f2(30))
}

func closure() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func calc(base int) (func(int) int, func(int) int) {
	sum := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return sum, sub
} */
