package main

import (
	"fmt"
)

func main() {
	bTest1()
	bTest2()
	bTest3()
	bTest4()
	bTest5()
	bTest6()
	bTest7()
}

func bTest1() {
	fmt.Println("---------- 类型默认值 ----------")
	var a bool // 布尔类型 | 1 | false | 不能用数字代表 true 或 false
	var b byte // 字节类型 | 1 | 0 | uint8 别名
	var c rune // 字符类型 | 4 | 0 | 专用于存储 unicode 编码，等价于 uint32

	var d int  // 整型 | 4 或 8 | 0 | 32 位
	var e uint // 整型 | 8 或 8 | 0 | 64 位

	var f int8  // 整型 | 1 | 0 | -128 ~ 127
	var g uint8 // 整型 | 1 | 0 | 0 ~ 255

	var h int16  // 整型 | 2 | 0 | -32768 ~ 32767
	var i uint16 // 整型 | 2 | 0 | 0 ~ 65535

	var j int32  // 整型 | 4 | 0 | -21亿 ~ 21亿
	var k uint32 // 整型 | 4 | 0 | 0 ~ 42亿

	var l int64  // 整型 | 8 | 0
	var m uint64 // 整型 | 8 | 0

	var o float32 // 浮点型 | 4 | 0.0 | 小数点精确到 7 位
	var p float64 // 浮点型 | 8 | 0.0 | 小数点精确到 15 位

	var q complex64  // 复数类型 | 8 | (0+0i)
	var r complex128 // 复数类型 | 16 | (0+0i)

	var s uintptr // 整型 | 4 或 8 |  | 足以存储指针的 uint32 或 uint64
	var t string  // 整型 | | "" | utf-8 字符串

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, o, p, q, r, s, t)
}

func bTest2() {
	fmt.Println("---------- 测试自动推导的类型 ----------")
	// 字符类型自动类型推导为 int32
	a := 'A'

	var b rune
	b = 'a'

	var c byte = 1
	// 整型自动类型推导为 int
	d := 2
	// 浮点型自动类型推导为 float64
	e := 3.14

	fmt.Printf("%T %T %T %T %T\n", a, b, c, d, e) // int32 int32 uint8 int float64
}

func bTest3() {
	fmt.Println("---------- 字符 ----------")
	var ch byte = 97
	fmt.Printf("%c %d\n", ch, ch) // a 97
	ch = 'a'
	fmt.Printf("%c %d\n", ch, ch)    // a 97
	fmt.Printf("%d %d\n", 'A', 'a')  // 65 97
	fmt.Printf("小写转大写：%c\n", 'a'-32) // 小写转大写：A
	fmt.Printf("大写转小写：%c\n", 'A'+32) // 大写转小写：a

	str := "BGA" // 字符串都是隐藏了一个结束符，'\0'
	fmt.Printf("%T %d\n", str, len(str))
	fmt.Printf("str[0] = %c\n", str[0])
}

func bTest4() {
	fmt.Println("---------- 复数类型 ----------")
	var a complex128
	a = 2.1 + 3.14i
	fmt.Println(a) // (2.1+3.14i)

	b := 3.3 + 4.4i
	fmt.Printf("%T\n", b) // complex128

	fmt.Println("实部 =", real(b), "虚部 =", imag(b)) // 实部 = 3.3 虚部 = 4.4
}

func bTest5() {
	fmt.Println("---------- 类型转换 ----------")
	var a bool = true
	fmt.Printf("a = %d\n", a) // a = %!d(bool=true)
	fmt.Printf("a = %t\n", a) // a = true

	// bool 和 int 不能互转

	var b = 'a'
	var c int
	c = int(b)
	fmt.Println("b =", b, "c =", c) // b = 97 c = 97
}

func bTest6() {
	fmt.Println("---------- 类型别名 ----------")
	type bigint int64
	var a bigint          // 等价于 var a int64
	fmt.Printf("%T\n", a) // main.bigint

	type (
		long   int64
		double float64
	)
}

func bTest7() {
	fmt.Println("---------- 终端输入 ----------")
	var name string
	fmt.Print("请输入姓名：")

	// & 取地址运算符
	// * 取值运算符，指针变量所指向内存的值

	// 阻塞等待用户的输入
	fmt.Scanf("%s", &name)

	var age int
	fmt.Print("请输入年龄：")
	fmt.Scanf("%d", &age)

	fmt.Println("姓名 =", name, "年龄 =", age)

	fmt.Print("请输入姓名：")
	fmt.Scan(&name)

	fmt.Print("请输入年龄：")
	fmt.Scan(&age)

	fmt.Println("姓名 =", name, "年龄 =", age)
}
