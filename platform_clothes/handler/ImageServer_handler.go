package handler

import "github.com/gin-gonic/gin"

//对账系统handler
type ImageHandler struct {
}

func ImageHandlerInstance() *ImageHandler {
	return &ImageHandler{}
}

// Register 注册handler
func (h *ImageHandler) Register(e *gin.Engine) {
	e.GET("/DownloadFile/:captchaId", h.DownloadFile)
	e.GET("/similar/:captchaId", h.Similar)
}

func (h *ImageHandler) DownloadFile(c *gin.Context) {
	imageName := c.Param("captchaId")
	c.File("./DownloadFile/" + imageName)
}

func (h *ImageHandler) Similar(c *gin.Context) {
	imageName := c.Param("captchaId")
	c.File("./SimilarClothes/" + imageName)
}
