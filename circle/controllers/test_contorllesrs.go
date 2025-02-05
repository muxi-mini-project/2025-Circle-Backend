package controllers

import (
	"circle/models"
	"circle/views"
	"circle/dao"
	"strconv"

	"github.com/gin-gonic/gin"
)
func Createtest(c *gin.Context) {
	token := c.GetHeader("Authorization")
	name := Username(token)
	discription := c.PostForm("discription")
	circle := c.PostForm("circle")
	testname:=c.PostForm("testname")
	id,_:=dao.GetIdByUser(name)
	test := models.Test{
		Userid:       id,
		Testname: testname,
		Discription: discription,
		Circle:     circle,
		Good:       0,
		Status:     "approved", // 待审核
	}
	id, err := dao.CreateTest(&test)
	if err != nil {
		views.Fail(c, "创建测试记录失败")
		return
	}
	views.Showid(c, id)
}
func Createquestion(c *gin.Context) {
	testid := c.PostForm("testid")
	p, _ := strconv.Atoi(testid)
	content := c.PostForm("content")
	difficulty := c.PostForm("difficulty")
	answer := c.PostForm("answer")
	variety := c.PostForm("variety")
	imageurl := c.PostForm("imageurl")
	explain:= c.PostForm("explain")
	question := models.TestQuestion{
		Content:   content,
		Testid:    p,
		Difficulty: difficulty,
		Answer:    answer,
		Variety:   variety,
		Imageurl:  imageurl,
		Explain:   explain,
	}
	id, err := dao.CreateQuestion(&question)
	if err != nil {
		views.Fail(c, "创建问题记录失败")
		return
	}
	views.Showid(c, id)
}
func Createtestoption(c *gin.Context) {
	practiceid := c.PostForm("practiceid")
	p, _ := strconv.Atoi(practiceid)
	content := c.PostForm("content")
	option := c.PostForm("option")
	options := models.TestOption{
		Content:   content,
		Practiceid: p,
		Option:    option,
	}
	id, err := dao.CreateTestOption(&options)
	if err != nil {
		views.Fail(c, "创建测试选项失败")
		return
	}
	views.Showid(c, id)
}
func Gettest(c *gin.Context) {
	testid := c.PostForm("testid")
	p, _ := strconv.Atoi(testid)
	test, err := dao.GetTestByID(p)
	if err != nil {
		views.Fail(c, "获取测试信息失败")
		return
	}
	token := c.GetHeader("Authorization")
	name := Username(token)
	id,_:=dao.GetIdByUser(name)
	err = dao.RecordTestHistory(p, id)
	if err != nil {
		views.Fail(c, "记录测试历史失败")
		return
	}
	views.Showtest(c, test)
}
func Getquestion(c *gin.Context) {
	testid := c.PostForm("testid")
	p, _ := strconv.Atoi(testid)
	questions, err := dao.GetQuestionsByTestID(p)
	if err != nil {
		views.Fail(c, "获取测试题目失败")
		return
	}
	views.Showtestquestion(c, questions)
}
func Gettestoption(c *gin.Context) {
	practiceid := c.PostForm("practiceid")
	p, _ := strconv.Atoi(practiceid)
	options, err := dao.GetTestOptionsByPracticeID(p)
	if err != nil {
		views.Fail(c, "获取测试选项失败")
		return
	}
	views.Showtestoption(c, options)
}
func Getscore(c *gin.Context) {
	token := c.GetHeader("Authorization")
	name := Username(token)
	user, err := dao.GetUserByName(name)
	if err != nil {
		views.Fail(c, "获取用户失败")
		return
	}
	testid := c.PostForm("testid")
	time := c.PostForm("time")  
	correctnumStr := c.PostForm("correctnum")
	correctnum, _ := strconv.Atoi(correctnumStr)
	testidInt, _ := strconv.Atoi(testid)
	t,_:=strconv.Atoi(time)
	top := models.Top{
		Userid:     user.Id,
		Correctnum: correctnum,
		Time:       t,
		Testid:     testidInt,
	}
	err = dao.SaveTopRecord(top)
	if err != nil {
		views.Fail(c, "保存成绩失败")
		return
	}
	views.Success(c, "成功")
}
func Showtop(c *gin.Context) {
	testid := c.PostForm("testid")
	tops, err := dao.GetTopByTestID(testid)
	if err != nil {
		views.Fail(c, "查询排行榜失败")
		return
	}
	views.Showtop(c, tops)
}
func Commenttest(c *gin.Context) {
	testid := c.PostForm("testid")
	p, _ := strconv.Atoi(testid)
	content := c.PostForm("content")
	token := c.GetHeader("Authorization")
	name := Username(token)
	user,_:=dao.GetUserByName(name)
	comment := models.TestComment{
		Content: content,
		Testid:  p,
		Userid:  user.Id,
	}
	if err := dao.CreateTestComment(&comment); err != nil {
		views.Fail(c, "评论失败")
		return
	}
	views.Success(c, "评论成功")
}
func GettestComment(c *gin.Context) {
	testid := c.PostForm("testid")
	p, _ := strconv.Atoi(testid)
	comments, err := dao.GetTestComments(p)
	if err != nil {
		views.Fail(c, "获取评论失败")
		return
	}
	views.Showtestcomment(c, comments)
}
func Lovetest(c *gin.Context) {
	testid := c.PostForm("testid")
	p, _ := strconv.Atoi(testid)
	test,_:= dao.GetTestByTestID(p)
	test.Good++
	_ = dao.UpdateTest(&test)
	views.Success(c, "点赞成功")
}
func RecommentTest(c *gin.Context) {
	circle:=c.PostForm("circle")
	test:=dao.RecommentTest(circle)
	views.ShowManytest(c, test)
}
func HotTest(c *gin.Context) {
	circle:=c.PostForm("circle")
	test:=dao.HotTest(circle)
	views.ShowManytest(c, test)
}
func NewTest(c *gin.Context) {
	circle:=c.PostForm("circle")
	test:=dao.NewTest(circle)
	views.ShowManytest(c, test)
}
func FollowCircleTest(c *gin.Context) {
	token:=c.GetHeader("Authorization")
	name:=Username(token)
	userid,_:=dao.GetIdByUser(name)
	test:=dao.FollowCircleTest(userid)
	views.ShowManytest(c, test)
}
