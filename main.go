package main

import (
	"api/configs"
	"api/internal/router"
	"api/pkg/logger"
	"api/pkg/shutdown"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	config := configs.Get().Server
	fmt.Println(config)
	accessLogger, err := logger.NewJSONLogger(
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, configs.Get().Server.Env)),
		logger.WithFileP(configs.ProjectAccessLogFile),
		logger.WithDebugLevel(),
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = accessLogger.Sync()
	}()

	e, err := router.New(accessLogger)
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler: e.Engine,
		Addr:    fmt.Sprintf("%s:%d", configs.Get().Server.Addr, configs.Get().Server.Port),
	}
	// 监听端口
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			accessLogger.Fatal("http server startup error", zap.Error(err))
		}
		accessLogger.Info("dddd")

	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				accessLogger.Error("server shutdown error", zap.Error(err))
			}
		},

		// 关闭缓存
		func() {
			if e.Cache != nil {
				if err := e.Cache.Close(); err != nil {
					accessLogger.Error("cache close error", zap.Error(err))
				}
			}
		},
	)
}
