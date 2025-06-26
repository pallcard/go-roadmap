package main

import "fmt"

import _ "go-roadmap/base/init" //导入init包，其init先执行

func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}
func init() {
	fmt.Println("init3")
}

// defer离return越近约先执行
func main() {
	fmt.Println("main")
	defer fmt.Println("defer4")
	Func()
	defer fmt.Println("defer3")
}

func Func() {
	defer fmt.Println("defer2")
	fmt.Println("func")
	defer fmt.Println("defer1")
}
