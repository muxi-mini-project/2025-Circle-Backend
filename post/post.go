package post

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"circle/database"
	"circle/user"
)

var lock sync.Mutex

type Comment struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
type Post struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Love     string `json:"love"`
	Collect  string `json:"collect"`
	Title    string `json:"title"`
	Imageurl string `json:"imageurl"`
	Postid   int    `json:"postid"`
	Circle   string `json:"circle"`
}
type Post1 struct {
	Name     string `json:"name"`
	Postid   int    `json:"postid"`
	Title    string `json:"title"`
	Love     int    `json:"love"`
	Imageurl string `json:"imageurl"`
	Circle   string `json:"circle"`
}

// Sendpost 发布帖子
// @Summary 创建新帖子
// @Description 创建一个新的帖子，用户需要在圈子内才可以发帖，支持上传一张图片。
// @Accept json
// @Produce json
// @Param Authorization header string true "User Authorization Token"
// @Param circleid formData string true "圈子ID"
// @Param content formData string true "帖子内容"
// @Param title formData string true "帖子标题"
// @Param imageurl formData string false "帖子图片URL"
// @Success 200 {object} map[string]interface{}{"message": string} "帖子创建成功"
// @Failure 404 {object} map[string]interface{}{"error": string} "用户不在该圈子内"
// @Router /sendpost [post]
func Sendpost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	name := user.Username(token)
	circleid := c.PostForm("circleid")
	var userid int
	_ = database.DB.QueryRow("SELECT id FROM user WHERE name=?", name).Scan(&userid)
	var count int
	query := "SELECT COUNT(*) FROM user_circles WHERE user_id = ? AND circle_id = ?"
	_ = database.DB.QueryRow(query, userid, circleid).Scan(&count)
	if count == 0 {
		c.JSON(404, gin.H{"error": "用户不在该圈子内"})
		return
	}
	content := c.PostForm("content")
	title := c.PostForm("title")
	var circle string
	query = `SELECT name FROM circles WHERE id = ?`
	_ = database.DB.QueryRow(query, circleid).Scan(&circle)
	imagePath := c.PostForm("imageurl")
	// 插入帖子数据
	_, _ = database.DB.Exec("INSERT INTO post (name, content,circle,love,collect,title,imageurl) VALUES (?,?,?,?,?,?,?)", name, content, circle, 0, 0, title, imagePath)
    _,_=database.DB.Exec("UPDATE topcircle SET number=number+1 WHERE circleid=?", circleid)
	c.JSON(200, gin.H{"message": "帖子创建成功"})
}

// Randpost 获取随机帖子
// @Summary 获取随机帖子
// @Description 获取指定圈子的随机帖子列表。
// @Accept json
// @Produce json
// @Param id formData string true "圈子ID"
// @Success 200 {object} map[string]interface{}{"message": []Post1} "返回随机帖子列表"
// @Router /randpost [post]
func Randpost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	id := c.PostForm("id")
	var circle string
	query := `SELECT name FROM circles WHERE id = ?`
	_ = database.DB.QueryRow(query, id).Scan(&circle)
	rows, _ := database.DB.Query("SELECT id,name, title, love, imageurl,circle FROM post WHERE circle=? ORDER BY RAND() LIMIT 10", circle)
	defer rows.Close()
	var pp []Post1
	for rows.Next() {
		var p Post1
		_ = rows.Scan(&p.Postid, &p.Name, &p.Title, &p.Love, &p.Imageurl, &p.Circle)
		pp = append(pp, p)
	}
	c.JSON(200, gin.H{"message": pp})
}

// Myfollowpost 获取我关注的帖子的随机列表
// @Summary 获取用户关注的帖子的随机列表
// @Description 获取用户关注的特定圈子的随机帖子。
// @Accept json
// @Produce json
// @Param Authorization header string true "User Authorization Token"
// @Param id formData string true "圈子ID"
// @Success 200 {object} map[string]interface{}{"message": []Post1} "返回我关注的随机帖子列表"
// @Router /myfollowpost [post]
func Myfollowpost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	username := user.Username(token)
	id := c.PostForm("id")
	var circle string
	query := `SELECT name FROM circles WHERE id = ?`
	_ = database.DB.QueryRow(query, id).Scan(&circle)
	rows, _ := database.DB.Query("SELECT p.id,p.circle,p.name,p.title,p.love,p.imageurl FROM post p JOIN fan f ON p.name = f.follower WHERE f.fan=? AND p.circle=? ORDER BY RAND() LIMIT 10", username, circle)
	defer rows.Close()
	var pp []Post1
	for rows.Next() {
		var p Post1
		_ = rows.Scan(&p.Postid, &p.Circle, &p.Name, &p.Title, &p.Love, &p.Imageurl)
		pp = append(pp, p)
	}
	c.JSON(200, gin.H{"message": pp})
}

// Readpost 阅读帖子
// @Summary 查看帖子详细信息
// @Description 查看指定ID的帖子详情，并记录阅读历史。
// @Accept json
// @Produce json
// @Param Authorization header string true "User Authorization Token"
// @Param id formData string true "帖子ID"
// @Success 200 {object} map[string]interface{}{"message": Post} "返回帖子详情"
// @Router /readpost [post]
func Readpost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	username := user.Username(token)
	id := c.PostForm("id")
	var pp Post
	query := "SELECT id,name,content,circle,love,collect,title,imageurl FROM post WHERE id=?"
	_ = database.DB.QueryRow(query, id).Scan(&pp.Postid, &pp.Name, &pp.Content, &pp.Circle, &pp.Love, &pp.Collect, &pp.Title, &pp.Imageurl)
	//+发帖人的头像
	_, _ = database.DB.Exec("INSERT INTO history (name,postid) VALUES (?,?)", username, pp.Postid)
	c.JSON(200, gin.H{"message": pp})
}

// Lovepost 点赞帖子
// @Summary 点赞帖子
// @Description 点赞指定ID的帖子。
// @Accept json
// @Produce json
// @Param Authorization header string true "User Authorization Token"
// @Param id formData string true "帖子ID"
// @Success 200 {object} map[string]interface{}{"love": int} "返回帖子点赞数量"
// @Router /lovepost [post]
func Lovepost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	id := c.PostForm("id")
	token := c.GetHeader("Authorization")
	username := user.Username(token)
	var name, title string
	_ = database.DB.QueryRow("SELECT title,name FROM post WHERE id=?", id).Scan(&title, &name)
	// 更新帖子的点赞数量
	message := username + "点赞你帖子-" + title
	_, _ = database.DB.Exec("UPDATE post SET love = love + 1 WHERE id = ?", id)
	_, _ = database.DB.Exec("INSERT INTO notifications (name,message,sendname) VALUES (?,?,?)", name, message, username)
	var love int
	_ = database.DB.QueryRow("SELECT love FROM post WHERE id=?", id).Scan(&love)
	c.JSON(http.StatusOK, gin.H{"love": love})
}

// Collectpost 收藏帖子
// @Summary 收藏帖子
// @Description 用户收藏帖子，并更新帖子收藏数。
// @Accept json
// @Produce json
// @Param Authorization header string true "User Authorization Token"
// @Param id formData string true "帖子ID"
// @Success 200 {object} map[string]interface{}{"collect": int} "返回当前收藏数"
// @Router /collectpost [post]
func Collectpost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	id := c.PostForm("id")
	token := c.GetHeader("Authorization")
	username := user.Username(token)
	var count int
	_ = database.DB.QueryRow("SELECT COUNT(*) FROM collect WHERE name=? AND postid=?", username, id).Scan(&count)
	if count > 0 {
		c.JSON(404, gin.H{"error": "已经收藏过了"})
		return
	}
	_, _ = database.DB.Exec("INSERT INTO collect (name,postid) VALUES (?,?)", username, id)
	var name, title string
	_ = database.DB.QueryRow("SELECT title,name FROM post WHERE id=?", id).Scan(&title, &name)
	// 更新帖子的点赞数量
	message := username + "收藏你帖子-" + title
	_, _ = database.DB.Exec("UPDATE post SET collect = collect + 1 WHERE id = ?", id)
	_, _ = database.DB.Exec("INSERT INTO notifications (name,message,sendname) VALUES (?,?,?)", name, message, username)
	var collect int
	_ = database.DB.QueryRow("SELECT collect FROM post WHERE id=?", id).Scan(&collect)
	c.JSON(http.StatusOK, gin.H{"collect": collect})
}

// Sharepost 分享帖子
// @Summary 分享帖子
// @Description 分享指定ID的帖子。
// @Accept json
// @Produce json
// @Param Authorization header string true "User Authorization Token"
// @Param id formData string true "帖子ID"
// @Success 200 {object} map[string]interface{}{"message": Post} "返回帖子详情"
// @Router /sharepost [post]
func Sharepost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	//获取帖子id
	id := c.PostForm("id")
	var name, content, circle, imageurl, title string
	var love, collect int
	_ = database.DB.QueryRow("SELECT title,name, content, circle, love, collect, imageurl FROM post WHERE id=?", id).Scan(&title, &name, &content, &circle, &love, &collect, &imageurl)
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

// Followpost 关注发帖子的人
// @Summary 关注发帖子的人
// @Description 用户关注指定ID的帖子的人。
// @Accept json
// @Produce json
// @Param Authorization header string true "User Authorization Token"
// @Param id formData string true "帖子ID"
// @Success 200 {object} map[string]interface{}{"message": string} "返回关注成功"
// @Router /followpost [post]
func Followpost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	id := c.PostForm("id")
	token := c.GetHeader("Authorization")
	username := user.Username(token)
	var name string //帖主
	_ = database.DB.QueryRow("SELECT name FROM post WHERE id=?", id).Scan(&name)
	var count int
	_ = database.DB.QueryRow("SELECT COUNT(*) FROM fan WHERE fan=? AND follower=?", username, name).Scan(&count)
	if count > 0 {
		c.JSON(404, gin.H{"error": "已经关注过了"})
		return
	}
	_, _ = database.DB.Exec("INSERT INTO fan (fan,follower) VALUES (?,?)", username, name)
	_, _ = database.DB.Exec("UPDATE user SET follower = follower + 1 WHERE name = ?", username)
	_, _ = database.DB.Exec("UPDATE user SET fan = fan + 1 WHERE name = ?", name)
	message := username + "关注了你"
	_, _ = database.DB.Exec("INSERT INTO notifications (name,message,sendname) VALUES (?,?,?)", name, message, username)
	c.JSON(200, gin.H{"message": "关注成功"})
}

// Commentpost 评论帖子
// @Summary 评论帖子
// @Description 用户评论帖子，并通知帖主。
// @Accept json
// @Produce json
// @Param Authorization header string true "User Authorization Token"
// @Param id formData string true "帖子ID"
// @Param content formData string true "评论内容"
// @Success 200 {object} map[string]interface{}{"message": string} "评论成功"
// @Failure 404 {object} map[string]interface{}{"error": string} "用户不在该圈子内"
// @Router /commentpost [post]
func Commentpost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	postid := c.PostForm("id")
	content := c.PostForm("content")
	token := c.GetHeader("Authorization")
	username := user.Username(token)
	var circleid int
	var circle string
	_ = database.DB.QueryRow("SELECT circle FROM post WHERE id=?", postid).Scan(&circle)
	_ = database.DB.QueryRow("SELECT id FROM circles WHERE name=?", circle).Scan(&circleid)
	var userid int
	_ = database.DB.QueryRow("SELECT id FROM user WHERE name=?", username).Scan(&userid)
	var count int
	query := "SELECT COUNT(*) FROM user_circles WHERE user_id = ? AND circle_id = ?"
	_ = database.DB.QueryRow(query, userid, circleid).Scan(&count)
	if count == 0 {
		c.JSON(404, gin.H{"error": "用户不在该圈子内"})
		return
	}
	var name, title string
	_ = database.DB.QueryRow("SELECT title,name FROM post WHERE id=?", postid).Scan(&title, &name)
	message := username + "评论了你的帖子-" + title
	_, _ = database.DB.Exec("INSERT INTO notifications (name,message,content,sendname) VALUES (?,?,?,?)", name, message, content, username)
	_, _ = database.DB.Exec("INSERT INTO comment (content,name,postid) VALUES (?,?,?)", content, username, postid)
	c.JSON(200, gin.H{"message": "评论创建成功"})
}

// Readcomment 阅读评论
// @Summary 查看帖子评论
// @Description 查看指定帖子的所有评论。
// @Accept json
// @Produce json
// @Param id formData string true "帖子ID"
// @Success 200 {object} map[string]interface{}{"comment": []Comment} "返回评论列表"
// @Router /readcomment [post]
func Readcomment(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	postid := c.PostForm("id")
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
