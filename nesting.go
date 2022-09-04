/* package main

import "fmt"

type Fly interface {
	fly()
}

type Swim interface {
	swim()
}

type FlySwim interface {
	Fly
	Swim
}

type User struct {
}

func (user User) fly() {
	fmt.Println("fly")
}

func (user User) swim() {
	fmt.Println("swim")
}

func main() {
	var ff FlySwim
	ff = User{}
	ff.fly()
	ff.swim()
}
*/