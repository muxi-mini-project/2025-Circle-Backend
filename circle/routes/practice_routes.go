package routes
import (
	"circle/dao"
	"circle/service"
	"circle/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func RunPractice(db *gorm.DB,r *gin.Engine){
	ud := dao.NewPracticeDao(db)
	us := service.NewPracticeServices(ud)
	uc := controllers.NewPracticeControllers(us)
    practice := r.Group("/practice")
	{
		practice.POST("/createpractice", uc.Createpractice)
		practice.POST("/createoption", uc.Createoption)
		practice.POST("/getpractice", uc.Getpractice)
		practice.POST("/getoption", uc.Getoption)
		practice.POST("/commentpractice", uc.Commentpractice)
		practice.POST("/getcomment", uc.GetComment)
		practice.POST("/checkanswer", uc.Checkanswer)
		practice.POST("/getrank", uc.Getrank)
		practice.POST("/getuserpractice", uc.GetUserPractice)
		practice.POST("/lovepractice", uc.Lovepractice)
	}
}