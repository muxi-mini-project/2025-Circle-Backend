package controllers

import (
	"circle/database"
	"circle/models"
	"circle/views"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)
func timetosecond(t string) int {
	parts := strings.Split(t, ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	s, _ := strconv.Atoi(parts[2])
	return h*60*60+m*60+s
}
func Createtest(c *gin.Context){
	token := c.GetHeader("Authorization")
	name:=Username(token)
	discription:=c.PostForm("discription")
	circle:=c.PostForm("circle")
	test:=models.Test{
		Name:name,
		Discription:discription,
		Circle:circle,
		Good:0,
		Allcomment: 0,
		Status:"pending", //待审核
	}
	database.DB.Create(&test)
	id:=test.Testid
	views.Showid(c,id)
}
func Createquestion(c *gin.Context){
	testid:=c.PostForm("testid")
	p,_:=strconv.Atoi(testid)
	content:=c.PostForm("content")
	difficulty:=c.PostForm("difficulty")
	answer:=c.PostForm("answer")
	variety:=c.PostForm("variety")
	imageurl:=c.PostForm("imageurl")
	question:=models.TestQuestion{
		Content:content,
		Testid:p,
		Difficulty: difficulty,
		Answer: answer,
		Variety: variety,
		Imageurl: imageurl,
	}
	database.DB.Create(&question)
	id:=question.Questionid
	views.Showid(c,id)
}
func Createtestoption(c *gin.Context){
    practiceid:=c.PostForm("practiceid")
	p,_:=strconv.Atoi(practiceid)
	content:=c.PostForm("content")
	option:=c.PostForm("option")
	options:=models.TestOption{
		Content:content,
		Practiceid:p,
		Option:option,
	}
	database.DB.Create(&options)
	views.Success(c,"等待审核")
}
func Gettest(c *gin.Context){
	testid:=c.PostForm("testid")
	p,_:=strconv.Atoi(testid)
	var test models.Test
	database.DB.Where("testid = ?",p).First(&test)
	var user models.User
	database.DB.Where("name = ?",test.Name).First(&user)
	testhistory:=models.Testhistory{
		Testid:p,
		Userid:user.Id,
	}
	database.DB.Create(&testhistory)
	views.Showtest(c,test)
}
func Getquestion(c *gin.Context){
	testid:=c.PostForm("testid")
	p,_:=strconv.Atoi(testid)
    var practices []models.TestQuestion
	database.DB.Where("testid = ?",p).Find(&practices)
	views.Showtestquestion(c,practices)
}
func Gettestoption(c *gin.Context){
    practiceid:=c.PostForm("practiceid")
	p,_:=strconv.Atoi(practiceid)
	var options []models.TestOption
	database.DB.Where("practiceid = ?",p).Find(&options)
	views.Showtestoption(c,options)
}
func Getscore(c *gin.Context){
	token := c.GetHeader("Authorization")
	name:=Username(token)
	var user models.User
	database.DB.Where("name = ?",name).First(&user)
	testid:=c.PostForm("testid")
	time:=c.PostForm("time")  //h.m.s
	second:=timetosecond(time)
	n:=c.PostForm("correctnum")
	correctnum,_:=strconv.Atoi(n)	
	t,_:=strconv.Atoi(testid)
	feedback:=c.PostForm("feedback")
	var test models.Test
	database.DB.Where("testid = ?",testid).First(&test)
	test.Allcomment+=1
	if feedback=="good"{
		test.Good+=1
	}
	database.DB.Save(&test)
    top:=models.Top{
		Userid:user.Id,
		Correctnum:correctnum,
		Time:time,
		Second:second,
		Testid:t,
	}
	database.DB.Create(&top)
	views.Success(c,"成功")
}
func Showtop(c *gin.Context){
	testid:=c.PostForm("testid")
	var tops []models.Top
	database.DB.Order("correctnum desc , second asc").Where("testid=?",testid).Limit(10).Find(&tops)
	views.Showtop(c,tops)
}
func Commenttest(c *gin.Context){
	testid:=c.PostForm("testid")
	p,_:=strconv.Atoi(testid)
	content:=c.PostForm("content")
	token := c.GetHeader("Authorization")
	name:=Username(token)
	var user models.User
	database.DB.Where("name = ?",name).Find(&user)
	comment:=models.TestComment{
		Content:content,
		Testid:p,
		Userid:user.Id,
	}
	database.DB.Create(&comment)
	views.Success(c,"评论成功")
}
func GettestComment(c *gin.Context){
	testid:=c.PostForm("testid")
	p,_:=strconv.Atoi(testid)
	var comments []models.TestComment
	database.DB.Where("testid = ?",p).Find(&comments)
	views.Showtestcomment(c,comments)
}