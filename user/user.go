package user

import (
	"circle/database"
	"circle/email"
	"database/sql"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var lock sync.Mutex
var BlacklistedTokens = make(map[string]int)
var WhitelistedTokens = make(map[string]int)

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

func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	code := rand.Intn(9000) + 1000   // 生成 1000 到 9999 的随机整数
	return fmt.Sprintf("%04d", code)
}

// Getcode 获取验证码
// @Summary 获取注册时的验证码
// @Description 根据邮箱获取验证码，如果邮箱没有注册过，则会生成一个新的验证码并发送到邮箱。
// @Accept json
// @Produce json
// @Param email formData string true "用户邮箱"
// @Success 200 {object} map[string]interface{}{"code": int, "message": string} "返回验证码发送状态"
// @Failure 409 {object} map[string]interface{}{"code": int, "message": string} "邮箱已注册"
// @Router /getcode [post]
func Getcode(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	ee := c.PostForm("email")
	var code int
	err := database.DB.QueryRow("SELECT code FROM email WHERE email=?", ee).Scan(&code)
	if err == sql.ErrNoRows {
		rightcode := generateVerificationCode()
		_, _ = database.DB.Exec("INSERT INTO email (email, code) VALUES (?, ?)", ee, rightcode)
		c.JSON(200, gin.H{"code": 0, "message": "验证码已发送"})
		email.Getemail(ee, rightcode)
		return
	}
	if code == 1 {
		c.JSON(409, gin.H{"code": 1, "message": "该邮箱已经注册"})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "验证码已发送"})
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册，注册成功后返回token。
// @Accept json
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param email formData string true "邮箱"
// @Param VerificationCode formData string true "验证码"
// @Success 200 {object} map[string]interface{}{"code": int, "message": string} "返回注册状态"
// @Failure 409 {object} map[string]interface{}{"code": int, "message": string} "用户名已存在"
// @Router /register [post]
func Register(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	token := c.GetHeader("Authorization")
	if token != "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "你已登录，请先登出"})
		return
	}
	lock.Lock()
	defer lock.Unlock()
	username := c.PostForm("username")
	password := c.PostForm("password")
	ee := c.PostForm("email")
	VerificationCode := c.PostForm("VerificationCode")
	imagePath:= c.PostForm("imageurl")
	var rightcode string
	_ = database.DB.QueryRow("SELECT code FROM email WHERE email=?", ee).Scan(&rightcode)
	if VerificationCode != rightcode {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "验证码错误"})
		return
	}
	var err error
	_ , err = database.DB.Exec("INSERT INTO user (name, password, fan, follower) VALUES (?, ?, ?, ?)", username, password, 0, 0)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"code": 1, "message": "用户名已存在"})
		return
	}
	_, _ = database.DB.Exec("INSERT INTO userimage (name, imageurl) VALUES (?,?)", username, imagePath)
	_, _ = database.DB.Exec("update email set code='1' where email=?", ee)
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册成功"})
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户输入用户名和密码登录，成功后返回Token。
// @Accept json
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} map[string]interface{}{"code": int, "message": string, "token": string} "返回登录成功及Token"
// @Failure 400 {object} map[string]interface{}{"code": int, "message": string} "已登录，或其他错误"
// @Failure 401 {object} map[string]interface{}{"code": int, "message": string} "密码错误"
// @Router /login [post]
func Login(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	token := c.GetHeader("Authorization")
	if token != "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "你已登录，请先登出"})
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	lock.Lock()
	defer lock.Unlock()
	var pw string
	err := database.DB.QueryRow("SELECT password FROM user WHERE name=?", username).Scan(&pw)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"code": 1, "message": "无当前用户名"})
		return
	}
	if pw != password {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "密码错误"})
		return
	}

	token = Token(username)
	WhitelistedTokens[token] = 1
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功", "token": token})
}

// Logout 用户登出
// @Summary 用户登出
// @Description 用户登出，登出后原token失效。
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}{"code": int, "message": string} "返回登出状态"
// @Failure 400 {object} map[string]interface{}{"code": int, "message": string} "你还没登录"
// @Router /logout [get]
func Logout(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "你还没登录"})
		return
	}

	lock.Lock()
	defer lock.Unlock()
	BlacklistedTokens[token] = 1 // 将原token加入黑名单
	delete(WhitelistedTokens, token)
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登出成功"})
}

// Change 修改密码
// @Summary 修改密码
// @Description 用户修改密码，需要提供当前密码和新密码。
// @Accept json
// @Produce json
// @Param password formData string true "当前密码"
// @Param changepassword formData string true "新密码"
// @Success 200 {object} map[string]interface{}{"code": int, "message": string} "返回修改密码成功"
// @Failure 400 {object} map[string]interface{}{"code": int, "message": string} "未登录，或其他错误"
// @Failure 401 {object} map[string]interface{}{"code": int, "message": string} "当前密码错误"
// @Router /change [post]
func Change(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	token := c.GetHeader("Authorization")
	lock.Lock()
	defer lock.Unlock()

	if token == "" || BlacklistedTokens[token] == 1 || WhitelistedTokens[token] != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "你还没登录"})
		return
	}

	password := c.PostForm("password")
	changepassword := c.PostForm("changepassword")
	uname := Username(token)

	var pw string
	err := database.DB.QueryRow("SELECT password FROM user WHERE name=?", uname).Scan(&pw)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "用户未找到"})
		return
	}

	if pw != password {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "密码错误"})
		return
	}

	_, _ = database.DB.Exec("UPDATE user SET password=? WHERE name=?", changepassword, uname)
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改成功"})
}
