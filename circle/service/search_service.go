package service
import (
	"circle/dao"
	"circle/models"
)
type SearchServices struct {
	ud *dao.SearchDao
}
func NewSearchServices(ud *dao.SearchDao) *SearchServices {
	return &SearchServices{
		ud: ud,
	}
}
func (us *SearchServices) SearchCircle(name string,circlekey string) []models.Circle {
	userid,_:=us.ud.GetIdByUser(name)
    circle:=us.ud.SearchCircle(circlekey)
	us.ud.SearchHistory(circlekey,userid)
	return circle
}
func (us *SearchServices) SearchTest(name string,testkey string) []models.Test {
	userid,_:=us.ud.GetIdByUser(name)
    test:=us.ud.SearchTest(testkey)
	us.ud.SearchHistory(testkey,userid)
	return test
}
func (us *SearchServices) SearchHistory(name string) []models.SearchHistory {
	userid,_:=us.ud.GetIdByUser(name)
	search:=us.ud.ShowSearchHistory(userid)
	return search
}
func (us *SearchServices) DeleteHistory(name string) {
	userid,_:=us.ud.GetIdByUser(name)
	us.ud.DeleteHistory(userid)
}
func (us *SearchServices) SearchPractice(circle string) []models.Practice {
	practice:=us.ud.SelectPracticeByCircle(circle)
	return practice
}