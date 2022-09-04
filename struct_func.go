/* package main

import "fmt"

type Person struct {
	name string
}

type User struct {
	name  string
	money int
}

type Login struct {
	username string
	password string
}

func main() {
	per := Person{
		name: "小吴",
	}
	per.doing()

	user := User{
		name:  "小何",
		money: 100,
	}
	user.pay()

	login := Login{
		username: "xu",
		password: "123abc",
	}
	fmt.Printf("%v\n", login.login(login.username, login.password))
}

// 属性和方法分开写
func (per Person) doing() {
	fmt.Printf("%v学习中....", per.name)
}

func (user User) pay() {
	fmt.Printf("%v花了%v元", user.name, user.money)
}

func (login Login) login(username string, password string) bool {
	if username == "xu" && password == "123abc" {
		return true
	} else {
		return false
	}
}
*/