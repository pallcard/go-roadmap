package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	if age <= 0 {
		fmt.Println("未出生")
		return
	}
	if age <= 18 {
		fmt.Println("未成年")
		return
	}
	if age <= 35 {
		fmt.Println("青年")
		return
	}
	fmt.Println("中年")

	fmt.Println("请输入你的性别：")
	var sex string
	fmt.Scan(&sex)
	switch sex {
	case "男":
		fmt.Print("man")
		fallthrough //让程序执行下一个case语句，而不进行条件判断
	case "女":
		fmt.Print("woman")
	default:
		fmt.Println("默认")
	}

	var sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	list := make([]string, 10, 10)
	for i, _ := range list {
		list[i] = strconv.Itoa(i)
	}
	fmt.Println(list)

	m := make(map[string]string)
	m["1"] = "1"
	for k, v := range m {
		fmt.Print(k, v)
	}

}
