package main

import (
	"circle/database" 
	"circle/user"     
	"circle/post"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	// 创建一个新的 Gin 引擎
	r := gin.Default()
    database.InitDB()
	//跨域
	// CORS 中间件配置
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // 允许所有来源
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // 允许的方法
		AllowHeaders:     []string{"Authorization", "Content-Type"}, // 允许的头部
		ExposeHeaders:    []string{"Content-Length"}, // 允许暴露的头部
		MaxAge: 12 * 60 * 60, // 预检请求的最大缓存时间
	}))
	// 定义路由
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)
	r.GET("/logout", user.Logout)
	r.POST("/change", user.Change)
    r.POST("/sendpost",post.Sendpost)
	r.POST("/readpost",post.Readpost)
	r.POST("/lovepost",post.Lovepost)
	r.POST("/collectpost",post.Collectpost)
	r.POST("/sharepost",post.Sharepost)
	r.POST("/followpost",post.Followpost)
	r.POST("/commentpost",post.Commentpost)
	r.POST("/readcomment",post.Readcomment)
	r.Run("0.0.0.0:8080")
}
