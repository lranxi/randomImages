package images

import (
	"api/internal/repository/cache"
	"api/internal/service/images"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type handler struct {
	logger       *zap.Logger
	cache        cache.Cache
	imageService images.ImageService
}

type Handler interface {
	// Upload 上传文件.多文件上传
	Upload(c *gin.Context)
	UncheckList(c *gin.Context)
	Audit(c *gin.Context)
	List(c *gin.Context)
	// Image 返回一张审核通过的随机图片
	Image(c *gin.Context)
}

func New(logger *zap.Logger, cache cache.Cache, service images.ImageService) Handler {
	return &handler{
		logger:       logger,
		cache:        cache,
		imageService: service,
	}
}
