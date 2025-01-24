package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"net/smtp"
	"math/rand"
	"time"
	"fmt"
	"sync"
	"io/ioutil"
	"strings"
	"encoding/base64"
	"circle/views"
	"circle/models"
	"circle/database"
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
func Getcode(c *gin.Context){
	email:=c.PostForm("email")
	var count int64
	database.DB.Model(&models.User{}).Where("email =?", email).Count(&count)
	if count>0{
		views.Fail(c,"该邮箱已注册")
		return
	}
	code:=generateVerificationCode()
	Getemail(email,code)
	lock.Lock()
	defer lock.Unlock()
    m[email]=code
	views.Success(c,"验证码已发送")
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
func Register(c *gin.Context){
	email:=c.PostForm("email")
	password:=c.PostForm("password")
	var count int64
	database.DB.Model(&models.User{}).Count(&count)
	name:="Circle_"+fmt.Sprintf("%04d",count+1)
	user:=models.User{
		Email:email,
		Password:password,
		Name:name,
	}
	database.DB.Create(&user)
	userpractice:=models.UserPractice{
		Userid:user.Id,
		Practicenum:0,
		Correctnum:0,
	}
	database.DB.Create(&userpractice)
	lock.Lock()
	defer lock.Unlock()
    delete(m,email)
	views.Success(c,"注册成功")
}
func Login(c *gin.Context){
	email:=c.PostForm("email")
	password:=c.PostForm("password")
	var user models.User
	err:=database.DB.Where("email = ?", email).First(&user).Error
	if err!=nil{
		views.Fail(c,"该邮箱未注册")
		return
	}
	if user.Password!=password{
		views.Fail(c,"密码错误")
		return
	}
	lock.Lock()
	defer lock.Unlock()
	token:=Token(user.Name)
	WhitelistedTokens[token]=1
	views.Success(c,token)
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
func Changepassowrd(c *gin.Context){
	lock.Lock()
	defer lock.Unlock()
	password:=c.PostForm("password")
	newpassword:=c.PostForm("newpassword")	
	token := c.GetHeader("Authorization")
	if _,ok:=WhitelistedTokens[token]; !ok {
		views.Fail(c,"token无效")
		return
	}
	name:=Username(token)
	var user models.User
	database.DB.Where("name = ?", name).First(&user)
	if user.Password!=password{
		views.Fail(c,"原密码错误")
		return
	}
	user.Password=newpassword
	database.DB.Save(&user)
	views.Success(c,"修改成功")
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
	var count int64
	database.DB.Model(&models.User{}).Where("name =?", newusername).Count(&count)
	if count>0{
		views.Fail(c,"该用户名已存在")
		return
	}
	name:=Username(token)
	var user models.User
	var practice models.Practice
	var practicecomment models.PracticeComment
	var test models.Test
	database.DB.Where("name = ?", name).First(&user)
	database.DB.Where("name = ?", name).First(&practice)
	database.DB.Where("name = ?", name).First(&practicecomment)
	database.DB.Where("name = ?", name).First(&test)
	user.Name=newusername
	practice.Name=newusername
	practicecomment.Name=newusername
	test.Name=newusername
	database.DB.Save(&user)
	database.DB.Save(&practice)
	database.DB.Save(&practicecomment)
	database.DB.Save(&test)
	newtoken:=Token(user.Name)
	WhitelistedTokens[newtoken]=1
	delete(WhitelistedTokens,token)
	views.Success(c,newtoken)
}
func Setphoto(c *gin.Context){
	lock.Lock()
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	if _,ok:=WhitelistedTokens[token]; !ok {
		views.Fail(c,"token无效")
		return
	}
	name:=Username(token)
	var user models.User
	database.DB.Where("name = ?", name).First(&user)
	user.Imageurl=c.PostForm("imageurl")
	database.DB.Save(&user)
	views.Success(c,"头像添加成功")
}
func Setdiscription(c *gin.Context){
	discription:=c.PostForm("discription")
	token := c.GetHeader("Authorization")
	name:=Username(token)
	var user models.User
	database.DB.Where("name = ?", name).First(&user)
	user.Discription=discription
	database.DB.Save(&user)
	views.Success(c,"简介修改成功")
}
func Getname(c *gin.Context){
	id:=c.PostForm("id")
	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	views.Success(c,user.Name)
}