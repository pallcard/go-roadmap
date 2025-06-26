package main

import (
	"errors"
	"fmt"
	"time"
)

// 无返回值
func fun1() {
	return // 也可以不写
}

// 单返回值
func fun2() int {
	return 1
}

// 多返回值
func fun3() (int, error) {
	return 0, errors.New("错误")
}

// 命名返回值
func fun4() (res string) {
	return // 相当于先定义再赋值
	//return "abc"
}
func main() {
	//匿名函数
	var add = func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))

	var funcMap = map[int]func(){
		1: func() {
			fmt.Println("登录")
		},
		2: func() {
			fmt.Println("个人中心")
		},
		3: func() {
			fmt.Println("注销")
		},
	}
	funcMap[1]()

	fmt.Println(awaitAdd(2)(1, 2, 3))
}

func awaitAdd(t int) func(...int) int {
	time.Sleep(time.Duration(t) * time.Second)
	return func(numList ...int) int {
		var sum int
		for _, i2 := range numList {
			sum += i2
		}
		return sum
	}
}
