package main

import "fmt"

func main() {
	// 变量的定义
	//静态语言的变量和动态语言的变量定义差异大

	// 1. 最基础的变量定义
	//var i int
	//i = 20
	//fmt.Println(i)

	//2.定义并初始化
	//var i int = 10
	//fmt.Println(i)

	//3.根据值自动判断变量类型（类型推断）
	//var i = 100
	//fmt.Println(i)

	//4.使用 :=
	//i := 90
	//fmt.Println(i)

	// 总结：
	//var a int = 10
	//var b = 23
	//c := 45
	//fmt.Println(a, b, c)

	//多变量的定义
	//var a, b, c int
	//a, b, c = 10, 20, 30
	//fmt.Println(a, b, c)

	//var a, b, c int = 10, 20, 30
	//fmt.Println(a, b, c)

	//集合类型
	//var (
	//	a int
	//	b string
	//)
	//fmt.Println(a, b)

	//修改变量的值
	var i int = 10
	i = 20
	fmt.Println(i)

}
