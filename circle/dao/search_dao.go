package dao
import (
	"circle/models"
	"circle/database"
	"gorm.io/gorm"
)
type SearchDao struct {
	db *gorm.DB
}
func NewSearchDao(db *gorm.DB) *SearchDao {
	return &SearchDao{db: db}
}
func (ud *SearchDao) SearchCircle(circlekey string) ([]models.Circle) {
	var circle []models.Circle
	_ = database.DB.Where("name LIKE ?", "%"+circlekey+"%").Find(&circle).Error
	return circle
}
func (ud *SearchDao) SearchTest(testkey string) ([]models.Test) {
	var test []models.Test
	_ = database.DB.Where("testname LIKE ?", "%"+testkey+"%").Find(&test).Error
	return test
}
func (ud *SearchDao) SearchHistory(searchkey string,userid int)  {
	searchhistory:=models.SearchHistory{
		SearchKey:searchkey,
		Userid:userid,
	}
	database.DB.Create(&searchhistory)
}
func (ud *SearchDao) ShowSearchHistory(userid int) ([]models.SearchHistory){
	var searchhistory []models.SearchHistory
	_ = database.DB.Where("userid = ?", userid).Find(&searchhistory).Error
	return searchhistory
}
func (ud *SearchDao) DeleteHistory(id int)  {
	database.DB.Where("userid = ?", id).Delete(&models.SearchHistory{})
}
func (ud *SearchDao) SelectPracticeByCircle(circle string) ([]models.Practice) {
	var practice []models.Practice
	_ = database.DB.Where("status = ?", "approved").
		Where("circle = ?", circle).
		Order("RAND()").
		Limit(10).
		Find(&practice)
	return practice
}
func (ud *SearchDao) GetIdByUser(name string) (int, error) {
	var id int
	err := database.DB.Model(&models.User{}).Where("name = ?", name).Select("id").First(&id).Error
	return id, err
}
