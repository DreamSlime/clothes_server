package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	_const "platform_clothes/const"
	"platform_clothes/httpclient"
)

//对账系统handler
type SimilarHandler struct {
}

func SimilarHandlerInstance() *SimilarHandler {
	return &SimilarHandler{}
}

// Register 注册handler
func (h *SimilarHandler) Register(e *gin.Engine) {
	e.POST("/similar/:captchaId", h.Similar)
}

func (h *SimilarHandler) Similar(c *gin.Context) {
	filenames, _ := GetAllFile(_const.IMAGE_PATH + _const.SimilarClothes)
	imageName := c.Param("captchaId")
	httpclient.HttpClientInstance().Get("http://127.0.0.1:8888/" + imageName)
	c.JSON(http.StatusOK, gin.H{
		"filenames": filenames,
	})
}

func GetAllFile(pathname string) ([]string, error) {
	var s []string
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := "http://127.0.0.1:6789/similar/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}
