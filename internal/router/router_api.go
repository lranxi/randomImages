package router

import (
	"api/configs"
	"api/internal/api/v1/images"
	imagesService "api/internal/service/images"
	"github.com/gin-contrib/pprof"
)

// api 路由
func setApiRouter(s *Server) {
	imageService := imagesService.New(s.Cache)
	imagesController := images.New(s.Logger, s.Cache, imageService)

	// debug环境开启pprof
	if configs.Get().Server.Env == "debug" {
		pprof.Register(s.Engine)
	}

	api := s.Engine

	api.POST("/upload/:category", imagesController.Upload)
	api.GET("/uncheck", imagesController.UncheckList)
	api.POST("/audit", imagesController.Audit)
	api.GET("/list", imagesController.List)
	api.GET("/image", imagesController.Image)
}
