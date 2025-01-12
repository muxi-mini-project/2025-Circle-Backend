package post

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"circle/database"
)

var lock sync.Mutex

type Comment struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func Sendpost(c *gin.Context) {
	//获取帖子内容(图片只能一张)
	lock.Lock()
	defer lock.Unlock()
	name := c.PostForm("name")
	circle := c.PostForm("circle")
	var userid, circleid int
	_ = database.DB.QueryRow("SELECT id FROM user WHERE name=?", name).Scan(&userid)
	_ = database.DB.QueryRow("SELECT id FROM circles WHERE name=?", circle).Scan(&circleid)
	var count int
	query := "SELECT COUNT(*) FROM user_circles WHERE user_id = ? AND circle_id = ?"
	_ = database.DB.QueryRow(query, userid, circleid).Scan(&count)
	if count == 0 {
		c.JSON(404, gin.H{"error": "不属于这个圈子"})
		return
	}
	content := c.PostForm("content")
	title := c.PostForm("title")
	file, err := c.FormFile("image")
	if err != nil {
		// 处理错误，例如返回 Bad Request 响应
		c.JSON(http.StatusBadRequest, gin.H{"error": "表单错误"})
		return
	}
	// 构建保存路径
	imagePath := "./uploads/" + file.Filename
	// 保存上传的文件
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		// 处理文件保存错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存图片失败"})
		return
	}
	// 插入帖子数据
	_, _ = database.DB.Exec("INSERT INTO post (name, content,circle,love,collect,title,imageurl) VALUES (?,?,?,?,?,?,?)", name, content, circle, 0, 0, title, imagePath)

	c.JSON(200, gin.H{"message": "帖子创建成功"})
}

func Readpost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	//获取帖子id
	title := c.PostForm("title")
	var name, content, circle, imageurl string
	var love, collect int
	err := database.DB.QueryRow("SELECT name,content,circle,love,collect,imageurl FROM post WHERE title=?", title).Scan(&name, &content, &circle, &love, &collect, &imageurl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无该帖子"})
		return
	}
	c.JSON(200, gin.H{
		"name":     name,
		"content":  content,
		"circle":   circle,
		"love":     love,
		"collect":  collect,
		"title":    title,
		"imageurl": imageurl,
	})
}

func Lovepost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	title := c.PostForm("title")
	// 更新帖子的点赞数量
	_, _ = database.DB.Exec("UPDATE post SET love = love + 1 WHERE title = ?", title)
	var love int
	_ = database.DB.QueryRow("SELECT love FROM post WHERE title=?", title).Scan(&love)
	c.JSON(http.StatusOK, gin.H{"love": love})
}

func Collectpost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	title := c.PostForm("title")
	// 更新帖子的点赞数量
	_, _ = database.DB.Exec("UPDATE post SET collect = collect + 1 WHERE title = ?", title)
	var collect int
	_ = database.DB.QueryRow("SELECT collect FROM post WHERE title=?", title).Scan(&collect)
	c.JSON(http.StatusOK, gin.H{"collect": collect})
} //有问题

func Sharepost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	//获取帖子id
	title := c.PostForm("title")
	var name, content, circle, imageurl string
	var love, collect int
	_ = database.DB.QueryRow("SELECT name, content, circle, love, collect, imageurl FROM post WHERE title=?", title).Scan(&name, &content, &circle, &love, &collect, &imageurl)
	c.JSON(200, gin.H{
		"name":     name,
		"content":  content,
		"circle":   circle,
		"love":     love,
		"collect":  collect,
		"title":    title,
		"imageurl": imageurl,
	})
}

func Followpost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	title := c.PostForm("title")
	uname := c.PostForm("name") //粉丝
	_, _ = database.DB.Exec("UPDATE user SET follower = follower + 1 WHERE name = ?", uname)
	var name string //帖主
	_ = database.DB.QueryRow("SELECT name FROM post WHERE title=?", title).Scan(&name)
	_, _ = database.DB.Exec("UPDATE user SET fan = fan + 1 WHERE name = ?", name)
	_, _ = database.DB.Exec("INSERT INTO fan (fan,follower) VALUES (?,?)", uname, name)
	c.JSON(200, gin.H{"message": "关注成功"})
}

func Commentpost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	name := c.PostForm("name")
	circle := c.PostForm("circle")
	var userid, circleid int
	_ = database.DB.QueryRow("SELECT id FROM user WHERE name=?", name).Scan(&userid)
	_ = database.DB.QueryRow("SELECT id FROM circles WHERE name=?", circle).Scan(&circleid)
	var count int
	query := "SELECT COUNT(*) FROM user_circles WHERE user_id = ? AND circle_id = ?"
	_ = database.DB.QueryRow(query, userid, circleid).Scan(&count)
	if count == 0 {
		c.JSON(404, gin.H{"error": "不属于这个圈子"})
		return
	}
	title := c.PostForm("title")
	content := c.PostForm("content")
	var postid int
	_ = database.DB.QueryRow("SELECT postid FROM post WHERE title=?", title).Scan(&postid)
	_, _ = database.DB.Exec("INSERT INTO comment (content,name,postid) VALUES (?,?,?)", content, name, postid)
	c.JSON(200, gin.H{"message": "评论创建成功"})
}

func Readcomment(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	title := c.PostForm("title")
	var postid int
	_ = database.DB.QueryRow("SELECT postid FROM post WHERE title=?", title).Scan(&postid)
	rows, _ := database.DB.Query("SELECT name,content FROM comment WHERE postid=?", postid)
	defer rows.Close()
	var com []Comment
	for rows.Next() {
		var cc Comment
		_ = rows.Scan(&cc.Name, &cc.Content)
		com = append(com, cc)
	}
	c.JSON(200, gin.H{"comment": com})
}
