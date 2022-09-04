/* package main

import "fmt"

// type USB interface {
// 	read()
// 	write()
// }

type Mobile struct {
	model string
}

type Computer struct {
	name string
}

func (c Computer) read() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("read.............")
}

func (c Computer) write() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("writr......")
}

func (m Mobile) read() {
	fmt.Printf("m.model: %v\n", m.model)
}
func main() {
	c := Computer{
		name: "macBook",
	}
	c.read()
	c.write()
	m := Mobile{
		model: "Iphone",
	}
	m.read()
}
*/