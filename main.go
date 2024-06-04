package main

import _case "go-demo/demo_case"
import "github.com/gin-gonic/gin"
import "github.com/goproxy/goproxy"

func main() {
	gin.New()
	&goproxy.Goproxy{}
	_case.UserCase()
}
