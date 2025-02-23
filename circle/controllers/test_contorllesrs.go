package controllers

import (
	"circle/request"
	"circle/service"

	"github.com/gin-gonic/gin"
)
type TestControllers struct {
	us *service.TestServices
}
func NewTestControllers(us *service.TestServices) *TestControllers {
	return &TestControllers{
		us: us,
	}
}
func (uc *TestControllers) Createtest(c *gin.Context) {
	var get request.Test
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
	id:= uc.us.CreateTest(n, get)
	if id==-1 {
		c.JSON(400, gin.H{"error": "创建测试失败"})
		return
	}
	c.JSON(200, gin.H{
		"id": id,
	    "message": "等待审核",
	})
}
func (uc *TestControllers) Createquestion(c *gin.Context) {
	var get request.TestQuestion
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	id:= uc.us.Createquestion(get)
	if id==-1 {
		c.JSON(400, gin.H{"error": "创建测试题目失败"})
		return
	}
	c.JSON(200, gin.H{
		"id": id,
		"success": "等待审核",
	})
}
func (uc *TestControllers) Createtestoption(c *gin.Context) {
	var get request.Option
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	id:= uc.us.Createtestoption(get)
	if id==-1 {
		c.JSON(400, gin.H{"error": "创建测试题目失败"})
		return
	}
	c.JSON(200, gin.H{
		"id": id,
		"success": "等待审核",
	})
}
func (uc *TestControllers) Gettest(c *gin.Context) {
	var get request.Gettest
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
	test:= uc.us.Gettest(n, get)
	c.JSON(200, gin.H{"test": test})
}
func (uc *TestControllers) Getquestion(c *gin.Context) {
	var get request.Gettest
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	question:= uc.us.Getquestion(get)
	c.JSON(200, gin.H{"question": question})
}
func (uc *TestControllers) Gettestoption(c *gin.Context) {
	var get request.GetPractice
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	option:= uc.us.Gettestoption(get)
	c.JSON(200, gin.H{"option": option})
}
func (uc *TestControllers) Getscore(c *gin.Context) {
	var get request.Score
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
	message:= uc.us.Getscore(n, get)
	c.JSON(200, gin.H{"message": message})
}
func (uc *TestControllers) Showtop(c *gin.Context) {
	var get request.Gettest
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	top:= uc.us.Showtop(get)
	c.JSON(200, gin.H{"top": top})
}
func (uc *TestControllers) Commenttest(c *gin.Context) {
	var get request.Commenttest
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
	message:= uc.us.Commenttest(n, get)
	c.JSON(200, gin.H{"message": message})
}
func (uc *TestControllers) GettestComment(c *gin.Context) {
	var get request.Gettest
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	comment:= uc.us.Gettestcomment(get)
	c.JSON(200, gin.H{"comment": comment})
}
func (uc *TestControllers) Lovetest(c *gin.Context) {
	var get request.Gettest
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	message:= uc.us.Lovetest(get)
	c.JSON(200, gin.H{"message": message})
}
func (uc *TestControllers) RecommentTest(c *gin.Context) {
	var get request.GetCircle
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	test:=uc.us.RecommentTest(get)
	c.JSON(200, gin.H{"test": test})
}
func (uc *TestControllers) HotTest(c *gin.Context) {
	var get request.GetCircle
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	test:=uc.us.HotTest(get)
	c.JSON(200, gin.H{"test": test})
}
func (uc *TestControllers) NewTest(c *gin.Context) {
	var get request.GetCircle
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	test:=uc.us.NewTest(get)
	c.JSON(200, gin.H{"test": test})
}
func (uc *TestControllers) FollowCircleTest(c *gin.Context) {
	name,_:=c.Get("username")
	n,_:=name.(string)
	test:=uc.us.FollowCircleTest(n)
	c.JSON(200, gin.H{"test": test})
}
