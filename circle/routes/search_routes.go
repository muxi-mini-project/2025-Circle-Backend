package routes
import (
	"circle/dao"
	"circle/service"
	"circle/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func RunSearch(db *gorm.DB,r *gin.Engine){
	ud := dao.NewSearchDao(db)
	us := service.NewSearchServices(ud)
	uc := controllers.NewSearchControllers(us)
    search:=r.Group("/search")
	{
        search.POST("/searchcircle",uc.SearchCircle)
		search.POST("/searchtest",uc.SearchTest)
		search.GET("/searchhistory",uc.SearchHistory)
		search.GET("/deletehistory",uc.DeleteHistory)
		search.POST("/searchpractice",uc.SearchPractice)
	}
}