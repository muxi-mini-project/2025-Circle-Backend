package controllers

import (
	"circle/request"
	"circle/service"

	"github.com/gin-gonic/gin"
)
type CircleControllers struct {
	us *service.CircleServices
}
func NewCircleControllers(us *service.CircleServices) *CircleControllers {
	return &CircleControllers{
		us: us,
	}
}
func (uc *CircleControllers) CreateCircle(c *gin.Context) {
    var get request.CreateCircle
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
    name,_:=c.Get("username")
	n,_:=name.(string)
    message:=uc.us.CreateCircle(n,get)
    c.JSON(200, gin.H{"message":message})
}
func (uc *CircleControllers) PendingCircle(c *gin.Context){
    name,_:=c.Get("username")
	n,_:=name.(string)
    circle,ok:=uc.us.PendingCircle(n)
    if !ok {
        c.JSON(400,gin.H{"error":"权限不足"})
    }
    c.JSON(200,gin.H{"circle":circle})
}
func (uc *CircleControllers) ApproveCircle(c *gin.Context){
    name,_:=c.Get("username")
	n,_:=name.(string)
    var get request.ApproveCircle
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
    message:=uc.us.ApproveCircle(n,get)
    c.JSON(200,gin.H{"message":message})
}
func (uc *CircleControllers) GetCircle(c *gin.Context){
    var get request.Circleid
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
    circle:=uc.us.GetCircle(get)
    c.JSON(200,gin.H{"circle":circle})
}
func (uc *CircleControllers) SelectCircle(c *gin.Context){
    circle:=uc.us.SelectCircle()
    c.JSON(200,gin.H{"circle":circle})
}
func (uc *CircleControllers) FollowCircle(c *gin.Context){
    var get request.Circleid
	if err := c.ShouldBindJSON(&get); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
    name,_:=c.Get("username")
	n,_:=name.(string)
    message:=uc.us.FollowCircle(n,get)
    c.JSON(200,gin.H{"message":message})
}