package main

import (
	"fmt"
	"math"
)

var globalName = "xx" // 可以不使用

const constName string = "xx" // 定义就要赋值

func main() {
	// 先定义，再赋值
	var name string
	name = "xx"
	fmt.Println(name)

	// var 变量名 类型 = "变量值"
	var userName string = "xx"
	fmt.Println(userName)

	// 错误，常量不能修改
	//constName = "b"

	print()

}

func print() {
	fmt.Printf("%v\n", "你好")          // 可以作为任何值的占位符输出
	fmt.Printf("%v %T\n", "xx", "xx") // 打印类型
	fmt.Printf("%d\n", 3)             // 整数
	fmt.Printf("%.2f\n", 1.25)        // 小数
	fmt.Printf("%s\n", "xxx")         // 字符串
	fmt.Printf("%#v\n", "")           // 用go的语法格式输出，很适合打印空字符串
	fmt.Printf("%.0f\n", math.Pow(2, 63))
}
