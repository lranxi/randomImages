package images

import (
	"api/internal/model"
	"api/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) Audit(c *gin.Context) {
	var params []model.Image
	c.ShouldBind(&params)
	if params == nil {
		response.Fail(c, http.StatusBadRequest, "param invalid")
	}

	h.imageService.Audit(params)

	response.Success(c, "")
}
