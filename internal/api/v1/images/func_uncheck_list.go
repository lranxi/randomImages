package images

import (
	"api/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UncheckList 返回未检查图片列表
func (h *handler) UncheckList(c *gin.Context) {
	images, err := h.imageService.UncheckList()
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, images)

}
