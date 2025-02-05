package main

import (
	"circle/controllers"
	"circle/database"

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
	user := r.Group("/user")
	{
		user.POST("/register", controllers.Register)
		user.POST("/login", controllers.Login)
		user.GET("/logout", controllers.Logout)
		user.POST("/changepassword", controllers.Changepassowrd)
		user.POST("/changeusername", controllers.Changeusername)
		user.POST("/getcode", controllers.Getcode)
		user.POST("/checkcode", controllers.Checkcode)
		user.POST("/setphoto", controllers.Setphoto)
		user.POST("/setdiscription", controllers.Setdiscription)
		user.POST("/getname", controllers.Getname)
		user.GET("/mytest", controllers.Mytest)
		user.GET("/mypractice", controllers.Mypractice)
		user.GET("/mydotest", controllers.MyDoTest)
		user.GET("/mydopractice", controllers.MyDoPractice)
		user.GET("/myuser", controllers.MyUser)
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
		practice.POST("/getrank", controllers.Getrank)
		practice.POST("/getuserpractice", controllers.GetUserPractice)
		practice.POST("/lovepractice", controllers.Lovepractice)
	}
	test := r.Group("/test")
	{
		test.POST("/createtest", controllers.Createtest)
		test.POST("/gettest", controllers.Gettest)
		test.POST("/getquestion", controllers.Getquestion)
		test.POST("/createquestion", controllers.Createquestion)
		test.POST("/createtestoption", controllers.Createtestoption)
		test.POST("/gettestoption", controllers.Gettestoption)
		test.POST("/commenttest", controllers.Commenttest)
		test.POST("/gettestcomment", controllers.GettestComment)
		test.POST("/showtop", controllers.Showtop)
		test.POST("/getscore", controllers.Getscore)
		test.POST("/lovetest", controllers.Lovetest)
		test.POST("/recommenttest", controllers.RecommentTest)
		test.POST("/hottest", controllers.HotTest)
		test.POST("/newtest", controllers.NewTest)
		test.GET("/followcircletest", controllers.FollowCircleTest)
	}
	circle := r.Group("/circle")
	{
		circle.POST("/createcircle", controllers.CreateCircle)
		circle.POST("/getcircle", controllers.GetCircle)
		circle.POST("/followcircle", controllers.FollowCircle)
		circle.GET("/selectcircle", controllers.SelectCircle)
		circle.GET("/pendingcircle", controllers.PendingCircle)
		circle.POST("/approvecircle", controllers.ApproveCircle)
	}
	search:=r.Group("/search")
	{
        search.POST("/searchcircle",controllers.SearchCircle)
		search.POST("/searchtest",controllers.SearchTest)
		search.GET("/searchhistory",controllers.SearchHistory)
		search.GET("/deletehistory",controllers.DeleteHistory)
		search.POST("/searchpractice",controllers.SearchPractice)
	}
	r.Run("0.0.0.0:8080")
}
