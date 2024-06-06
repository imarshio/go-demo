package main

import (
	"fmt"
	"unsafe"
)

func DemoInt() {

	var num int64 = 20
	fmt.Printf("the type of num is %T\n", num)
	fmt.Printf("the size of num is %d\n", unsafe.Sizeof(num))

	// go 默认整数数值类型为int
	var num0 = 10
	fmt.Println("num0 =", num0)
	fmt.Printf("the type of num0 is %T\n", num0)

	// 指定类型int
	var num1 int = 10
	fmt.Println("num1 =", num1)
	fmt.Printf("the type of num1 is %T\n", num1)
	fmt.Printf("the size of num1 is %d bytes\n", unsafe.Sizeof(num1))

	// 8代表一个字节 int8 表示有符号支持八个字节的数字
	// 最大为 0111 1111 =
	// 最小为 1111 1111 =
	var num2 int8 = 10
	fmt.Println("num2 =", num2)
	fmt.Printf("the type of num2 is %T\n", num2)
}
