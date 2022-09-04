/* package main

import "fmt"

// 嵌套结构体
func main() {
	type Dog struct {
		name  string
		age   int
		color string
	}

	type Person struct {
		dog  Dog
		name string
		age  int
		sex  string
	}

	dog := Dog{
		name:  "欢欢",
		age:   2,
		color: "yellow",
	}

	per := Person{
		dog:  dog,
		name: "小吴",
		age:  18,
		sex:  "male",
	}
	fmt.Printf("per: %v\n", per)
	fmt.Printf("per.dog.name: %v\n", per.dog.name)
}
*/