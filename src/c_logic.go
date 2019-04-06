package main

import (
	"fmt"
	"time"
)

func main() {
	cTest1()
	cTest2()
	cTest3()
	cTest4()
}

func cTest1() {
	fmt.Println("---------- if ----------")

	a := 1
	if a == 1 {
		fmt.Println("11aa")
	} else if a == 2 {
		fmt.Println("11bb")
	} else {
		fmt.Println("11cc")
	}

	// 支持一个初始化语句，初始化语句和判断条件以分号分隔
	if b := 1; b == 1 {
		fmt.Println("22aa")
	} else if b == 2 {
		fmt.Println("22bb")
	} else {
		fmt.Println("22cc")
	}
}

func cTest2() {
	fmt.Println("---------- switch ----------")
	// go 语言保留了 break 关键字，跳出 switch 语句。但是可以不写，默认就包含了
	a := 1
	switch a {
	case 1:
		fmt.Println("aa")
		fallthrough // 不跳出 switch 语句，后面紧挨着 case 的无条件执行
	case 2:
		fmt.Println("bb")
	case 3:
		fmt.Println("cc")
	default:
		fmt.Println("ee")
	}

	fmt.Println("----------")

	// 支持一个初始化语句，初始化语句和变量本身以分号分隔
	switch b := 5; b {
	case 1:
		fmt.Println("aa")
	case 2, 3, 4:
		fmt.Println("bb")
	case 5:
		fmt.Println("cc")
		fallthrough // 不跳出 switch 语句，后面紧挨着 case 的无条件执行
	default:
		fmt.Println("ee")
	}

	fmt.Println("----------")

	c := 60
	// switch 后面可以指定变量，在 case 后面指定条件
	switch {
	case c > 90:
		fmt.Println("优秀")
	case c > 80:
		fmt.Println("良好")
	case c > 70:
		fmt.Println("一般")
	case c >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}

func cTest3() {
	fmt.Println("---------- for ----------")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("原始 str[%d] = %c\n", i, str[i])
	}

	for i, item := range str {
		fmt.Printf("range index 和 value str[%d] = %c\n", i, item)
	}

	for i := range str {
		fmt.Printf("range index str[%d] = %c\n", i, str[i])
	}

	for i, _ := range str {
		fmt.Printf("range 忽略 value str[%d] = %c\n", i, str[i])
	}

	for _, item := range str {
		fmt.Printf("range 忽略 index %c\n", item)
	}

	i := 0
	for {
		i++
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func cTest4() {
	fmt.Println("---------- goto ----------")

	fmt.Println("111")
	goto End
	fmt.Println("222")
End:
	fmt.Println("333")
}
