package search

import (
	"circle/database"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Postid   int    `json:"postid"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Circle   string `json:"circle"`
	Love     string `json:"love"`
	Collect  string `json:"collect"`
	Title    string `json:"title"`
	Imageurl string `json:"imageurl"`
}
type Circle struct {
	Name        string `json:"name"`
	Discription string `json:"discription"`
	Variety     string `json:"variety"`
	Imageurl    string `json:"imageurl"`
	Id          int    `json:"id"`
}

// Searchpost 搜索帖子
// @Summary 根据关键字搜索帖子
// @Description 根据用户输入的关键字搜索帖子，支持在标题、内容或用户名中查找。
// @Accept json
// @Produce json
// @Param search formData string true "搜索关键字"
// @Success 200 {object} map[string]interface{}{"post": []Post} "返回匹配的帖子列表"
// @Router /searchpost [post]
func Searchpost(c *gin.Context) {
	s := c.PostForm("search")
	query := "SELECT id,name,content,circle,love,collect,title,imageurl FROM post WHERE title LIKE ? OR content LIKE ? OR name LIKE ? ORDER BY RAND() LIMIT 10"
	se := "%" + s + "%"
	rows, _ := database.DB.Query(query, se, se,se)
	defer rows.Close()
	var post []Post
	for rows.Next() {
		var p Post
		rows.Scan(&p.Postid, &p.Name, &p.Content, &p.Circle, &p.Love, &p.Collect, &p.Title, &p.Imageurl)
		post = append(post, p)
	}
	c.JSON(200, gin.H{"post": post})
}

// Searchcircle 搜索圈子
// @Summary 根据关键字搜索圈子
// @Description 根据用户输入的关键字搜索圈子，支持在名称或简介中查找。
// @Accept json
// @Produce json
// @Param search formData string true "搜索关键字"
// @Success 200 {object} map[string]interface{}{"circle": []Circle} "返回匹配的圈子列表"
// @Router /searchcircle [post]
func Searchcircle(c *gin.Context) {
	s := c.PostForm("search")
	query := "SELECT id,name,discription,variety,imageurl FROM circles WHERE name LIKE ? OR discription LIKE ? ORDER BY RAND() LIMIT 10"
	se := "%" + s + "%"
	rows, _ := database.DB.Query(query, se, se)
	defer rows.Close()
	var circle []Circle
	for rows.Next() {
		var p Circle
		rows.Scan(&p.Id, &p.Name, &p.Discription, &p.Variety, &p.Imageurl)
		circle = append(circle, p)
	}
	c.JSON(200, gin.H{"circle": circle})
}
