package main

import (
	"fmt"
	"math"
)

/** go语言的基本数据类型有
1.整数形
2.浮点型
3.复数
4.布尔
5.字符串
**/

var b byte // uint8别名
var n1 uint8 = 2
var n2 uint16 = 2
var n3 uint32 = 2
var n4 uint64 = 2
var n5 uint = 2
var n6 int8 = 2
var n7 int16 = 2
var n8 int32 = 2
var n9 int64 = 2
var n10 int = 2 // int按照机器长度 32/64

var f float32
var f2 float64 //默认

func main() {
	// 整型 浮点型
	fmt.Printf("\n\n-------整型 浮点型------\n\n")

	fmt.Printf("%.0f\n", math.Pow(2, 63))
	var n1 int = 9223372036854775807
	fmt.Println(n1)
	//var n2 int = 9223372036854775808 // 执行报错
	//fmt.Println(n2)
	//var u int8 = 128 // 报错
	//fmt.Print(u)

	var i1 = 1
	fmt.Printf("%T\n", i1) //int
	var f1 = 1.0
	fmt.Printf("%T\n", f1) //float64

	// 字符型 字符的本质是一个整数
	fmt.Printf("\n\n-------字符型------\n\n")
	var c1 = 'a'
	var c2 = 97
	fmt.Println(c1) // 直接打印都是数字
	fmt.Println(c2)

	fmt.Printf("%c %c\n", c1, c2) // 以字符的格式打印

	var r1 rune = '中' //  \u4e2d ---->十六进制 0x4E2D = 十进制 20013
	fmt.Printf("%c\n", r1)
	fmt.Println(r1) //Unicode码

	var s string = "爱你中国"
	fmt.Println(s)
	fmt.Println("\t")   // 制表符
	fmt.Println("\n")   // 回车
	fmt.Println("\"\"") // 双引号
	fmt.Println("\r")   // 回到行首
	fmt.Println("\\")   // 反斜杠

	//转义字符原样输出
	s = `
a
b
c
`
	fmt.Print(s)

	fmt.Printf("\n\n-------布尔类型------\n\n")
	fmt.Print(true)
	fmt.Print(false)

	fmt.Printf("\n\n-------数组------\n\n") //长度固定
	var arr [2]int = [2]int{1, 2}
	fmt.Print(arr)
	fmt.Print(len(arr))
	var arr2 = [...]int{1, 2, 3, 4, 5} //等于[5]int
	fmt.Print(arr2)
	fmt.Print(len(arr2))
	arr2 = [5]int{5, 4, 3} //只能辅助为 [5]int
	fmt.Print(arr2)

	fmt.Printf("\n\n-------切片------\n\n") //长度不固定
	var list []int
	fmt.Print(list)
	fmt.Print(len(list))
	list = append(list, 1, 2, 4)
	fmt.Print(list)
	fmt.Print(len(list))
	list[2] = 1
	fmt.Print(list)

	var list2 []string //nil
	fmt.Print(list2)
	fmt.Print(list2 == nil) //true
	fmt.Printf("%T\n", list2)
	list2 = append(list2, "") //append对nil切片自动初始化

	list22 := []string{}
	fmt.Print(list22 == nil) //false

	var list3 []struct{ a int }
	fmt.Print(list3)
	fmt.Printf("%T\n", list3)
	list3 = append(list3, struct{ a int }{a: 3})
	fmt.Print(list3)

	list4 := make([]int, 1, 10)
	fmt.Print(list4)
	fmt.Print(len(list4))
	fmt.Print(list4 == nil)

	fmt.Printf("\n\n-------map------\n\n") //map使用之前一定要初始化
	// 声明
	var m1 map[string]string
	fmt.Print(m1 == nil)
	//m1["a"] = "1" //panic: assignment to entry in nil map
	// 初始化1
	m1 = make(map[string]string)
	// 初始化2
	m1 = map[string]string{}
	// 设置值
	m1["name"] = "xx"
	fmt.Println(m1)
	// 取值
	fmt.Println(m1["name"])
	// 删除值
	delete(m1, "name")
	fmt.Println(m1)

	// 声明并赋值初始化
	var m2 = map[string]string{}
	var m3 = make(map[string]string)
	m2["a"] = "1"
	m3["a"] = "1"

	a, ok := m2["a"]
	fmt.Println(a, ok)

	b, ok := m2["b"]
	fmt.Println(b, ok) //b为零值

	fmt.Printf("\n\n-------零值问题------\n\n")
	var a1 int
	var a2 float32
	var a3 string
	var a4 bool

	fmt.Printf("%#v\n", a1)
	fmt.Printf("%#v\n", a2)
	fmt.Printf("%#v\n", a3)
	fmt.Printf("%#v\n", a4)
}
