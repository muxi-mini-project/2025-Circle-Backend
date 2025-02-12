package controllers

import (
	"circle/service"
	"circle/request"

	"github.com/gin-gonic/gin"
)
type SearchControllers struct {
	us *service.SearchServices
}
func NewSearchControllers(us *service.SearchServices) *SearchControllers {
	return &SearchControllers{
		us: us,
	}
}
func (uc *SearchControllers) SearchCircle(c *gin.Context) {
	name,_:=c.Get("username")
	n,_:=name.(string)
	var circlekey request.Circlekey
	if err:=c.ShouldBindJSON(&circlekey);err!=nil{
		c.JSON(400,gin.H{"error":"错误的参数"})
		return
	}
	circle:=uc.us.SearchCircle(n,circlekey.Circlekey)
	c.JSON(200,gin.H{"circle":circle})
}
func (uc *SearchControllers) SearchTest(c *gin.Context) {
	name,_:=c.Get("username")
	n,_:=name.(string)
	var testkey request.Testkey
	if err:=c.ShouldBindJSON(&testkey);err!=nil{
		c.JSON(400,gin.H{"error":"错误的参数"})
		return
	}
	test:=uc.us.SearchTest(n,testkey.Testkey)
	c.JSON(200,gin.H{"test":test})
}
func (uc *SearchControllers) SearchHistory(c *gin.Context) {
	name,_:=c.Get("username")
	n,_:=name.(string)
	history:=uc.us.SearchHistory(n)
	c.JSON(200,gin.H{"history":history})
}
func (uc *SearchControllers) DeleteHistory(c *gin.Context) {
	name,_:=c.Get("username")
	n,_:=name.(string)
	uc.us.DeleteHistory(n)
	c.JSON(200,gin.H{"message":"删除成功"})
}
func (uc *SearchControllers) SearchPractice(c *gin.Context){
    var circle request.Circle
	if err:=c.ShouldBindJSON(&circle);err!=nil{
		c.JSON(400,gin.H{"error":"错误的参数"})
		return
	}
	practice:=uc.us.SearchPractice(circle.Circle)
	c.JSON(200,gin.H{"message":practice})
}