package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 静态文件目录 路由
func setStaticRouter(s *Server) {
	// 显示未审核文件
	s.Engine.StaticFS("/static/images/uncheck", http.Dir("static/images/uncheck"))

	s.Engine.LoadHTMLGlob("templates/**")
	s.Engine.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "hello",
		})
	})
}
