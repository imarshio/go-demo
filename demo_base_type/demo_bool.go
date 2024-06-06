package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var b = false
	fmt.Println("b = ", b)
	fmt.Println("bool 类型占用的空间大小", unsafe.Sizeof(b))
}
