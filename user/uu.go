package user

import (
	"circle/database"
	"database/sql"
	//"encoding/json"
	"encoding/base64"
	"math/rand"
	"net/http"
	"strings"
	"sync"

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

// Register 用户注册
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

	username := c.PostForm("username")
	password := c.PostForm("password")

	lock.Lock()
	defer lock.Unlock()
	_, err := database.DB.Exec("INSERT INTO user (name, password, fan, follower) VALUES (?, ?, ?, ?)", username, password, 0, 0)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"code": 1, "message": "用户名已存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册成功"})
}

// Login 用户登录
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
