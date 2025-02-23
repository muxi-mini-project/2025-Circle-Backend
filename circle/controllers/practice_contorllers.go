package controllers

import (
	"circle/request"
	"circle/service"

	"github.com/gin-gonic/gin"
)
type PracticeControllers struct {
	us *service.PracticeServices
}
func NewPracticeControllers(us *service.PracticeServices) *PracticeControllers {
	return &PracticeControllers{
		us: us,
	}
}
func (uc *PracticeControllers) Createpractice(c *gin.Context) {
	var practice request.Practice
	if err := c.ShouldBindJSON(&practice); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
	id:=uc.us.Createpractice(n,practice)
	if id==-1 {
		c.JSON(400, gin.H{"error": "创建失败"})
	}
	c.JSON(200, gin.H{
		"practiceid": id,
		"success":"等待审核",
	})
}
func (uc *PracticeControllers) Createoption(c *gin.Context) {
	var option request.Option
	if err := c.ShouldBindJSON(&option); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	message:=uc.us.Createoption(option)
	c.JSON(200, gin.H{"message":message})
}
func (uc *PracticeControllers) Getpractice(c *gin.Context) {
	var getpractice request.GetPractice
    if err := c.ShouldBindJSON(&getpractice); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
    pracctice:= uc.us.GetPractice(getpractice)
	c.JSON(200, gin.H{"practice": pracctice})
}
func (uc *PracticeControllers) Getoption(c *gin.Context) {
	var getpractice request.GetPractice
    if err := c.ShouldBindJSON(&getpractice); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	option:= uc.us.GetPracticeOption(getpractice)
	c.JSON(200, gin.H{"option": option})
}
func (uc *PracticeControllers) Commentpractice(c *gin.Context) {
	var get request.Comment
    if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
    message:=uc.us.CommentPractice(n, get)
	c.JSON(200, gin.H{"message":message})
}
func (uc *PracticeControllers) GetComment(c *gin.Context) {
	var getpractice request.GetPractice
    if err := c.ShouldBindJSON(&getpractice); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
    comment:= uc.us.GetComment(getpractice)
	c.JSON(200, gin.H{"comment": comment})
}
func (uc *PracticeControllers) Checkanswer(c *gin.Context) {
	var get request.CheckAnswer
    if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
	message:=uc.us.CheckAnswer(n, get)
	c.JSON(200, gin.H{"message":message})
}
func (uc *PracticeControllers) Getrank(c *gin.Context) {
	var get request.GetPractice
    if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
	message:=uc.us.Getrank(n, get)
    c.JSON(200, gin.H{"message":message})
} 
func (uc *PracticeControllers) GetUserPractice(c *gin.Context) {
	var get request.GetPractice
    if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
	practice:=uc.us.GetUserPractice(n, get)
	c.JSON(200, gin.H{"userpractice": practice})
}
func (uc *PracticeControllers) Lovepractice(c *gin.Context) {
	var get request.GetPractice
    if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	message:=uc.us.Lovepractice(get)
	c.JSON(200, gin.H{"message":message})
}