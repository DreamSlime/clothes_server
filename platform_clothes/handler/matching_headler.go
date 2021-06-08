package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//对账系统handler
type MatchingHandler struct {
}

func MatchingHandlerInstance() *MatchingHandler {
	return &MatchingHandler{}
}

// Register 注册handler
func (h *MatchingHandler) Register(e *gin.Engine) {
	e.POST("/upload", h.UploadImage)
	e.GET("/show", h.Show)
}

func (h *MatchingHandler) UploadImage(c *gin.Context) {
	// 单个文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println(file.Filename)
	dst := fmt.Sprintf("./DownloadFile/%s", file.Filename)
	// 上传文件到指定的目录
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"name":   "xxx.png",
		"status": "done",
		"url":    "http://127.0.0.1:6789/DownloadFile/" + file.Filename,
		"uid":    "http://127.0.0.1:6789/DownloadFile/" + file.Filename,
	})
}

func (h *MatchingHandler) Show(c *gin.Context) {
	// 单个文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println(file.Filename)
	dst := fmt.Sprintf("./DownloadFile/%s", file.Filename)
	// 上传文件到指定的目录
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}
