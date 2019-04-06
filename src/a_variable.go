// go 语言已包作为管理单位，每个文件必须先声明包，程序「必须有一个 main 包」
package main

import (
	"fmt"
)

/*
生成可执行文件，然后执行可执行文件

go build a_variable.go
./a_variable
*/

/*
直接运行

go run a_variable.go
*/

// 入口函数
func main() {
	fmt.Println("Hello World")
	aTest1()
	aTest2()
	aTest3()
	aTest4()
	aTest5()
	aTest6()
}

func aTest1() {
	fmt.Println("---------- 先声明变量，然后再初始化 ----------")
	var a int // 声明变量
	a = 10    // 给变量赋值

	var b, c string      // 同时声明多个相同类型的变量
	b, c = "bb", "ccc"   // 同时给多个变量赋值
	fmt.Println(a, b, c) // 10 bb ccc
}

func aTest2() {
	fmt.Println("---------- 声明变量并初始化 ----------")
	var a int = 10                // 声明变量并初始化
	var b, c string = "bb", "ccc" // 同时声明多个相同类型变量并初始化

	// 声明变量并赋初值时可以不写变量类型，根据值自动推导类型。可以为不同类型的变量
	var d = 10          // 声明变量并初始化
	var e, f = 20, "ff" // 同时声明多个变量并初始化

	// := 为特殊写法，表示声明变量并赋初值
	g := 10          // 声明变量并初始化
	h, i := 20, "ii" // 同时声明多个变量并初始化

	fmt.Println(a, b, c, d, e, f, g, h, i) // 10 bb ccc 10 20 ff 10 20 ii

	// fmt.Printf 格式化输出，%T 打印变量所属类型
	fmt.Printf("%T %T %T %T %T %T\n", d, e, f, g, h, i) // int int string int int string
}

func aTest3() {
	fmt.Println("---------- 交换多个变量的值，以及匿名变量 ----------")
	a, b := 10, 20
	fmt.Println(a, b) // 10 20
	// 交换多个变量的值
	a, b = b, a
	fmt.Println(a, b) // 20 10

	// _ 表示匿名变量，丢弃数据不处理，用于处理函数返回值。_ 不能作为 value
	var c int
	c, _ = a, b
	var d, _, _ = b, a, c
	fmt.Println(c, d) // 20 10
}

func aTest4() {
	fmt.Println("---------- 常量 ----------")
	const A int = 10
	const B, C int = 20, 30
	// 可以不写常量类型，根据值自动推导类型
	const D = 40
	const E, F = 50, "哈哈"
	fmt.Println(A, B, C, D, E, F)                       // 10 20 30 40 50 哈哈
	fmt.Printf("%T %T %T %T %T %T\n", A, B, C, D, E, F) // int int int int int string
}

func aTest5() {
	fmt.Println("---------- () 多个变量或常量 ----------")
	var (
		a int
		b float64
	)
	a, b = 10, 3.14
	var (
		c int     = 10
		d float64 = 3.14
	)
	// 可以不写变量类型，根据值自动推导类型
	var (
		e = 10
		f = 3.14
	)
	const (
		G int     = 10
		H float32 = 3.14
	)
	// 可以不写常量类型，根据值自动推导类型
	const (
		I = 10
		J = 3.14
	)
	fmt.Println(a, b, c, d, e, f, G, H, I, J) // 10 3.14 10 3.14 10 3.14 10 3.14 10 3.14
	// 可以只写一个，后面的值和第一相等
	const (
		L = 2
		M
		N
	)
	fmt.Println(L, M, N) // 2 2 2
}

func aTest6() {
	fmt.Println("---------- iota 常量自动生成器 ----------")
	// 1、iota 常量自动生成器，给常量赋值使用，每隔一行自动加 1
	const (
		a = iota
		b = iota
		c = iota
	)
	// 2、iota 遇到 const 时重置为 0
	const d = iota
	const e = iota
	fmt.Println(a, b, c, d, e) // 0 1 2 0 0

	const f, g, h = iota, iota, iota
	fmt.Println(f, g, h) // 0 0 0

	// 3、可以只写一个 iota，后面的值也为 iota，但会自动累加
	const (
		i = iota
		j
		k
	)
	fmt.Println(i, j, k) // 0 1 2

	// 4、如果是同一行，值都一样
	const (
		l    = iota
		m, n = iota, iota
		o    = iota
	)
	fmt.Println(l, m, n, o) // 0 1 1 2
}
