package paper

import (
	"circle/database"
	"circle/user"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var lock sync.Mutex

type Question struct {
	Questionid    int
	QuestionText  string
	CorrectAnswer string
	Options       []Option
}
type Option struct {
	Optionid    int
	OptionLable string
	OptionText  string
}
type Test struct {
	Testid int
	Name string
	Title string
	Imageurl string
	Circle string
	Discription string
}
type Ownquestion struct {
	Questionid int
	Questiontext string
	Correctanswer string
	Type string
	Score int
	Difficulty int
	Ownoptions []OwnOption
}
type OwnOption struct {
	Optionid int
	Optionlabel string
	Optiontext string
}
type Comment struct {
	Id int
	Name string
	Comment string
	Testid int
}

func Getrandomquestions(circlename string, limit int) []Question {
	lock.Lock()
	defer lock.Unlock()
	query := `
        SELECT id, question_text, correct_answers 
        FROM questions 
        WHERE circle_name = ? 
        ORDER BY RAND() 
        LIMIT ?`
	rows, _ := database.DB.Query(query, circlename, limit)
	defer rows.Close()
	var count int
	query = "SELECT COUNT(*) FROM questions"
	_ = database.DB.QueryRow(query).Scan(&count)
	questions := make([]Question, count+1)
	for rows.Next() {
		var q Question
		_ = rows.Scan(&q.Questionid, &q.QuestionText, &q.CorrectAnswer)
		questions[q.Questionid] = q
	}
	return questions
}
func GetOptions(questionID int) []Option {
	lock.Lock()
	defer lock.Unlock()
	query := `SELECT id, option_label, option_text FROM options WHERE question_id = ?`
	rows, _ := database.DB.Query(query, questionID)
	defer rows.Close()
	var options []Option
	for rows.Next() {
		var o Option
		_ = rows.Scan(&o.Optionid, &o.OptionLable, &o.OptionText)
		options = append(options, o)
	}
	return options
}

// Getquestion 获取问题和选项(加入圈子)
// @Summary 获取指定测试的题目和选项(加入圈子)
// @Description 根据测试ID获取题目和选项，返回测试的所有题目和对应的选项
// @Accept json
// @Produce json
// @Param id formData string true "测试ID" // 测试ID
// @Success 200 {object} map[string]interface{}{"message": []Question} "成功返回问题和选项列表"
// @Failure 400 {string} string "测试ID无效" // 如果测试ID无效
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /getquestion [post]
func Getquestion(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	id := c.PostForm("id")
	var circle string
	query := "SELECT name FROM circles WHERE id = ?"
	_ = database.DB.QueryRow(query, id).Scan(&circle)
	questions := Getrandomquestions(circle, 10)
	for _, q := range questions {
		options := GetOptions(q.Questionid)
		questions[q.Questionid].Options = options
	}
	//删除空切片
	var newq []Question
	for _, value := range questions {
		if value.Questionid != 0 {
			newq = append(newq, value)
		}
	}
	c.JSON(200, gin.H{"message": newq})
}

// Getscore 提交分数(加入圈子)
// @Summary 提交分数并记录考试成绩(加入圈子)
// @Description 提交用户的考试分数，并在分数大于等于100时保存到数据库
// @Accept json
// @Produce json
// @Param id formData string true "圈子ID" // 圈子ID
// @Param score formData string true "考试分数" // 用户的考试分数
// @Success 200 {string} string "考试成功"
// @Failure 400 {string} string "分数不及格" // 如果分数小于100
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /getscore [post]
func Getscore(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	circleid := c.PostForm("id")
	token := c.GetHeader("Authorization")
	name := user.Username(token)
	score := c.PostForm("score")
	s, _ := strconv.Atoi(score)
	if s < 100 {
		c.JSON(200, gin.H{"message": "失败"})
		return
	}
	_,_=database.DB.Exec("UPDATE topcircle SET number=number+1 WHERE circleid=?", circleid)
	var userid int
	query := `SELECT id FROM user WHERE name = ?`
	_ = database.DB.QueryRow(query, name).Scan(&userid)
	query = `SELECT id FROM circles WHERE name = ?`
	_, _ = database.DB.Exec("INSERT INTO user_circles (user_id,circle_id) VALUES (?,?)", userid, circleid)
	c.JSON(200, gin.H{"message": "考试成功"})
}

// Generatetest 生成考试
// @Summary 创建一个新的考试
// @Description 创建一个新的考试，指定考试的圈子、标题、描述等
// @Accept json
// @Produce json
// @Param circle formData string true "圈子名称" // 圈子名称
// @Param title formData string true "考试标题" // 考试标题
// @Param discription formData string true "考试描述" // 考试描述
// @Param imageurl formData string true "考试图片URL" // 考试图片URL
// @Success 200 {object} map[string]interface{}{"testid": int, "message": string} "成功返回考试ID"
// @Failure 400 {string} string "圈子不存在或未加入圈子" // 如果圈子不存在或用户未加入圈子
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /generatetest [post]
func Generatetest(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	name := user.Username(token)
	circle := c.PostForm("circle")
	title := c.PostForm("title")
	discription := c.PostForm("discription")
	imageurl := c.PostForm("imageurl")
	var count,cc int
	_= database.DB.QueryRow("SELECT COUNT(*) FROM circles WHERE name = ?", circle).Scan(&count)
	if count == 0 {
		c.JSON(400, gin.H{"message": "圈子不存在"})
		return
	}
	var circleid,userid int
	_ = database.DB.QueryRow("SELECT id FROM user WHERE name = ?", name).Scan(&userid)
	_ = database.DB.QueryRow("SELECT id FROM circles WHERE name = ?", circle).Scan(&circleid)
	_= database.DB.QueryRow("SELECT COUNT(*) FROM user_circles WHERE user_id = ? ADN circle_id=?",userid,circleid).Scan(&cc)
    if cc==0 {
		c.JSON(400, gin.H{"message": "你没有加入该圈子"})
		return 
	}
	_, _ = database.DB.Exec("INSERT INTO owntest (name, circle,title, discription, imageurl) VALUES (?,?,?,?,?)", name, circle, title, discription, imageurl)
	var testid int
	_ = database.DB.QueryRow("SELECT testid FROM owntest WHERE name = ?", name).Scan(&testid)
	c.JSON(200, gin.H{
		"testid":  testid,
		"message": "成功",
	})
}

// Addquestionandoption 添加问题和选项
// @Summary 添加问题和选项
// @Description 添加问题和选项，包括题目文本、题目类型、正确答案、分数、难度
// @Accept json
// @Produce json
// @Param testid formData string true "考试ID" // 考试ID
// @Param questiontext formData string true "题目文本" // 题目文本
// @Param type formData string true "题目类型" // 题目类型
// @Param correctanswer formData string true "正确答案" // 正确答案
// @Param score formData string true "分数" // 分数
// @Param difficulty formData string true "难度" // 难度
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "考试ID无效" // 如果考试ID无效
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /addquestionandoption [post]
func Addquestionandoption(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	testid := c.PostForm("testid")
	testid2, _ := strconv.Atoi(testid)
	questiontext := c.PostForm("questiontext")
	t := c.PostForm("type")
	correctanswer := c.PostForm("correctanswer")
	score := c.PostForm("score")
	score2, _ := strconv.Atoi(score)
	difficulty := c.PostForm("difficulty")
	difficulty2, _ := strconv.Atoi(difficulty)
	_,_=database.DB.Exec("INSERT INTO ownquestion (testid,questiontext,type,correctanswer,score,difficulty) VALUES(?,?,?,?,?,?)",testid2,questiontext,t,correctanswer,score2,difficulty2)
	var questionid int
	_ = database.DB.QueryRow("SELECT questionid FROM ownquestion WHERE questiontext = ?", questiontext).Scan(&questionid)
	if t == "判断题" {
		_, _ = database.DB.Exec("INSERT INTO ownoption (questionid,optiontext,optionlabel) VALUES (?,?,?)", questionid, "true", "true")
		_, _ = database.DB.Exec("INSERT INTO ownoption (questionid,optiontext,optionlabel) VALUES (?,?,?)", questionid, "false", "false")
	} else {
		aa := c.PostForm("A")
		bb := c.PostForm("B")
		cc := c.PostForm("C")
		dd := c.PostForm("D")
		_, _ = database.DB.Exec("INSERT INTO ownoption (questionid,optiontext,optionlabel) VALUES (?,?,?)", questionid, aa, "A")
		_, _ = database.DB.Exec("INSERT INTO ownoption (questionid,optiontext,optionlabel) VALUES (?,?,?)", questionid, bb, "B")
		_, _ = database.DB.Exec("INSERT INTO ownoption (questionid,optiontext,optionlabel) VALUES (?,?,?)", questionid, cc, "C")
		_, _ = database.DB.Exec("INSERT INTO ownoption (questionid,optiontext,optionlabel) VALUES (?,?,?)", questionid, dd, "D")
	}
	c.JSON(200, gin.H{"message": "成功"})
}

// Findtest 查找考试
// @Summary 查找考试
// @Description 查找考试，返回考试ID列表
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}{"testid": []string} "成功返回考试ID列表"
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /findtest [post]
func Findtest(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	name := user.Username(token)
	var userid int
	_ = database.DB.QueryRow("SELECT id FROM user WHERE name = ?", name).Scan(&userid)
	query := `SELECT ot.testid FROM owntest ot
	          JOIN circles c ON ot.circle = c.name
			  WHERE c.id IN (
			  SELECT circle_id FROM user_circles WHERE user_id = ?)
			  ORDER BY RAND() LIMIT 10
			  `
	rows, _ := database.DB.Query(query, userid)
	defer rows.Close()
	var t []string
	for rows.Next() {
		var testid string
		_=rows.Scan(&testid)
		t = append(t, testid)
	}
	c.JSON(200, gin.H{"testid": t})
}

// Gettest 获取考试
// @Summary 获取考试
// @Description 获取考试，返回考试信息
// @Accept json
// @Produce json
// @Param testid formData string true "考试ID" // 考试ID
// @Success 200 {object} map[string]interface{}{"test": Test, "question": []Ownquestion} "成功返回考试信息"
// @Failure 400 {string} string "考试ID无效" // 如果考试ID无效
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /gettest [post]
func Gettest(c *gin.Context){
    lock.Lock()
	defer lock.Unlock()
	testid := c.PostForm("testid")
	token := c.GetHeader("Authorization")
	name := user.Username(token)
	var userid,done int
	_ = database.DB.QueryRow("SELECT id FROM user WHERE name = ?", name).Scan(&userid)
	_= database.DB.QueryRow("SELECT done FROM user_test WHERE testid = ? AND userid = ?", testid, userid).Scan(&done)
	if done==1{
		c.JSON(200, gin.H{"message": "你已经考过该试卷了"})
		return
	}
	var t Test
	query:="SELECT testid,name,title,imageurl,circle,discription FROM owntest WHERE testid = ?"
	_= database.DB.QueryRow(query, testid).Scan(&t.Testid,&t.Name,&t.Title,&t.Imageurl,&t.Circle,&t.Discription)	
	query = "SELECT questionid,questiontext,type,correctanswer,score,difficulty FROM ownquestion WHERE testid = ?"
	rows, _ := database.DB.Query(query, testid)
	defer rows.Close()
	var qq []Ownquestion
	for rows.Next() {
		var q Ownquestion
		_=rows.Scan(&q.Questionid,&q.Questiontext,&q.Type,&q.Correctanswer,&q.Score,&q.Difficulty)
		query = "SELECT optionid,optiontext,optionlabel FROM ownoption WHERE questionid = ?"
		rows2, _ := database.DB.Query(query, q.Questionid)
		defer rows2.Close()
		var oo []OwnOption
		for rows2.Next() {
			var o OwnOption
			_=rows2.Scan(&o.Optionid,&o.Optiontext,&o.Optionlabel)
			oo = append(oo, o)
		}
		q.Ownoptions = oo
		qq = append(qq, q)
	}
	c.JSON(200, gin.H{"test": t, "question": qq})
}

// Getownscore 获取自己的分数
// @Summary 获取自己的分数
// @Description 获取自己的分数
// @Accept json
// @Produce json
// @Param testid formData string true "考试ID" // 考试ID
// @Param score formData string true "分数" // 分数
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "考试ID无效" // 如果考试ID无效
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /getownscore [post]
func Getownscore(c *gin.Context){
	lock.Lock()	
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	name := user.Username(token)
	var userid int
	_ = database.DB.QueryRow("SELECT id FROM user WHERE name = ?", name).Scan(&userid)
	score:=c.PostForm("score")
	testid:=c.PostForm("testid")
    ss,_:= strconv.Atoi(score)
	_,_ = database.DB.Exec("INSERT INTO user_test (userid,testid,score,done) VALUES (?,?,?,?)",userid,testid,ss,1)
	c.JSON(200, gin.H{"message": "成功"})
}

// Commenttest 评论考试
// @Summary 评论考试
// @Description 评论考试
// @Accept json
// @Produce json
// @Param testid formData string true "考试ID" // 考试ID
// @Param comment formData string true "评论" // 评论
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "考试ID无效" // 如果考试ID无效
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /commenttest [post]
func Commenttest(c *gin.Context){
	lock.Lock()
	defer lock.Unlock()
	token := c.GetHeader("Authorization")
	name := user.Username(token)
	testid := c.PostForm("testid")
	comment := c.PostForm("comment")
	_,_=database.DB.Exec("INSERT INTO commenttest (name,comment,testid) VALUES (?,?,?)",name,comment,testid)
	c.JSON(200, gin.H{"message": "成功"})
}

// Getcommenttest 获取考试评论
// @Summary 获取考试评论
// @Description 获取考试评论
// @Accept json
// @Produce json
// @Param testid formData string true "考试ID" // 考试ID
// @Success 200 {object} map[string]interface{}{"message": []Comment} "成功返回考试评论"
// @Failure 400 {string} string "考试ID无效" // 如果考试ID无效
// @Failure 500 {string} string "数据库查询错误" // 查询数据库时发生错误
// @Router /getcommenttest [post]
func Getcommenttest(c *gin.Context){
	lock.Lock()
	defer lock.Unlock()
	testid := c.PostForm("testid")
	query := "SELECT id,name,comment,testid FROM commenttest WHERE testid = ?"
	rows, _ := database.DB.Query(query, testid)
	defer rows.Close()
	var pp []Comment
	for rows.Next() {
		var p Comment
		_=rows.Scan(&p.Id,&p.Name,&p.Comment,&p.Testid)
		pp = append(pp, p)
	}
	c.JSON(200, gin.H{"message": pp})
}