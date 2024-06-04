// go.mod 文件是 go 用于依赖管理的文件
// 可以通过 go mod init 来创建该文件

// 模块名称
module go-demo

// go sdk version
go 1.22

// 第三方依赖
require (
	github.com/gin-gonic/gin v1.10.0
	github.com/goproxy/goproxy v0.16.9
)

// 排除以来继承里不需要的依赖
exclude (
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