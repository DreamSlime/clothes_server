package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	port = ":6789"
)

// Handler handler接口，提供一个注册函数
type Handler interface {
	Register(e *gin.Engine)
}

var (
	handlers []Handler
)

func Init() {
	// 将新的handler实例化，并且添加到slice中
	handlers = make([]Handler, 0)
	handlers = append(handlers, ImageHandlerInstance())
	handlers = append(handlers, MatchingHandlerInstance())
	handlers = append(handlers, SimilarHandlerInstance())
}

// RegisterHandlers 注册所有的api路径，包括函数注册和Handler接口注册
func RegisterHandlers(e *gin.Engine) {

	for _, h := range handlers {
		h.Register(e)
	}

	e.Run(port)
}
