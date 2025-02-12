package service
import(
    "circle/request"
	"circle/dao"
	"circle/models"
)
type TestServices struct {
	ud *dao.TestDao
}
func NewTestServices(ud *dao.TestDao) *TestServices {
	return &TestServices{
		ud: ud,
	}
}
func (us *TestServices) CreateTest(name string,get request.Test) int {
	id,_:=us.ud.GetIdByUser(name)
	test := models.Test{
		Userid:       id,
		Testname: get.Testname,
		Discription: get.Discription,
		Circle:     get.Circle,
		Good:       0,
		Status:     "approved", // 待审核
	}
	id, err := us.ud.CreateTest(&test)
	if err != nil {
		return -1
	}
	return id
}
func (us *TestServices) Createquestion(get request.TestQuestion) int {
	question := models.TestQuestion{
		Content:   get.Content,
		Testid:     get.Testid,
		Difficulty: get.Difficulty,
		Answer:    get.Answer,
		Variety:   get.Variety,
		Imageurl:  get.Imageurl,
		Explain:   get.Explain,
	}
	id, err := us.ud.CreateQuestion(&question)
	if err != nil {
		return -1
	}
	return id
}
func (us *TestServices) Createtestoption(get request.Option) int {
	options := models.TestOption{
		Content:   get.Content,
		Practiceid: get.Practiceid,
		Option:    get.Option,
	}
	id, err := us.ud.CreateTestOption(&options)
	if err != nil {
		return -1
	}
	return id
}
func (us *TestServices) Gettest(name string,get request.Gettest) models.Test {
	test, _ := us.ud.GetTestByID(get.Testid)
	id,_:=us.ud.GetIdByUser(name)
	_=us.ud.RecordTestHistory(get.Testid, id)
	return test
}
func (us *TestServices) Getquestion(get request.Gettest) []models.TestQuestion {
	questions, _:= us.ud.GetQuestionsByTestID(get.Testid)
	return questions
}
func (us *TestServices) Gettestoption(get request.GetPractice) []models.TestOption {
	options, _:= us.ud.GetTestOptionsByPracticeID(get.Practiceid)
	return options
}
func (us *TestServices) Getscore(name string,get request.Score) string {
	user, _ := us.ud.GetUserByName(name)
	top := models.Top{
		Userid:     user.Id,
		Correctnum: get.Correctnum,
		Time:       get.Time,
		Testid:     get.Testid,
	}
	_ = us.ud.SaveTopRecord(top)
	return "成功"
}
func (us *TestServices) Showtop(get request.Gettest) []models.Top{
	tops, _ := us.ud.GetTopByTestID(get.Testid)
	return tops
}
func (us *TestServices) Commenttest(name string,get request.Commenttest) string {
	user,_:=us.ud.GetUserByName(name)
	comment := models.TestComment{
		Content: get.Content,
		Testid:  get.Testid,
		Userid:  user.Id,
	}
	_ = us.ud.CreateTestComment(&comment)
	return "成功"
}
func (us *TestServices) Gettestcomment(get request.Gettest)  []models.TestComment{
	comments, _:= us.ud.GetTestComments(get.Testid)
	return comments
}
func (us *TestServices) Lovetest(get request.Gettest) string {
	test,_:= us.ud.GetTestByTestID(get.Testid)
	test.Good++
	_ = us.ud.UpdateTest(&test)
	return "点赞成功"
}
func (us *TestServices) RecommentTest(get request.GetCircle) []models.Test{
	test:=us.ud.RecommentTest(get.Circle)
	return test
}
func (us *TestServices) HotTest(get request.GetCircle) []models.Test{
	test:=us.ud.HotTest(get.Circle)
	return test
}
func (us *TestServices) NewTest(get request.GetCircle) []models.Test{
	test:=us.ud.NewTest(get.Circle)
	return test
}
func (us *TestServices) FollowCircleTest(name string) []models.Test{
	userid,_:=us.ud.GetIdByUser(name)
	test:=us.ud.FollowCircleTest(userid)
	return test
}