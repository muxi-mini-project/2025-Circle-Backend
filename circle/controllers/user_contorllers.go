package controllers

import (
	"circle/request"
	"circle/service"

	"github.com/gin-gonic/gin"
)
type UserControllers struct {
	us *service.UserServices
}
func NewUserControllers(us *service.UserServices) *UserControllers {
	return &UserControllers{
		us: us,
	}
}
func (uc *UserControllers) Getcode(c *gin.Context) {
    var email request.Email
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	uc.us.Getcode(email)
	c.JSON(200, gin.H{"success": "验证码已发送"})
}
func (uc *UserControllers) Checkcode(c *gin.Context){
	var email request.Email
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	ok:=uc.us.Checkcode(email)
	if !ok{
		c.JSON(400, gin.H{"error": "验证码错误"})
		return
	}
	c.JSON(200, gin.H{"success": "验证码正确"})
}
func (uc *UserControllers) Register(c *gin.Context) {
    var user request.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	message,ok:=uc.us.Register(user) 
    if !ok {
		c.JSON(400, gin.H{"error": message})
		return
	}
    c.JSON(200, gin.H{"success": message})
}
func (uc *UserControllers) Login(c *gin.Context) {
    var user request.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	message,ok:=uc.us.Login(user)
    if !ok {
		c.JSON(400, gin.H{"error": message})
		return
	}
    c.JSON(200, gin.H{"success": message})
}
func (uc *UserControllers) Logout(c *gin.Context){
	//uc.us.Logout(token)
	c.JSON(200, gin.H{"success": "登出成功"})
}
func (uc *UserControllers) Changepassowrd(c *gin.Context) {
    var newpassword request.Newpassword
	if err := c.ShouldBindJSON(&newpassword); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
	message,ok:=uc.us.Changepassword(newpassword,n)
	if !ok {
		c.JSON(400, gin.H{"error": message})
		return
	}
	c.JSON(200, gin.H{"success": message})
}
func (uc *UserControllers) Changeusername(c *gin.Context){
	var newusername request.Newusername
	if err := c.ShouldBindJSON(&newusername); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
    message,ok:=uc.us.Changeusername(newusername,n)
	if !ok {
		c.JSON(400, gin.H{"error": message})
		return
	}
	c.JSON(200, gin.H{"success": message})
}
func (uc *UserControllers) Setphoto(c *gin.Context) {
	name,_:=c.Get("username")
	n,_:=name.(string)
	var imageurl request.Imageurl
	if err := c.ShouldBindJSON(&imageurl); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	message,ok:=uc.us.Setphoto(n, imageurl.Imageurl)
	if !ok {
		c.JSON(400, gin.H{"error": message})
		return
	}
	c.JSON(200, gin.H{"success": message})
}
func (uc *UserControllers) Setdiscription(c *gin.Context) {
	var discription request.Discription
	if err := c.ShouldBindJSON(&discription); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	name,_:=c.Get("username")
	n,_:=name.(string)
	message,ok:=uc.us.Setdiscription(n, discription.Discription)
	if !ok {
		c.JSON(400, gin.H{"error": message})
		return
	}
	c.JSON(200, gin.H{"success": message})
}
func (uc *UserControllers) Getname(c *gin.Context) {
	var id request.Userid
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(400, gin.H{"error": "无效的参数"})
		return
	}
	message,ok:=uc.us.Getname(id)
	if !ok{
		c.JSON(400, gin.H{"error": message})
		return 
	}
	c.JSON(200, gin.H{"success": message})
}
func (uc *UserControllers) Mytest(c *gin.Context) {
	name,_:=c.Get("username")
	n,_:=name.(string)
	test:=uc.us.Mytest(n)
    c.JSON(200, gin.H{"success": test})
}
func (uc *UserControllers) Mypractice(c *gin.Context) {
	name,_:=c.Get("username")
	n,_:=name.(string)
	practice:=uc.us.Mypractice(n)
	c.JSON(200, gin.H{"success": practice})
}
func (uc *UserControllers) MyDoTest(c *gin.Context) {
	name,_:=c.Get("username")
	n,_:=name.(string)
	test:=uc.us.MyDoTest(n)
	c.JSON(200, gin.H{"success": test})
}
func (uc *UserControllers) MyDoPractice(c *gin.Context) {
	name,_:=c.Get("username")
	n,_:=name.(string)
	practice:=uc.us.MyDoPractice(n)
	c.JSON(200, gin.H{"success": practice})
}
func (uc *UserControllers) MyUser(c *gin.Context){
	name,_:=c.Get("username")
	n,_:=name.(string)
	user:=uc.us.MyUser(n)
	c.JSON(200, gin.H{"success": user})
}