package images

import (
	"api/configs"
	"api/internal/response"
	"api/pkg/file"
	"api/pkg/random"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *handler) Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["images"]
	category, err := strconv.Atoi(c.Param("category"))
	if err != nil {
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	for _, f := range files {
		randomStr := random.Str(configs.FileNameLength)
		dst := configs.UncheckStaticFileDir + randomStr + file.Filetype(f.Filename)
		c.SaveUploadedFile(f, dst)

		// 开启一个goroutine对上传文件进行检查
		err := h.imageService.UploadPreCheck(dst, category)
		if err != nil {
			response.Fail(c, http.StatusBadRequest, err.Error())
			return
		}
	}
	response.Success(c, "")
}
