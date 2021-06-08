package main

import (
	"github.com/gin-gonic/gin"
	"platform_clothes/handler"
)

var (
	e *gin.Engine
)

func main() {
	e = gin.Default()
	initHandler()
}

func initHandler() {
	handler.Init()
	handler.RegisterHandlers(e)
}

//s
