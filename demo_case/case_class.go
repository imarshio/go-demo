package demo_case

import "fmt"

const (
	// Male 常量，相当于一个static final
	Male = iota
	Female
	_
	Third
)

type Gender uint8

type User struct {
	// 结构体，相当于一个class
	// 在go当中没有类的概念，但是有结构体
	Name   string
	Age    int
	Gender Gender
}

// func (*User) run()
// 声明这是一个函数，名称是run()
// (*User) 代表传参是一个指针类型，具体对象类型是 User 类，只不过传的是对象的地址，即指针
// 这样的本质还是面对对象编程

func (*User) run() {
	fmt.Println("run")
}

func (User) sleep() {
	fmt.Println("sleep")
}

func UserCase() {

	//u := new(User)
	// 得到的是一个指针类型变量：var u *User = &User{}

	//u := User{}
	// 得到的是一个User类型变量，var u User = User{}

	u := &User{}
	// 得到的是一个指针类型变量：var u *User = &User{}
	// 如果想通过 User{} 方式获取一个指针类型，
	// 那么我们需要使用 & （取地址符号）来拿到对象的地址

	// 所以，指针类型的变量 == 指针指向的对象的地址

	u.run()
	u.sleep()
}
