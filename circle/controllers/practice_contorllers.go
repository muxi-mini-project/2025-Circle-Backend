package controllers

import (
	"circle/models"
	"circle/views"
	"circle/dao"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
)
func Createpractice(c *gin.Context) {
	variety := c.PostForm("variety")
	difficulty := c.PostForm("difficulty")
	circle := c.PostForm("circle")
	imageurl := c.PostForm("imageurl")
	content := c.PostForm("content")
	answer := c.PostForm("answer")
	explain := c.PostForm("explain")
	token := c.GetHeader("Authorization")
	name := Username(token)
	id,_:=dao.GetIdByUser(name)
	practice := models.Practice{
		Userid:      id,
		Content:   content,
		Difficulty: difficulty,
		Circle:    circle,
		Answer:    answer,
		Variety:   variety,
		Imageurl:  imageurl,
		Explain: 	explain,
		Good:       0,
		Status:    "approved", // 待审核
	}
	err := dao.CreatePractice(&practice)
	if err != nil {
		views.Fail(c, "创建练习失败")
		return
	}
	views.Showid(c, practice.Practiceid)
}
func Createoption(c *gin.Context) {
	practiceid := c.PostForm("practiceid")
	p, _ := strconv.Atoi(practiceid)
	content := c.PostForm("content")
	option := c.PostForm("option")
	options := models.PracticeOption{
		Content:   content,
		Practiceid: p,
		Option:    option,
	}
	err := dao.CreatePracticeOption(&options)
	if err != nil {
		views.Fail(c, "创建选项失败")
		return
	}
	views.Success(c, "等待审核")
}
func Getpractice(c *gin.Context) {
	circle := c.PostForm("circle")
	practiceid:= c.PostForm("practiceid")
	var practice models.Practice
	if practiceid != "" {
		p, _ := strconv.Atoi(practiceid)
		practice= dao.GetPracticeByPracticeID(p)
	}else {
	    practice, _ = dao.GetPracticeByCircle(circle)
	}
	views.Showpractice(c, practice)
}
func Getoption(c *gin.Context) {
	practiceid := c.PostForm("practiceid")
	pid, err := strconv.Atoi(practiceid)
	if err != nil {
		views.Fail(c, "无效的 practiceid")
		return
	}
	options, err := dao.GetPracticeOptionsByPracticeID(pid)
	if err != nil {
		views.Fail(c, "获取 PracticeOption 失败")
		return
	}
	views.Showoptions(c, options)
}
func Commentpractice(c *gin.Context) {
	practiceid := c.PostForm("practiceid")
	p, _ := strconv.Atoi(practiceid) 
	content := c.PostForm("content")
	token := c.GetHeader("Authorization")
	name := Username(token)
	id,_:=dao.GetIdByUser(name)
	comment := models.PracticeComment{
		Content:    content,
		Practiceid: p,
		Userid:       id,
	}
	err := dao.CreatePracticeComment(&comment)
	if err != nil {
		views.Fail(c, "评论失败")
		return
	}
	views.Success(c, "评论成功")
}
func GetComment(c *gin.Context) {
	practiceid := c.PostForm("practiceid")
	p, _ := strconv.Atoi(practiceid)
	comments, err := dao.GetPracticeCommentsByPracticeID(p)
	if err != nil {
		views.Fail(c, "获取评论失败")
		return
	}
	views.Showcomment(c, comments)
}
func Checkanswer(c *gin.Context) {
	circle:=c.PostForm("circle")
	practiceid := c.PostForm("practiceid")
	answer := c.PostForm("answer")
	time:=c.PostForm("time")
	p, _ := strconv.Atoi(practiceid) 
	t,_:=strconv.Atoi(time)
	token := c.GetHeader("Authorization")
	name := Username(token)
	user, err := dao.GetUserByUsername(name)
	if err != nil {
		views.Fail(c, "获取用户信息失败")
		return
	}
	userpractice, err := dao.GetUserPracticeByUserID(user.Id,circle)
	if err != nil {
		views.Fail(c, "获取用户练习记录失败")
		return
	}
	userpractice.Alltime += t
	userpractice.Practicenum++
	if answer == "true" {
		userpractice.Correctnum++
	}
	err = dao.UpdateUserPractice(userpractice)
	if err != nil {
		views.Fail(c, "更新用户练习记录失败")
		return
	}
	practicehistory := models.Practicehistory{
		Practiceid: p,
		Userid:     user.Id,
		Answer:     answer,
	}
	err = dao.CreatePracticeHistory(&practicehistory)
	if err != nil {
		views.Fail(c, "创建练习历史记录失败")
		return
	}
	views.Success(c, "成功")
}
func Getrank(c *gin.Context) {
	circle:=c.PostForm("circle")
	name:=Username(c.GetHeader("Authorization"))
	id,_:=dao.GetIdByUser(name)
	rank:= dao.Showrank(id,circle)
	views.Success(c, fmt.Sprintf("%d", rank))
} 
func GetUserPractice(c *gin.Context) {
	circle:=c.PostForm("circle")
	name:=Username(c.GetHeader("Authorization"))
	id,_:=dao.GetIdByUser(name)
	userpractice, err := dao.GetUserPracticeByUserID(id,circle)
	if err != nil {
		views.Fail(c, "获取用户练习记录失败")
		return
	}
	views.Showuserpractice(c, *userpractice)
}
func Lovepractice(c *gin.Context) {
	practiceid := c.PostForm("practiceid")
	p, _ := strconv.Atoi(practiceid)
	practice := dao.GetPracticeByPracticeID(p)
	practice.Good++
	_ = dao.UpdatePractice(&practice)
	views.Success(c, "点赞成功")
}