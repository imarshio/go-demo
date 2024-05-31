package main

import "fmt"

func main() {
	VoidMethod()
	NoReturnMethod(1, 2)
	fmt.Println(Sum(1, 2))
	fmt.Println(Reverse(1, 2))
}

func VoidMethod() {
	println("this is a void method with no variables and no return value")
}

func NoReturnMethod(a, b int) {
	println("this is a method with two variables and no return value")
}

func Sum(a, b int) int {
	println("this is a method with two variables and one return value")
	return a + b
}

func Reverse(a, b int) (c, d int) {
	fmt.Println("this is a method with two variables and two return values")
	println("this is a method with two variables and two return values")
	return b, a
}
