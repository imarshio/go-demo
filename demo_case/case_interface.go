package demo_case

import "fmt"

/**
 * @Description: 在go当中，接口的声明如下
 */

type Person interface {
	eat()
	sleep()
}

type Worder interface {
	work()
}

type Runner interface {
	run()
}

// 与Java不同的是，在go当中，继承接口不需要关键字
// 当一个类实现了接口的所有方法，那么这个类就继承了这个接口
// 如上我定义了一个带三个方法的接口

// Student 不工作,会eat,run,sleep所以他继承了Person 和 Runner 接口
type Student struct {
	// 结构体，相当于一个class
	Name   string
	Age    int
	Gender Gender
}

func (*Student) run() {
	fmt.Println("run")
}

func (*Student) eat() {
	fmt.Println("run")
}

func (*Student) sleep() {
	fmt.Println("sleep")
}

// worker 不会run，没有继承 Runner接口
type Worker struct {
}

func (*Worker) eat() {
	fmt.Println("run")
}

func (*Worker) sleep() {
	fmt.Println("sleep")
}

func (*Worker) work() {
	fmt.Println("sleep")
}

// Teacher 啥都会，继承了Person,Runner,Worker接口

type Teacher struct {
	// 结构体，相当于一个class
	Name   string
	Age    int
	Gender Gender
}

func (*Teacher) eat() {
	fmt.Println("run")
}

func (*Teacher) run() {
	fmt.Println("run")
}

func (*Teacher) sleep() {
	fmt.Println("sleep")
}

func (*Teacher) work() {
}
