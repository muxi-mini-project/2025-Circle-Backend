package controllers

import (
	"circle/dao"
	"circle/views"

	"github.com/gin-gonic/gin"
)
func SearchCircle(c *gin.Context) {
	token:=c.GetHeader("Authorization")
	name:=Username(token)
	userid,_:=dao.GetIdByUser(name)
    circlekey:=c.PostForm("circlekey")
    circle:=dao.SearchCircle(circlekey)
	dao.SearchHistory(circlekey,userid)
	views.ShowManyCircle(c,circle)
}
func SearchTest(c *gin.Context) {
	token:=c.GetHeader("Authorization")
	name:=Username(token)
	userid,_:=dao.GetIdByUser(name)
    testkey:=c.PostForm("testkey")
    test:=dao.SearchTest(testkey)
	dao.SearchHistory(testkey,userid)
	views.ShowManytest(c,test)
}
func SearchHistory(c *gin.Context) {
	token:=c.GetHeader("Authorization")
	name:=Username(token)
	userid,_:=dao.GetIdByUser(name)
	search:=dao.ShowSearchHistory(userid)
	views.ShowSearchHistory(c,search)
}
func DeleteHistory(c *gin.Context) {
	token:=c.GetHeader("Authorization")
	name:=Username(token)
	userid,_:=dao.GetIdByUser(name)
	dao.DeleteHistory(userid)
	views.Success(c,"删除成功")
}
func SearchPractice(c *gin.Context){
    circle:=c.PostForm("circle")
	practice:=dao.SelectPracticeByCircle(circle)
	views.ShowManyPractice(c,practice)
}