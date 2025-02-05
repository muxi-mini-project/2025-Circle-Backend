package controllers

import (
	"circle/dao"
	"circle/models"
	"circle/views"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/smtp"
	"strings"
	"sync"
	"time"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)
var lock sync.Mutex
var m=make(map[string]string)
var WhitelistedTokens=make(map[string]int)
func Token(username string) string {
    randomBytes := make([]byte, 16)
    rand.Read(randomBytes)
    token := base64.URLEncoding.EncodeToString(randomBytes) + "|" + username
    return token
}
func Username(token string) string {
	name := strings.Split(token, "|")
	return name[1]
}
func Getemail(ee string,VerificationCode string)  {
	data, _ := ioutil.ReadFile("data.txt")
	m:=string(data)
	html := "<h1>验证码：" + VerificationCode + "</h1>"
	e := email.NewEmail()
	e.From = "luohuixi <2388287244@qq.com>"    
	e.To = []string{ee}         
	e.Subject = "验证码"                            
	e.Text = []byte("This is a plain text body.") 
	e.HTML = []byte(html)                    
	smtpHost := "smtp.qq.com"                                           
	smtpPort := "587"                                                             
	auth := smtp.PlainAuth("", "2388287244@qq.com", m, smtpHost)
	e.Send(smtpHost+":"+smtpPort, auth)  
}
func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano()) 
	code := rand.Intn(9000) + 1000   
	return fmt.Sprintf("%04d", code)
}
func Getcode(c *gin.Context) {
    email := c.PostForm("email")
    code := generateVerificationCode()
    Getemail(email, code)
    lock.Lock()
    defer lock.Unlock()
    m[email] = code
    views.Success(c, "验证码已发送")
}
func Checkcode(c *gin.Context){
	email:=c.PostForm("email")
	code:=c.PostForm("code")
	if m[email]!=code{
		views.Fail(c,"验证码错误")
		return
	}
	views.Success(c,"验证成功")
}
func Register(c *gin.Context) {
    email := c.PostForm("email")
    password := c.PostForm("password")
    count, err := dao.CountUsersByEmail(email)
    if err != nil {
        views.Fail(c, "查询数据库失败")
        return
    }
    if count > 0 {
        views.Fail(c, "该邮箱已注册")
        return
    }
    totalUsers, err := dao.CountUsersByName("")
    if err != nil {
        views.Fail(c, "查询数据库失败")
        return
    }
    name := "Circle_" + fmt.Sprintf("%04d", totalUsers+1)
    user := models.User{
        Email:    email,
        Password: password,
        Name:     name,
		Discription: "这里空空如也",
    }
    if err := dao.CreateUser(&user); err != nil {
        views.Fail(c, "创建用户失败")
        return
    }
    views.Success(c, "注册成功")
}
func Login(c *gin.Context) {
    email := c.PostForm("email")
    password := c.PostForm("password")
    user, err := dao.GetUserByEmail(email)
    if err != nil {
        views.Fail(c, "该邮箱未注册")
        return
    }
    if user.Password != password {
        views.Fail(c, "密码错误")
        return
    }
    token := Token(user.Name)
    lock.Lock()
    WhitelistedTokens[token] = 1
    lock.Unlock()
    views.Success(c, token)
}
func Logout(c *gin.Context){
	lock.Lock()
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	if _,ok:=WhitelistedTokens[token];!ok {
		views.Fail(c,"token无效")
		return
	}
	delete(WhitelistedTokens,token)
	views.Success(c,"登出成功")
}
func Changepassowrd(c *gin.Context) {
    newpassword:=c.PostForm("newpassword")
	token := c.GetHeader("Authorization")
	if _,ok:=WhitelistedTokens[token]; !ok {
		views.Fail(c,"token无效")
		return
	}
	user, _ := dao.GetUserByName(Username(token))
	user.Password=newpassword
	_=dao.UpdateUser(user)
	views.Success(c,"密码修改成功")
}
func Changeusername(c *gin.Context){
	lock.Lock()
	defer lock.Unlock()
	newusername:=c.PostForm("newusername")
	token := c.GetHeader("Authorization")
	if _,ok:=WhitelistedTokens[token]; !ok {
		views.Fail(c,"token无效")
		return
	}
	count,_:=dao.CountUsersByName(newusername)
	if count>0{
		views.Fail(c,"该用户名已存在")
		return
	}
	name := Username(token)
	user, err := dao.GetUserByName(name)
	if err != nil {
		views.Fail(c, "用户查询失败")
		return
	}
	user.Name = newusername
	err = dao.UpdateUser(user)
	if err != nil {
		views.Fail(c, "用户名更新失败")
		return
	}
	newtoken:=Token(user.Name)
	WhitelistedTokens[newtoken]=1
	delete(WhitelistedTokens,token)
	views.Success(c,newtoken)
}
func Setphoto(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	if _, ok := WhitelistedTokens[token]; !ok {
		views.Fail(c, "token无效")
		return
	}
	name := Username(token)
	user, err := dao.GetUserByName(name)
	if err != nil {
		views.Fail(c, "用户查询失败")
		return
	}
	user.Imageurl = c.PostForm("imageurl")
	err = dao.UpdateUser(user)
	if err != nil {
		views.Fail(c, "头像更新失败")
		return
	}
	views.Success(c, "头像添加成功")
}
func Setdiscription(c *gin.Context) {
	discription := c.PostForm("discription")
	token := c.GetHeader("Authorization")
	name := Username(token)
	user, err := dao.GetUserByName(name)
	if err != nil {
		views.Fail(c, "用户查询失败")
		return
	}
	user.Discription = discription
	err = dao.UpdateUser(user)
	if err != nil {
		views.Fail(c, "简介更新失败")
		return
	}

	views.Success(c, "简介修改成功")
}
func Getname(c *gin.Context) {
	id := c.PostForm("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		views.Fail(c, "无效的用户ID")
		return
	}
	user, err := dao.GetUserByID(userID)
	if err != nil {
		views.Fail(c, "用户查询失败")
		return
	}
	views.Success(c, user.Name)
}
func Mytest(c *gin.Context) {
	token := c.GetHeader("Authorization")
	name := Username(token)
	userid, _ := dao.GetIdByUser(name)
	test,_:=dao.GetTestByUserid(userid)
	views.ShowManytest(c,test)
}
func Mypractice(c *gin.Context) {
	token := c.GetHeader("Authorization")
	name := Username(token)
	userid, _ := dao.GetIdByUser(name)
	practice,_:=dao.GetPracticeByUserid(userid)
	views.ShowManyPractice(c,practice)
}
func MyDoTest(c *gin.Context) {
	token := c.GetHeader("Authorization")
	name := Username(token)
	userid, _ := dao.GetIdByUser(name)
	test,_:=dao.GetHistoryTestByUserid(userid)
	views.ShowManyTestid(c,test)
}
func MyDoPractice(c *gin.Context) {
	token := c.GetHeader("Authorization")
	name := Username(token)
	userid, _ := dao.GetIdByUser(name)
	practice,_:=dao.GetHistoryPracticeByUserid(userid)
	views.ShowManyHistoryPractice(c,practice)
}
func MyUser(c *gin.Context){
	token := c.GetHeader("Authorization")
	name := Username(token)
	user, err := dao.GetUserByName(name)
	if err != nil {
		views.Fail(c, "用户查询失败")
		return
	}
	views.ShowUser(c,*user)
}