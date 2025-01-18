package circles

import (
	"circle/database"
	"circle/user"

	"github.com/gin-gonic/gin"
	"sync"
)
var lock sync.Mutex
type Circle struct {
	Name        string `json:"name"`
	Discription string `json:"discription"`
	Variety     string `json:"variety"`
	Imageurl    string `json:"imageurl"`
	Id          int    `json:"id"`
}
type Top struct {
	Circle string `json:"circle"`
	Top int `json:"top"`
}

// Allcircles 获取所有某类型的圈子
// @Summary 获取所有某类型的圈子
// @Description 根据圈子的类型 (variety) 获取所有符合条件的圈子
// @Accept json
// @Produce json
// @Param variety formData string true "圈子的类型" // 获取指定类型的圈子
// @Success 200 {object} map[string]interface{}{"message": []Circle} "成功返回圈子列表"
// @Failure 500 {string} string "数据库查询错误" // 查询错误时的返回
// @Router /allcircles [post]
func Allcircles(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	variety := c.PostForm("variety")
	query := "SELECT name,id,discription,variety,imageurl FROM circles WHERE variety = ? "
	rows, _ := database.DB.Query(query,variety)
	defer rows.Close()
	var pp []Circle
	for rows.Next() {
		var p Circle
		_ = rows.Scan(&p.Name, &p.Id, &p.Discription, &p.Variety, &p.Imageurl)
		pp = append(pp, p)
	}
	c.JSON(200, gin.H{"message": pp})
}

// Mycircles 获取某用户加入的圈子
// @Summary 获取某用户加入的圈子
// @Description 根据圈子的类型 (variety) 获取所有符合条件的圈子
// @Accept json
// @Produce json
// @Param variety formData string true "圈子的类型" // 获取指定类型的圈子
// @Param name formData string true "用户名" // 获取指定类型的圈子
// @Success 200 {object} map[string]interface{}{"message": []Circle} "成功返回圈子列表"
// @Failure 500 {string} string "数据库查询错误" // 查询错误时的返回
// @Router /mycircles [post]
func Mycircles(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	variety := c.PostForm("variety")
	name:=c.PostForm("name")
	if name == "" {
	    token := c.GetHeader("Authorization")
	    name = user.Username(token)
	}
	var userid int
	_ = database.DB.QueryRow("SELECT id FROM user WHERE name=?", name).Scan(&userid)
	query := `
        SELECT c.name, c.id, c.discription, c.variety, c.imageurl
        FROM user_circles uc
        JOIN circles c ON uc.circle_id = c.id
        WHERE uc.user_id = ? AND c.variety = ?
    ` //c,uc表的别名
	rows, _ := database.DB.Query(query, userid,variety)
	defer rows.Close()

	var pp []Circle
	for rows.Next() {
		var p Circle
		_ = rows.Scan(&p.Name, &p.Id, &p.Discription, &p.Variety, &p.Imageurl)
		pp = append(pp, p)
	}

	c.JSON(200, gin.H{"message": pp})
}

func Topcircle(c *gin.Context){
	query:="SELECT circleid FROM topcircle ORDER BY number DESC "
	rows, _ := database.DB.Query(query)
	defer rows.Close()
	var pp []Top
	var i int
	for rows.Next(){
		var circleid int
		var p Top
		i++
		_ = rows.Scan(&circleid)
		_=database.DB.QueryRow("SELECT name FROM circles WHERE id = ?", circleid).Scan(&p.Circle)
		p.Top=i
		pp=append(pp,p)
	}
	c.JSON(200,gin.H{"message":pp})
}