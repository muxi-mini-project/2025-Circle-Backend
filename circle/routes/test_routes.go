package routes
import (
	"circle/dao"
	"circle/service"
	"circle/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func RunTest(db *gorm.DB,r *gin.Engine){
	ud := dao.NewTestDao(db)
	us := service.NewTestServices(ud)
	uc := controllers.NewTestControllers(us)
    test := r.Group("/test")
	{
		test.POST("/createtest", uc.Createtest)
		test.POST("/gettest", uc.Gettest)
		test.POST("/getquestion", uc.Getquestion)
		test.POST("/createquestion", uc.Createquestion)
		test.POST("/createtestoption", uc.Createtestoption)
		test.POST("/gettestoption", uc.Gettestoption)
		test.POST("/commenttest", uc.Commenttest)
		test.POST("/gettestcomment", uc.GettestComment)
		test.POST("/showtop", uc.Showtop)
		test.POST("/getscore", uc.Getscore)
		test.POST("/lovetest", uc.Lovetest)
		test.POST("/recommenttest", uc.RecommentTest)
		test.POST("/hottest", uc.HotTest)
		test.POST("/newtest", uc.NewTest)
		test.GET("/followcircletest", uc.FollowCircleTest)
	}
}