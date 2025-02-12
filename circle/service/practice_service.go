package service
import (
    "circle/models"
	"circle/request"
	"circle/dao"

	"fmt"
)
type PracticeServices struct {
	ud *dao.PracticeDao
}
func NewPracticeServices(ud *dao.PracticeDao) *PracticeServices {
	return &PracticeServices{
		ud: ud,
	}
}
func (us *PracticeServices) Createpractice(name string,practice request.Practice) int {
	id,_:=us.ud.GetIdByUser(name)
	practices := models.Practice{
		Userid:      id,
		Content:   practice.Content,
		Difficulty: practice.Difficulty,
		Circle:    practice.Circle,
		Answer:    practice.Answer,
		Variety:   practice.Variety,
		Imageurl:  practice.Imageurl,
		Explain:    practice.Explain,
		Good:       0,
		Status:    "approved", // 待审核
	}
	err := us.ud.CreatePractice(&practices)
	if err != nil {
		return -1
	}
	return practices.Practiceid
}
func (us *PracticeServices) Createoption(option request.Option) string {
	options := models.PracticeOption{
		Content:   option.Content,
		Practiceid: option.Practiceid,
		Option:    option.Option,
	}
	err := us.ud.CreatePracticeOption(&options)
	if err != nil {
		return "创建失败"
	}
	return "等待审核"
}
func (us *PracticeServices) GetPractice(get request.GetPractice) models.Practice {
	var practice models.Practice
	if get.Circle == "" {
		practice= us.ud.GetPracticeByPracticeID(get.Practiceid)
	}else {
	    practice, _ = us.ud.GetPracticeByCircle(get.Circle)
	}
	return practice
}
func (us *PracticeServices) GetPracticeOption(get request.GetPractice) []models.PracticeOption {
	options, _ := us.ud.GetPracticeOptionsByPracticeID(get.Practiceid)
	return options
}
func (us *PracticeServices) CommentPractice(name string,comment request.Comment) string {
	id,_:=us.ud.GetIdByUser(name)
	comments := models.PracticeComment{
		Content:    comment.Content,
		Practiceid: comment.Practiceid,
		Userid:       id,
	}
	err := us.ud.CreatePracticeComment(&comments)
	if err != nil {
		return "评论失败"
	}
	return "评论成功"
}
func (us *PracticeServices) GetComment(comment request.GetPractice) []models.PracticeComment {
	comments, _ := us.ud.GetPracticeCommentsByPracticeID(comment.Practiceid)
	return comments
}
func (us *PracticeServices) CheckAnswer(name string,get request.CheckAnswer) string {
	user, err := us.ud.GetUserByUsername(name)
	if err != nil {
		return "用户不存在"
	}
	userpractice, err := us.ud.GetUserPracticeByUserID(user.Id, get.Circle)
	if err != nil {
		return "用户练习记录不存在"
	}
	userpractice.Alltime += get.Time
	userpractice.Practicenum++
	if get.Answer == "true" {
		userpractice.Correctnum++
	}
	err = us.ud.UpdateUserPractice(userpractice)
	if err != nil {
		return "更新练习记录失败"
	}
	practicehistory := models.Practicehistory{
		Practiceid: get.Practiceid,
		Userid:     user.Id,
		Answer:     get.Answer,
	}
	err = us.ud.CreatePracticeHistory(&practicehistory)
	if err != nil {
		return "创建练习记录失败"
	}
	return "成功"
}
func (us *PracticeServices) Getrank(name string,get request.GetPractice) string {
	id,_:=us.ud.GetIdByUser(name)
	rank:= us.ud.Showrank(id,get.Circle)
	return fmt.Sprintf("%d", rank)
}
func (us *PracticeServices) GetUserPractice(name string,get request.GetPractice) models.UserPractice {
	id,_:=us.ud.GetIdByUser(name)
	userpractice, _ := us.ud.GetUserPracticeByUserID(id,get.Circle)
	return *userpractice
}
func (us *PracticeServices) Lovepractice(get request.GetPractice) string {
	practice := us.ud.GetPracticeByPracticeID(get.Practiceid)
	practice.Good++
	_ = us.ud.UpdatePractice(&practice)
	return "点赞成功"
}