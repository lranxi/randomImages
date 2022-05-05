package router

import (
	"api/configs"
	cache "api/internal/repository/cache"
	"api/internal/router/interceptor/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	Engine *gin.Engine
	Logger *zap.Logger
	Cache  cache.Cache
}

func New(accessLogger *zap.Logger) (*Server, error) {
	cfg := configs.Get().Server
	// 设置启动模式
	gin.SetMode(cfg.Env)
	engine := gin.New()
	engine.Use(
		logger.LogMiddleware(accessLogger),
		gin.Recovery(),
	)

	// 创建redis
	redis, err := cache.NewCache()
	if err != nil {
		return nil, err
	}

	server := &Server{
		Logger: accessLogger,
		Cache:  redis,
		Engine: engine,
	}

	// 添加api路由
	setApiRouter(server)

	// 添加静态文件路由
	setStaticRouter(server)

	return server, nil

}
