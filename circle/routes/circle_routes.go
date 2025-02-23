package routes
import (
	"circle/dao"
	"circle/service"
	"circle/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func RunCircle(db *gorm.DB,r *gin.Engine){
	ud := dao.NewCircleDao(db)
	us := service.NewCircleServices(ud)
	uc := controllers.NewCircleControllers(us)
    circle := r.Group("/circle")
	{
		circle.POST("/createcircle", uc.CreateCircle)
		circle.POST("/getcircle",  uc.GetCircle)
		circle.POST("/followcircle",  uc.FollowCircle)
		circle.GET("/selectcircle",  uc.SelectCircle)
		circle.GET("/pendingcircle",  uc.PendingCircle)
		circle.POST("/approvecircle",  uc.ApproveCircle)
	}
}