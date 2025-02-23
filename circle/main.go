package main

import (
	"circle/database"
	"circle/routes"
	"circle/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.InitDB()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,                                                         // 允许所有来源
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // 允许的方法
		AllowHeaders:    []string{"Authorization", "Content-Type"},                    // 允许的头部
		ExposeHeaders:   []string{"Content-Length"},                                   // 允许暴露的头部
		MaxAge:          12 * 60 * 60,                                                 // 预检请求的最大缓存时间
	}))
	r.Use(service.JwtMiddleware())
    routes.RunUser(database.DB,r)
    routes.RunPractice(database.DB,r)
	routes.RunTest(database.DB,r)
	routes.RunCircle(database.DB,r)
	routes.RunSearch(database.DB,r)
	r.Run("0.0.0.0:8080")
}