package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {

	var a, b = 3, 4
	var c, d = 4.0, 5.9

	fmt.Println(GetMaxInt(a, b))

	fmt.Println(GetMaxFloat(c, d))

	fmt.Println(GetSumInt(a, b))

	fmt.Println(GetSumFloat(c, d))

	fmt.Println(GetSum(&a, &b))

	fmt.Println(GetMaxWith(a, b))
	// 由编译器推断泛型的类型
	fmt.Println(GetMax(a, b))
	// 显示的指定泛型类型
	fmt.Println(GetMax[float64](c, d))

	fmt.Println(Compare(a, b))

	Print(a)

	// 声明一个 map
	//var m map[uint8]any
	m := make(map[uint8]uint8)
	m[1] = 1
	m[2] = 2

	var list = MapToList(m)

	ch := make(chan uint8)
	go PrintChan(ch)
	for _, u := range list {
		ch <- u
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}

func GetSumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetSumFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// GetSum 指针类型的泛型需要用 interface {} 包起来
func GetSum[T interface{ *int | *float64 }](a, b T) T {
	//return a + b

	fmt.Println(a)
	fmt.Println(b)
	return nil
}

func GetMaxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func GetMaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetMax[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func GetMaxWith[T CustomNumberType](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type CustomNumberType interface {
	// ～ 表示支持int64及其衍生类型
	// ｜ 表示去并集
	// 多行之间取交集
	uint8 | int32 | float64 | ~int64 | int | DerivativeInt64
	uint16 | int32 | float64 | ~int64 | int | AliasInt32
}

// DerivativeInt64 int64的衍生类型
type DerivativeInt64 int64

// AliasInt32 = 代表将AliasInt32 设置为 int32的别名
type AliasInt32 = int32

// Compare comparable 内置泛型类型,只支持 == != 两个操作
func Compare[T comparable](a, b T) bool {
	return a == b
}

// Print any 内置泛型类型 是一个interface，支持任何类型
func Print[T any](a T) {
	fmt.Println(a)
}

// MapToList map[string]T 形参类型
// map[string]T 表示这是一个key为字符串，value为任意类型的映射（在Java中，这类似于 Map<String, T>）
func MapToList[C comparable, T any](m map[C]T) []T {
	var list []T
	for _, v := range m {
		list = append(list, v)
	}
	return list
}

// PrintChan any 内置泛型类型 是一个interface，支持任何类型
func PrintChan[T any](ch chan T) {
	fmt.Println("PrintChan")
	for data := range ch {
		fmt.Println(data)
	}
}

// List 定义泛型类型
type List[T any] []T

// ID 定义泛型接口
type ID interface {
	int | int8 | int16 | int32 | int64
}
