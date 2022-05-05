package router

import "net/http"

// 静态文件目录 路由
func setStaticRouter(s *Server) {
	s.Engine.StaticFS("/static/images/uncheck", http.Dir("static/images/uncheck"))
	//s.Engine.Static("/uncheck", "static/images")
}
