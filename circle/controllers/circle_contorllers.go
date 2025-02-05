package controllers

import (
	"circle/dao"
	"circle/models"
	"circle/views"
	"strconv"

	"github.com/gin-gonic/gin"
)
func CreateCircle(c *gin.Context) {
    name:=c.PostForm("name")
    discription:=c.PostForm("discription")
    Imageurl:=c.PostForm("imageurl")
    usename:=Username(c.GetHeader("Authorization"))
    userid,_:=dao.GetIdByUser(usename)
    circle:=models.Circle{
        Name:name,
        Discription:discription,
        Imageurl:Imageurl,
        Userid:userid,
        Status:"pending",
    }
    _=dao.CreateCircle(&circle)
    views.Success(c,"等待审核")
}
func PendingCircle(c *gin.Context){
    token:=c.GetHeader("Authorization")
    name:=Username(token)
    if name!="root" {
        views.Fail(c,"权限不足")
        return
    }
    circle,_:=dao.SelectPendingCircle()
    views.ShowCircle(c,circle)
}
func ApproveCircle(c *gin.Context){
    token:=c.GetHeader("Authorization")
    name:=Username(token)
    if name!="root" {
        views.Fail(c,"权限不足")
        return
    }
    circleid:=c.PostForm("circleid")
    decide:=c.PostForm("decide")
    id,_:=strconv.Atoi(circleid)
    circle,_:=dao.GetCircleByID(id)
    if decide=="false" {
        _=dao.DeleteCircleByID(id)
    }else {
        circle.Status="approved"
        _=dao.UpdateCircle(&circle,id)
    }
    views.Success(c,"审核结束")
}
func GetCircle(c *gin.Context){
    circleid:=c.PostForm("circleid")
    id,_:=strconv.Atoi(circleid)
    circle,_:=dao.GetCircleByID(id)
    views.ShowCircle(c,circle)
}
func SelectCircle(c *gin.Context){
    circle,_:=dao.SelectCircle()
    views.ShowManyCircle(c,circle)
}
func FollowCircle(c *gin.Context){
    circleid:=c.PostForm("circleid")
    cir,_:=strconv.Atoi(circleid)
    token:=c.GetHeader("Authorization")
    name:=Username(token)
    id,_:=dao.GetIdByUser(name)
    _=dao.FollowCircle(cir,id)
    views.Success(c,"关注成功")
}