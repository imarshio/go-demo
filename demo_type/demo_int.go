package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int64 = 20
	fmt.Printf("the type of num is %T\n", num)
	fmt.Printf("the size of num is %d\n", unsafe.Sizeof(num))
}
