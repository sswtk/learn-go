package main

import "fmt"

var aa = 3
var ss = "Kkk"

var (
	bb = 5
	vv = "www"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) //%q :quot
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 4, 8, false, "ddd" // := 只能在函数内使用
	b = 0                            //重新赋值
	fmt.Println(a, b, c, s)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb, vv)
}
