package main

import "fmt"

func demo_byte() {
	// 英文字符
	var s byte = 'a'
	// 直接输出会输出这个字符的码值
	fmt.Println("直接输出：", s)
	fmt.Printf("输出字符： %c\n", s)
	fmt.Printf("输出字符类型： %T\n", s)

	// 中文字符
	//var z byte = '中',会报错，overflow
	var z int = '中'
	fmt.Println("直接输出：", z)
	fmt.Printf("输出字符： %c\n", z)
	fmt.Printf("输出字符类型： %T\n", z)

	// 字符类型可以直接运算
	var n = 10 + 'a'
	fmt.Println("直接输出：", n)
	fmt.Printf("输出字符： %c\n", n)
	fmt.Printf("输出字符类型： %T\n", n)

}
