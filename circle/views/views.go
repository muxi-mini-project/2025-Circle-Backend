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
func Selectpractice(c *gin.Context,pracitce []models.Practice){
	c.JSON(200,gin.H{"practices":pracitce,})
}


func Showtest(c *gin.Context,test models.Test){
	c.JSON(200,gin.H{"test":test})
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