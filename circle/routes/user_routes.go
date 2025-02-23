package routes
import (
	"circle/dao"
	"circle/service"
	"circle/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func RunUser(db *gorm.DB,r *gin.Engine){
	ud := dao.NewUserDao(db)
	us := service.NewUserServices(ud)
	uc := controllers.NewUserControllers(us)
    user := r.Group("/user")
	{
		user.POST("/register", uc.Register)
		user.POST("/login", uc.Login)
		user.GET("/logout", uc.Logout)
		user.POST("/changepassword", uc.Changepassowrd)
		user.POST("/changeusername", uc.Changeusername)
		user.POST("/getcode", uc.Getcode)
		user.POST("/checkcode", uc.Checkcode)
		user.POST("/setphoto", uc.Setphoto)
		user.POST("/setdiscription", uc.Setdiscription)
		user.POST("/getname", uc.Getname)
		user.GET("/mytest", uc.Mytest)
		user.GET("/mypractice", uc.Mypractice)
		user.GET("/mydotest", uc.MyDoTest)
		user.GET("/mydopractice", uc.MyDoPractice)
		user.GET("/myuser", uc.MyUser)
	}
}