package main

import "fmt"

func main() {

	num1 := 10

	// 声明一个指针类型
	var num2 = &num1

	// 指针类型直接输出是一个地址
	fmt.Println(num2)
	// 想要取到指针指向的值 需要配合* 使用
	fmt.Println(*num2)

}
