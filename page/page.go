package page

import (
	"sync"

	"circle/database"
	"circle/user"

	"github.com/gin-gonic/gin"
)

var lock sync.Mutex

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Fan      int    `json:"fan"`
	Follower int    `json:"follower"`
}
type New struct {
	Message  string `json:"message"`
	Sendname string `json:"sendname"`
	Content  string `json:"content"`
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
type History struct {
	Postid int `json:"postid"`
}

// Information 获取用户的个人信息
// @Summary 获取用户的个人信息
// @Description 根据用户名获取用户的基本信息，包括用户名、ID、粉丝数、关注数、头像和未读通知数
// @Accept json
// @Produce json
// @Param name formData string false "用户名" // 用户名，若不传则从 Authorization 中获取
// @Success 200 {object} map[string]interface{}{"message": User, "imageurl": string, "未读信息": int} "成功返回用户信息"
// @Failure 400 {string} string "用户名不能为空" // 如果用户名为空，返回错误信息
// @Failure 500 {string} string "数据库查询错误" // 数据库查询错误时返回
// @Router /information [post]
func Information(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	name:=c.PostForm("name")
	if name == "" {
	    token := c.GetHeader("Authorization")
	    name = user.Username(token)
	}
	var u User
	var imageurl string
	query := "SELECT name,id,fan,follower FROM user WHERE name=?"
	_ = database.DB.QueryRow(query, name).Scan(&u.Name, &u.Id, &u.Fan, &u.Follower)
	query = "SELECT imageurl FROM userimage WHERE name=?"
	_ = database.DB.QueryRow(query, name).Scan(&imageurl)
	var count int
	_ = database.DB.QueryRow("SELECT count(*) FROM notifications WHERE name=? AND is_read=?", name, 0).Scan(&count)
	c.JSON(200, gin.H{
		"message":  u,
		"imageurl": imageurl,
		"未读信息":     count,
	})
}

// Myfollow 获取用户的关注列表
// @Summary 获取用户的关注列表
// @Description 根据用户名获取用户所关注的其他用户信息，包括用户名、ID、粉丝数和关注数
// @Accept json
// @Produce json
// @Param name formData string false "用户名" // 用户名，若不传则从 Authorization 中获取
// @Success 200 {object} map[string]interface{}{"message": []User} "成功返回关注的用户列表"
// @Failure 400 {string} string "用户名不能为空" // 如果用户名为空，返回错误信息
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /myfollow [post]
func Myfollow(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	name:=c.PostForm("name")
	if name == "" {
	    token := c.GetHeader("Authorization")
	    name = user.Username(token)
	}
	query := `
        SELECT u.name, u.id, u.fan,u.follower
        FROM fan f
        JOIN user u ON f.follower = u.name
        WHERE f.fan = ?
    `
	rows, _ := database.DB.Query(query, name)
	defer rows.Close()
	var uu []User
	for rows.Next() {
		var u User
		_ = rows.Scan(&u.Name, &u.Id, &u.Fan, &u.Follower)
		uu = append(uu, u)
	}
	c.JSON(200, gin.H{"message": uu})
}

// MyFan 获取用户的粉丝列表
// @Summary 获取用户的粉丝列表
// @Description 根据用户名获取用户的粉丝信息，包括用户名、ID、粉丝数和关注数
// @Accept json
// @Produce json
// @Param name formData string false "用户名" // 用户名，若不传则从 Authorization 中获取
// @Success 200 {object} map[string]interface{}{"message": []User} "成功返回粉丝的用户列表"
// @Failure 400 {string} string "用户名不能为空" // 如果用户名为空，返回错误信息
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /myfan [post]
func MyFan(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	name:=c.PostForm("name")
	if name == "" {
	    token := c.GetHeader("Authorization")
	    name = user.Username(token)
	}
	query := `
        SELECT u.name, u.id, u.fan,u.follower
        FROM fan f
        JOIN user u ON f.fan = u.name
        WHERE f.follower = ?
    `
	rows, _ := database.DB.Query(query, name)
	defer rows.Close()
	var uu []User
	for rows.Next() {
		var u User
		_ = rows.Scan(&u.Name, &u.Id, &u.Fan, &u.Follower)
		uu = append(uu, u)
	}
	c.JSON(200, gin.H{"message": uu})
}

// Mymessage 获取用户的消息
// @Summary 获取用户的消息
// @Description 获取用户的已读和未读消息，并将未读消息标记为已读
// @Accept json
// @Produce json
// @Param name formData string false "用户名" // 用户名，若不传则从 Authorization 中获取
// @Success 200 {object} map[string]interface{}{"未读消息": []New, "已读消息": []New} "成功返回已读和未读消息列表"
// @Failure 400 {string} string "用户名不能为空" // 如果用户名为空，返回错误信息
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /mymessage [get]
func Mymessage(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	name := user.Username(token)
	query := "SELECT message,sendname,content FROM notifications WHERE name=? AND is_read=? ORDER BY id DESC LIMIT 10"
	r, _ := database.DB.Query(query, name, 1)
	defer r.Close()
	var read []New
	for r.Next() {
		var n New
		_ = r.Scan(&n.Message, &n.Sendname, &n.Content)
		read = append(read, n)
	}
	query = "SELECT message,sendname,content FROM notifications WHERE name=? AND is_read=?"
	rows, _ := database.DB.Query(query, name, 0)
	defer rows.Close()
	var unread []New
	for rows.Next() {
		var n New
		_ = rows.Scan(&n.Message, &n.Sendname, &n.Content)
		unread = append(unread, n)
		_, _ = database.DB.Exec("UPDATE notifications SET is_read=? WHERE name=?", 1, name)
	}
	c.JSON(200, gin.H{
		"未读消息": unread,
		"已读消息": read,
	})
}

// Mypost 获取用户发布的帖子
// @Summary 获取用户发布的帖子
// @Description 获取用户发布的所有帖子，包括帖子ID、标题、内容等信息
// @Accept json
// @Produce json
// @Param name formData string false "用户名" // 用户名，若不传则从 Authorization 中获取
// @Success 200 {object} map[string]interface{}{"message": []Post} "成功返回用户发布的帖子列表"
// @Failure 400 {string} string "用户名不能为空" // 如果用户名为空，返回错误信息
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /mypost [post]
func Mypost(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	name:=c.PostForm("name")
	if name == "" {
	    token := c.GetHeader("Authorization")
	    name = user.Username(token)
	}
	rows, _ := database.DB.Query("SELECT id,name,circle, content, title, love, collect, imageurl FROM post WHERE name=?", name)
	defer rows.Close()
	var pp []Post
	for rows.Next() {
		var p Post
		_ = rows.Scan(&p.Postid, &p.Name, &p.Circle, &p.Content, &p.Title, &p.Love, &p.Collect, &p.Imageurl)
		pp = append(pp, p)
	}
	c.JSON(200, gin.H{"message": pp})
}

// Mycollect 获取用户收藏的帖子
// @Summary 获取用户收藏的帖子
// @Description 获取用户收藏的所有帖子，包括帖子ID、标题、内容等信息
// @Accept json
// @Produce json
// @Param name formData string false "用户名" // 用户名，若不传则从 Authorization 中获取
// @Success 200 {object} map[string]interface{}{"message": []Post} "成功返回用户收藏的帖子列表"
// @Failure 400 {string} string "用户名不能为空" // 如果用户名为空，返回错误信息
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /mycollect [post]
func Mycollect(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	name:=c.PostForm("name")
	if name == "" {
	    token := c.GetHeader("Authorization")
	    name = user.Username(token)
	}
	query := `
        SELECT p.id,p.name,p.circle,p.love,p.collect,p.title,p.content,p.imageurl
        FROM collect c
        JOIN post p ON c.postid = p.id
        WHERE c.name = ?
    `
	rows, _ := database.DB.Query(query, name)
	defer rows.Close()
	var pp []Post
	for rows.Next() {
		var p Post
		_ = rows.Scan(&p.Postid, &p.Name, &p.Circle, &p.Love, &p.Collect, &p.Title, &p.Content, &p.Imageurl)
		pp = append(pp, p)
	}
	c.JSON(200, gin.H{"message": pp})
}

// Myhistory 获取用户浏览历史
// @Summary 获取用户的浏览历史
// @Description 获取用户的浏览历史记录，包括浏览过的帖子ID
// @Accept json
// @Produce json
// @Param name formData string false "用户名" // 用户名，若不传则从 Authorization 中获取
// @Success 200 {object} map[string]interface{}{"message": []History} "成功返回用户的浏览历史"
// @Failure 400 {string} string "用户名不能为空" // 如果用户名为空，返回错误信息
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /myhistory [get]
func Myhistory(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	name := user.Username(token)
	query := "SELECT postid FROM history WHERE name=?"
	var hh []History
	rows, _ := database.DB.Query(query, name)
	defer rows.Close()
	for rows.Next() {
		var h History
		_ = rows.Scan(&h.Postid)
		hh = append(hh, h)
	}
	c.JSON(200, gin.H{"message": hh})
}

// Mytest 获取用户的考试记录
// @Summary 获取用户的考试记录
// @Description 获取用户的考试记录，包括用户参加的考试ID(即考过的卷子)
// @Accept json
// @Produce json
// @Param name formData string false "用户名" // 用户名，若不传则从 Authorization 中获取
// @Success 200 {object} map[string]interface{}{"testid": []int} "成功返回用户的考试ID列表"
// @Failure 400 {string} string "用户名不能为空" // 如果用户名为空，返回错误信息
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /mytest [post]
func Mytest(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	name:=c.PostForm("name")
	if name == "" {
	    token := c.GetHeader("Authorization")
	    name = user.Username(token)
	}
	var userid int
	_ = database.DB.QueryRow("SELECT id FROM user WHERE name=?", name).Scan(&userid)
	query:="SELECT testid FROM user_test WHERE userid=?"
	rows, _ := database.DB.Query(query, userid)
	defer rows.Close()
	var t []int
	for rows.Next() {
		var testid int
		_ = rows.Scan(&testid)
		t=append(t,testid)
	}
	c.JSON(200, gin.H{"testid": t})
}

// Myowntest 获取用户自己出的试卷
// @Summary 获取用户自己出的试卷
// @Description 获取用户自己出的试卷，包括考试ID
// @Accept json
// @Produce json
// @Param name formData string false "用户名" // 用户名，若不传则从 Authorization 中获取
// @Success 200 {object} map[string]interface{}{"testid": []int} "成功返回用户的自有考试ID列表"
// @Failure 400 {string} string "用户名不能为空" // 如果用户名为空，返回错误信息
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /myowntest [post]
func Myowntest(c *gin.Context){
	lock.Lock()
	defer lock.Unlock()
	name:=c.PostForm("name")
	if name == "" {
	    token := c.GetHeader("Authorization")
	    name = user.Username(token)
	}
	query:="SELECT testid FROM owntest WHERE name=?"
	rows, _ := database.DB.Query(query, name)
	defer rows.Close()
	var t []int
	for rows.Next() {
		var testid int
		_ = rows.Scan(&testid)
		t=append(t,testid)
	}
	c.JSON(200, gin.H{"testid": t})
}

// Mylevel 获取用户在某个圈子中的等级
// @Summary 获取用户在某个圈子中的等级
// @Description 获取用户在指定圈子中的等级，基于用户的发布帖子和自有考试数量
// @Accept json
// @Produce json
// @Param name formData string false "用户名" // 用户名，若不传则从 Authorization 中获取
// @Param circleid formData string true "圈子ID" // 圈子ID
// @Success 200 {object} map[string]interface{}{"circle": string, "level": int} "成功返回圈子和用户等级"
// @Failure 400 {string} string "用户名或圈子ID不能为空" // 如果缺少参数，返回错误信息
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /mylevel [post]
func Mylevel(c *gin.Context){
	lock.Lock()
	defer lock.Unlock()
	circleid:=c.PostForm("circleid")
	var circle string
	_=database.DB.QueryRow("SELECT name FROM circles WHERE id=?", circleid).Scan(&circle)
	name:=c.PostForm("name")
	if name == "" {
	    token := c.GetHeader("Authorization")
	    name = user.Username(token)
	}
	var count,count2 int
	_=database.DB.QueryRow("SELECT COUNT(*) FROM post WHERE name=? AND circle=?",name,circle).Scan(&count)
	_=database.DB.QueryRow("SELECT COUNT(*) FROM owntest WHERE name=? AND circle=?",name,circle).Scan(&count2)
    c.JSON(200,gin.H{
		"circle":circle,
		"level":count+count2,
	})
}