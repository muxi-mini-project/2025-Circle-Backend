package main

import (
	"circle/database"
	"circle/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)
func main(){
	r:=gin.Default()
	database.InitDB()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,                                                         // 允许所有来源
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // 允许的方法
		AllowHeaders:    []string{"Authorization", "Content-Type"},                    // 允许的头部
		ExposeHeaders:   []string{"Content-Length"},                                   // 允许暴露的头部
		MaxAge:          12 * 60 * 60,                                                 // 预检请求的最大缓存时间
	}))
	user := r.Group("/user")
    {
        user.POST("/register", controllers.Register)
        user.POST("/login", controllers.Login)
        user.GET("/logout", controllers.Logout)
        user.POST("/changepassword", controllers.Changepassowrd)
		user.POST("/changeusername", controllers.Changeusername)
        user.POST("/getcode", controllers.Getcode)
		user.POST("/checkcode",controllers.Checkcode)
		user.POST("/setphoto", controllers.Setphoto)
		user.POST("/setdiscription",controllers.Setdiscription)
		user.POST("/getname",controllers.Getname)
    }
	practice := r.Group("/practice")
	{
		practice.POST("/createpractice", controllers.Createpractice)
		practice.POST("/createoption", controllers.Createoption)
		practice.POST("/getpractice", controllers.Getpractice)
		practice.POST("/getoption", controllers.Getoption)
		practice.POST("/commentpractice", controllers.Commentpractice)
		practice.POST("/getcomment", controllers.GetComment)
		practice.POST("/checkanswer", controllers.Checkanswer)
		practice.POST("/selectpractice",controllers.Selectpractice)
	}
	test:=r.Group("/test")
	{
		test.POST("/createtest", controllers.Createtest)
		test.POST("/gettest", controllers.Gettest)
		test.POST("/getquestion", controllers.Getquestion)
		test.POST("/createquestion", controllers.Createquestion)
		test.POST("/createtestoption", controllers.Createtestoption)
		test.POST("/gettestoption", controllers.Gettestoption)
		test.POST("/commenttest", controllers.Commenttest)
		test.POST("/gettestcomment", controllers.GettestComment)
		test.POST("/showtop",controllers.Showtop)
		test.POST("/getscore",controllers.Getscore)
	}
	r.Run("0.0.0.0:8080")
}