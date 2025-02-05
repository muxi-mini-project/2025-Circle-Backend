package views
import (
	"github.com/gin-gonic/gin"
	"circle/models"
)
func Success(c *gin.Context,message string){
	c.JSON(200,gin.H{"success":message})
}
func Fail(c *gin.Context,message string){
	c.JSON(400,gin.H{"fail":message})
}
func ShowUser(c *gin.Context,user models.User){
	c.JSON(200,gin.H{"user":user})
}


func Showpractice(c *gin.Context,practices models.Practice){
	c.JSON(200,gin.H{"practices":practices})
}
func Showoptions(c *gin.Context,options []models.PracticeOption){
	c.JSON(200,gin.H{"options":options})
}
func Showcomment(c *gin.Context,comments []models.PracticeComment){
	c.JSON(200,gin.H{"comments":comments})
}
func Showid(c *gin.Context,id int){
	c.JSON(200,gin.H{
		"id":id,
		"success":"等待审核",
	})
}
func ShowManyPractice(c *gin.Context,pracitce []models.Practice){
	c.JSON(200,gin.H{"practices":pracitce,})
}
func Showuserpractice(c *gin.Context,userpractices models.UserPractice){
	c.JSON(200,gin.H{"userpractices":userpractices})
}
func ShowManyHistoryPractice(c *gin.Context,historypractices []models.Practicehistory){
	c.JSON(200,gin.H{"historypractices":historypractices})
}


func Showtest(c *gin.Context,test models.Test){
	c.JSON(200,gin.H{"test":test})
}
func ShowManytest(c *gin.Context,tests []models.Test){
	c.JSON(200,gin.H{"tests":tests})
}
func Showtestquestion(c *gin.Context,questions []models.TestQuestion){
	c.JSON(200,gin.H{"questions":questions})
}
func Showtestoption(c *gin.Context,options []models.TestOption){
	c.JSON(200,gin.H{"options":options})
}
func Showtop(c *gin.Context,tops []models.Top){
	c.JSON(200,gin.H{"tops":tops})
}
func Showtestcomment(c *gin.Context,comments []models.TestComment){
	c.JSON(200,gin.H{"comments":comments})
}
func ShowManyTestid(c *gin.Context,testids []models.Testhistory){
	c.JSON(200,gin.H{"testids":testids})
}


func ShowCircle(c *gin.Context,circles models.Circle){
	c.JSON(200,gin.H{"circles":circles})
}
func ShowManyCircle(c *gin.Context,circles []models.Circle){
	c.JSON(200,gin.H{"circles":circles})
}


func ShowSearchHistory(c *gin.Context,search []models.SearchHistory){
	c.JSON(200,gin.H{"search":search})
}