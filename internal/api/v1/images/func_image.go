package images

import (
	"api/internal/response"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func (h *handler) Image(c *gin.Context) {
	widthStr := c.Query("width")
	width, _ := strconv.Atoi(widthStr)
	heightStr := c.Query("height")
	height, _ := strconv.Atoi(heightStr)
	img, err := h.imageService.Image(width, height)
	// 如果图片的宽高 不是10的倍数直接删除
	defer func() {
		if (width != 0 && width%10 != 0) || (height != 0 && height%10 != 0) {
			os.Remove(img.Path)
		}
	}()
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	file, err := ioutil.ReadFile(img.Path)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Writer.WriteString(string(file))
}
