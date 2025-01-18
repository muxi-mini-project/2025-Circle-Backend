package main

import (
	"circle/circles"
	"circle/database"
	"circle/page"
	"circle/paper"
	"circle/post"
	"circle/search"
	"circle/user"
	"circle/photo"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title 圈子答题 API
// @version 1.0
// @description 这是一个用于圈子答题的API 文档
// @host 0.0.0.0:8080
// @BasePath /

func main() {
	// 创建一个新的 Gin 引擎
	r := gin.Default()
	database.InitDB()

	//跨域
	// CORS 中间件配置
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,                                                         // 允许所有来源
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // 允许的方法
		AllowHeaders:    []string{"Authorization", "Content-Type"},                    // 允许的头部
		ExposeHeaders:   []string{"Content-Length"},                                   // 允许暴露的头部
		MaxAge:          12 * 60 * 60,                                                 // 预检请求的最大缓存时间
	}))
	// 定义路由
	r.POST("/getcode", user.Getcode)
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)
	r.GET("/logout", user.Logout)
	r.POST("/change", user.Change)

	r.POST("/sendpost", post.Sendpost)
	r.POST("/randpost", post.Randpost)
	r.POST("/myfollowpost",post.Myfollowpost)
	r.POST("/readpost", post.Readpost)
	r.POST("/lovepost", post.Lovepost)
	r.POST("/collectpost", post.Collectpost)
	r.POST("/sharepost", post.Sharepost)
	r.POST("/followpost", post.Followpost)
	r.POST("/commentpost", post.Commentpost)
	r.POST("/readcomment", post.Readcomment)

	r.POST("/getquestion", paper.Getquestion)
	r.POST("/getscore", paper.Getscore)
	r.POST("/generatetest",paper.Generatetest)
	r.POST("/addquestionandoption",paper.Addquestionandoption)
	r.GET("/findtest",paper.Findtest)
	r.POST("/gettest",paper.Gettest)
	r.POST("/getownscore",paper.Getownscore)
	r.POST("/commenttest",paper.Commenttest)
	r.POST("/getcommenttest",paper.Getcommenttest)

	r.POST("/allcircles", circles.Allcircles)
	r.POST("/mycircles", circles.Mycircles)
	r.GET("/topcircle",circles.Topcircle)

	r.POST("/information", page.Information)
	r.POST("/myfollow", page.Myfollow)
	r.POST("/myfan", page.MyFan)
	r.GET("/mymessage", page.Mymessage)
	r.POST("/mypost", page.Mypost)
	r.POST("/mycollect", page.Mycollect)
	r.GET("/myhistory", page.Myhistory)
	r.POST("/mylevel",page.Mylevel)
	r.POST("/mytest",page.Mytest)
	r.POST("/myowntest",page.Myowntest)
	
	r.POST("searchpost",search.Searchpost)
	r.POST("searchcircle",search.Searchcircle)

	r.GET("/photo",photo.Photo)
	
	r.Run("0.0.0.0:8080")
}