package images

import (
	"api/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) List(c *gin.Context) {
	images, err := h.imageService.List()
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	}
	response.Success(c, images)
}
