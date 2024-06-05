# Go-demo

Demo project for Go.

Currently, the version of Go is 1.22

## 快速开始

[Go Get Started](https://go.dev/doc/)

### Installing Go

[Install Go](https://go.dev/doc/install)

> [!Note]
> 注：最新版的Go下载已经不需要手动配置环境变量了，前提是你使用的是安装包安装。
>
> 如果是解压缩的安装包，还是需要手动配置环境变量。

### Getting Started

## Go env 命令

[GOPATH](https://pkg.go.dev/cmd/go#hdr-GOPATH_and_Modules)

```shell
go env

[demo.DESKTOP-xxxxxxV] ➤ go env
set GO111MODULE=
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\shenqing\AppData\Local\go-build
set GOENV=C:\Users\shenqing\AppData\Roaming\go\env
set GOEXE=.exe
set GOEXPERIMENT=
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=C:\Users\shenqing\go\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=C:\Users\shenqing\go
set GOPRIVATE=
set GOPROXY=https://proxy.golang.org,direct
set GOROOT=C:\Env\Go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLCHAIN=auto
set GOTOOLDIR=C:\Env\Go\pkg\tool\windows_amd64
set GOVCS=
set GOVERSION=go1.22.3
set GCCGO=gccgo
set GOAMD64=v1
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=0
set GOMOD=NUL
set GOWORK=
set CGO_CFLAGS=-O2 -g
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-O2 -g
set CGO_FFLAGS=-O2 -g
set CGO_LDFLAGS=-O2 -g
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -fno-caret-diagnostics -Qunused-arguments -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=C:\Users\shenqing\AppData\Roaming\MobaXterm\slash\mx86_64b\tmp\go-build4034802447=/tmp/go-build -gno-record-gcc-switches

```

### 切换国内代理

```shell
# 查看默认的代理配置
go env GOPROXY

# https://proxy.golang.org,direct

# 设置代理
go env -w GOPROXY=https://goproxy.cn,direct

# 查看代理配置
go env GOPROXY

# https://goproxy.cn,direct
```

### goproxy

[goproxy](https://goproxy.cn/)

下面是 goproxy 官网提供的一个（zhuang bility）炫酷的演示案例。

```go
package main

import (
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
	// 	require "github.com/goproxy/goproxy@v1.6.9"
)

func main() {
	http.ListenAndServe("localhost:8080", &goproxy.Goproxy{
		GoBinEnv: append(
			os.Environ(),
			"GOPROXY=https://goproxy.cn,direct", // 使用 Goproxy.cn 作为上游代理
			"GOPRIVATE=git.example.com",         // 解决私有模块的拉取问题（比如可以配置成公司内部的代码源）
		),
		ProxiedSUMDBs: []string{
			"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org", // 代理默认的校验和数据库
		},
	})
}
```

## Go mod 命令

### Go mod 文件组成

```mod
// go.mod 文件是 go 用于依赖管理的文件
// 可以通过 go mod init 来创建该文件

// 模块名称
module go-demo

// go sdk version
go 1.22

// 第三方依赖
require (
	// dependency latest
)

// 排除以来继承里不需要的依赖
exclude (
	// dependency latest
)

// 修改依赖包的路径
// 依赖包发生迁移
// 原始包无法使用
// 使用本地包

// 即当依赖包不可用时，可以使用 replace 替换依赖包
replace (
	// source latest => target latest
)

// 撤回
// 如，当前项目作为其他项目的依赖，如果当前版本出现问题，则回退到指定的版本
retract (
)
```

### 创建一个项目

```shell
# 初始化模块（项目）
go mod init

# 初始化指定模块（项目）名称
go mod init demo
```

### 下载依赖

```shell

go mod download

go mod download github.com/gin-gonic/gin@v1.10.0

# 这里下载不会有任何提示，说明下载成功，有提示则不成功
# 下载位置在 GOPATH/pkg/mod/
# 只是下载指定的包，不会下载指定包依赖的包
```

### Go mod tidy

remove unused dependencies and add missing dependencies.

可以理解为对 go.mod 文件进行格式化整理

```shell
go mod tidy
```

## Go build 命令

## Go run 命令

## Go 开发规范

[Go Code Style](https://gocn.github.io/styleguide/)

以下是我个人精简后的规范

### package命名

package 名称应该与文件名一致（虽然 go 可以使 package 名与文件名不一致，但是不建议这样做），并且以小写字母开头。

使用下划线命名法(Underscore Naming Convention)， 在多个单词之间使用下划线作为分隔符，每个单词全部小写。

只有包名为`main`时，才可以通过`go run`命令直接运行。

### 变量命名

#### 全局变量

#### 局部变量

### 变量声明

- 没用的包不要import，在Go中，没用的（unused）变量会爆红
- 没用的变量不要声明，在Go中，没用的（unused）变量会爆红

### 函数命名

驼峰命名法(Camel Case Naming Convention)，在多个单词之间使用驼峰命名法，每个单词的首字母大写（包括开头首字母），其余字母小写。

## Go 数据类型

[Go Type](https://golang.org/ref/spec#Types)

### 基础类型

Go语言中，基础类型有如下这些

```textmate
<!--  布尔 -->
bool
<!--  字符-->
string

<!--  有符号整数-->
int  int8  int16  int32  int64

<!--  无符号整数-->
uint uint8 uint16 uint32 uint64 uintptr

<!--  字节-->
byte // alias for uint8

<!--  -->
rune // alias for int32
     // represents a Unicode code point

<!--  浮点型-->
float32 float64

<!--  -->
complex64 complex128
```

### int 类型

| 类型    | 含义                                                                                                                          |
|-------|-----------------------------------------------------------------------------------------------------------------------------|
| int   | 这是一个平台相关的整数类型，它的大小依赖于目标架构。在32位系统上，int通常等于int32；在64位系统上，则通常等于int64。它提供了对机器字长的良好默认选择，适用于大多数整数运算需求。                            |
| int8  | 这是一个8位有符号整数，范围是从-128到127。当你确切知道变量的值会在这个范围内，并且想要节省空间时，可以使用它。                                                                 |
| int16 | 这是一个16位有符号整数，范围是从-32,768到32,767。适合存储比int8更大一些的整数值。                                                                          |
| int32 | 这是一个32位有符号整数，范围是从-2^31到2^31-1（即-2,147,483,648到2,147,483,647）。当需要更大的整数范围，但又不需要64位的大小时，可以使用它。                                 |
| int64 | 这是一个64位有符号整数，范围是从-2^63到2^63-1（即-9,223,372,036,854,775,808到9,223,372,036,854,775,807）。它提供了极大的整数范围，适用于需要大量精度的场景，比如时间戳或者大规模计数。 |

### 复杂类型

```textmate
指针
Pointer

数组

结构体
struct

管道
channel

函数

切片
slice

接口
interface

map
```

## Go 变量声明

```go
package main

import "fmt"

// 全局变量声明
var num1 = 10
var name = "demo"

func main() {
	a := 10
	fmt.Println(a)
	fmt.Println(num1 + 1)
	fmt.Println(name)
}

```

## Go 函数声明

### main function

```go
package main

// 主函数声明，一个 module 只能有一个 main function
func main() {
}

```

### 一般函数声明

```go
package main

import "fmt"

func main() {
	VoidMethod()
	NoReturnMethod(1, 2)
	fmt.Println(Sum(1, 2))
	fmt.Println(Reverse(1, 2))
}

func VoidMethod() {
	fmt.Println("this is a void method with no variables and no return value")
}

func NoReturnMethod(a, b int) {
	fmt.Println("this is a method with two variables and no return value")
}

func Sum(a, b int) int {
	fmt.Println("this is a method with two variables and one return value")
	return a + b
}

func Reverse(a, b int) (c, d int) {
	fmt.Println("this is a method with two variables and two return values")
	println("this is a method with two variables and two return values")
	return b, a
}


```

- func 关键字
- VoidMethod 函数名
- (a, b int) 参数列表、参数名称、参数类型
- int （第二个int）返回值类型

> [!Note]
> 在 Go 中，函数的命名跟Java略有不同，Java中函数名开头字母一般都是小写，Go中函数名开头字母一般都是大写（除main函数外）。

## Go 泛型

Go语言在1.18版本中增加了泛型功能。

限制

- 匿名结构体和匿名函数不支持泛型
- 不支持类型断言
- 支持泛型方法，只能通过recover来实现方法的泛型处理
- ～后的类型必须为基本类型，不能为接口类型

## Go 结构体声明

