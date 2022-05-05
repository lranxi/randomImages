package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 请求响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// Success 成功并返回数据
func Success(c *gin.Context, data interface{}) {
	newResponse(c, http.StatusOK, "success", data)
}

// Fail 处理失败
func Fail(c *gin.Context, code int, message string) {
	newResponse(c, code, message, "")
}
