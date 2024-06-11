package main

import (
	"fmt"
	"strconv"
)

func main() {

	// go 语言的编码统一都是utf-8
	// 双引号包裹的字符可以转义
	var name = "demo"
	fmt.Printf("hello %s\n", name)
	fmt.Printf("hello %T\n", name)

	// 反引号包裹不会转义
	var code = `fmt.Printf("hello %T\n", name)`

	fmt.Println("反引号字符", code)

	var concat = "abc" + "def"
	fmt.Println(concat)

	// 换行拼接需要将 + 放在最后
	var concat_a = "abc" + "def" +
		"abc" + "def"

	fmt.Println(concat_a)

	// 将 int 类型转成 string
	var num1 int = 20
	var str string = fmt.Sprintf("%d", num1)
	fmt.Printf("convert result str = %s, the type of str is %T\n", str, str)

	// 将 string 类型转成 int
	var num2, _ = strconv.ParseInt(str, 10, 64)
	fmt.Printf("convert result num2 = %d, the type of str is %T", num2, int8(num2))

}
