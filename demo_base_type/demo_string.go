package main

import "fmt"

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
}
