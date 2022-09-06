package main

import (
	"fmt"
	"os"
)

func main() {
	createFile()
	fileOpenClose()
}

func fileOpenClose() {
	f, err := os.Open("test.js")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}

}

func createFile() {
	file, err := os.Create("test.js")
	if err != nil {
		fmt.Printf("file.Name(): %v\n", file.Name())
	}
}
