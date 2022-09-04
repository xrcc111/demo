// package main

// import "fmt"

// func main() {
// 	test()
// }

// func add() {
// 	var s = []string{}
// 	s = append(s, "number")
// 	s = append(s, "string")
// 	fmt.Printf("s: %v\n", s)
// }

// func remove() {
// 	var s = []int{1, 3, 4, 5, 6}
// 	s = append(s[:2], s[3:]...)
// 	fmt.Printf("s: %v\n", s)
// }

// func query() {
// 	s := []string{"nie", "ang", "wewe", "iiioio"}
// 	key := 3
// 	for i, v := range s {
// 		if key == i {
// 			fmt.Printf(" %v: %v\n", i, v)
// 		}
// 	}
// }

// func test() {
// 	var s = []int{10, 3, 4, 6, 8}
// 	m := make([]int, len(s))
// 	copy(m, s)
// 	s[0] = 100
// 	fmt.Printf("s: %v\n", s)
// 	fmt.Printf("m: %v\n", m)
// }
