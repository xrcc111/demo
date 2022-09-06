/* package main

import (
	"fmt"
	"os"
)

func main() {
	getRecord()
}

// 创建文件
func createFile() {
	file, err := os.Create("test.js")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("file.Name(): %v\n", file.Name())
	}
}

// 删除文件
func removeFile() {
	err := os.Remove("test.js")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 获取当前目录
func getRecord() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
}
*/