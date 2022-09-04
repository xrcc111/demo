/* package main

import "fmt"

func main() {
	// getName("zhang", sayName)
	ff := calc("+")
	r := ff(10, 2)
	fmt.Printf("r: %v\n", r)
}

func sayName(name string) {
	fmt.Printf("name: %v\n", name)
}

func getName(name string, f func(string)) {
	f(name)
}

// 计算器
func calc(operate string) func(int, int) int {
	switch operate {
	case "+":
		return sum
	case "-":
		return sub
	default:
		return nil
	}
}
func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
*/