package dao
import (
	"circle/models"
	"circle/database"
)
func SearchCircle(circlekey string) ([]models.Circle) {
	var circle []models.Circle
	_ = database.DB.Where("name LIKE ?", "%"+circlekey+"%").Find(&circle).Error
	return circle
}
func SearchTest(testkey string) ([]models.Test) {
	var test []models.Test
	_ = database.DB.Where("testname LIKE ?", "%"+testkey+"%").Find(&test).Error
	return test
}
func SearchHistory(searchkey string,userid int)  {
	searchhistory:=models.SearchHistory{
		SearchKey:searchkey,
		Userid:userid,
	}
	database.DB.Create(&searchhistory)
}
func ShowSearchHistory(uesrid int) ([]models.SearchHistory){
	var searchhistory []models.SearchHistory
	_ = database.DB.Where("uesrid = ?", uesrid).Find(&searchhistory).Error
	return searchhistory
}
func DeleteHistory(id int)  {
	database.DB.Where("userid = ?", id).Delete(&models.SearchHistory{})
}
func SelectPracticeByCircle(circle string) ([]models.Practice) {
	var practice []models.Practice
	_ = database.DB.Where("status = ?", "approved").
		Where("circle = ?", circle).
		Order("RAND()").
		Limit(10).
		Find(&practice)
	return practice
}