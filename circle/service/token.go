package service

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

// 定义密钥（用于签名和验证）
var secretKey = []byte("my_secret_key")

// 生成 Token
func GenerateToken(username string) (string, error) {
	// 创建 JWT 负载（Claims）
	claims := jwt.MapClaims{
		"username": username,
		"exp":     time.Now().AddDate(0, 1, 0).Unix(), // 过期时间1个月
		"iat":     time.Now().Unix(),                   // 签发时间
	}

	// 创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名 Token
	return token.SignedString(secretKey)
}

// 解析 Token
func ParseToken(tokenString string) (*jwt.MapClaims, error) {
	// 解析 Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("无效的签名方法")
		}
		return secretKey, nil
	})

	// 解析失败
	if err != nil {
		return nil, err
	}

	// 解析成功，返回 claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("无效的 Token")
}

// JWT 中间件
func JwtMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		//将不需要验证token的接口排除在外
		skipRoutes := []string{"/user/register", "/user/login","/user/getcode","/user/checkcode"}
        for _, route := range skipRoutes {
            if c.Request.URL.Path == route {
                // 如果是不需要验证的路由，直接继续处理请求
                c.Next()
                return
            }
        }
        // 从请求头中获取 Token
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(400, gin.H{"error": "未提供 Token"})
            c.Abort()
            return
        }
        // 去除 "Bearer " 前缀
        //tokenString := authHeader[7:]
        // 解析 Token
        claims, err := ParseToken(authHeader)
        if err != nil {
            c.JSON(400, gin.H{"error": "解析 Token 失败: " + err.Error()})
            c.Abort()
            return
        }
        // 获取 username
        username, exists := (*claims)["username"]
        if!exists {
            c.JSON(400, gin.H{"error": "Token 解析失败: 未找到 username 字段"})
            c.Abort()
            return
        }
		// 将 username 转换为 string
		name, ok := username.(string)
		if!ok {
			c.JSON(400, gin.H{"error": "Name in context is not a string"})
			return
		}
        // 将 username 加入 Gin 上下文
        c.Set("username", name)
        // 继续处理请求
        c.Next()
    }
}

