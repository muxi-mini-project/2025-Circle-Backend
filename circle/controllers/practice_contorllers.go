package controllers

import (
	"circle/database"
	"circle/models"
	"circle/views"
	"strconv"

	"github.com/gin-gonic/gin"
)
func Createpractice(c *gin.Context){
    variety:=c.PostForm("variety")
	difficulty:=c.PostForm("difficulty")
	circle:=c.PostForm("circle")
	imageurl:=c.PostForm("imageurl")
	content:=c.PostForm("content")
	answer:=c.PostForm("answer")
	name:=c.PostForm("name")
	if name==""{
	    token := c.GetHeader("Authorization")
	    name=Username(token)
	}
    pracitce:=models.Practice{
		Name:name,
		Content:content,
		Difficulty:difficulty,
		Circle:circle,
		Answer:answer,
		Variety:variety,
		Imageurl:imageurl,
		Status:"pending", //待审核
	}
	database.DB.Create(&pracitce)
	id:=pracitce.Practiceid
	views.Showid(c,id)
}
func Createoption(c *gin.Context){
	practiceid:=c.PostForm("practiceid")
	p,_:=strconv.Atoi(practiceid)
	content:=c.PostForm("content")
	option:=c.PostForm("option")
	options:=models.PracticeOption{
		Content:content,
		Practiceid:p,
		Option:option,
	}
	database.DB.Create(&options)
	views.Success(c,"等待审核")
}
func Getpractice(c *gin.Context){
	circle:=c.PostForm("circle")
    var practices models.Practice
	database.DB.Where("status = ?", "approved").Where("circle = ?",circle).Order("RAND()").First(&practices)
	views.Showpractice(c,practices)
}
func Getoption(c *gin.Context){
    practiceid:=c.PostForm("practiceid")
	var options []models.PracticeOption
	database.DB.Where("practiceid = ?",practiceid).Find(&options)
	views.Showoptions(c,options)
}
func Commentpractice(c *gin.Context){
	practiceid:=c.PostForm("practiceid")
	p,_:=strconv.Atoi(practiceid)
	content:=c.PostForm("content")
	token := c.GetHeader("Authorization")
	name:=Username(token)
	comment:=models.PracticeComment{
		Content:content,
		Practiceid:p,
		Name:name,
	}
	database.DB.Create(&comment)
	views.Success(c,"评论成功")
}
func GetComment(c *gin.Context){
	practiceid:=c.PostForm("practiceid")
	p,_:=strconv.Atoi(practiceid)
	var comments []models.PracticeComment
	database.DB.Where("practiceid = ?",p).Find(&comments)
	views.Showcomment(c,comments)
}
func Checkanswer(c *gin.Context){
	practiceid:=c.PostForm("practiceid")
	answer:=c.PostForm("answer")
	p,_:=strconv.Atoi(practiceid)
	token := c.GetHeader("Authorization")
	name:=Username(token)
	var user models.User
	database.DB.Where("name = ?",name).Find(&user)
	var userpractice models.UserPractice
	database.DB.Where("userid = ?",user.Id).First(&userpractice)
	userpractice.Practicenum=userpractice.Practicenum+1
	if answer=="true" {
		userpractice.Correctnum=userpractice.Correctnum+1
	}
	database.DB.Save(&userpractice)
	practicehistory := models.Practicehistory{
		Practiceid:p,
		Userid:user.Id,
		Answer:answer,
	}
	database.DB.Create(&practicehistory)
	views.Success(c,"回答成功")
}
func Selectpractice(c *gin.Context){
	circle:=c.PostForm("circle")
	var practices []models.Practice
	database.DB.Where("status = ?", "approved").Where("circle = ?",circle).Limit(5).Find(&practices)
	views.Selectpractice(c,practices)
}