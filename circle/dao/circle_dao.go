package dao
import (
	"circle/models"
	"circle/database"
	"gorm.io/gorm"
)
type CircleDao struct {
	db *gorm.DB
}
func NewCircleDao(db *gorm.DB) *CircleDao {
	return &CircleDao{db: db}
}
func (ud *CircleDao) CreateCircle(circle *models.Circle) error {
	return database.DB.Create(circle).Error
}
func (ud *CircleDao) SelectPendingCircle() (models.Circle, error) {
	var circles models.Circle
	err := database.DB.Where("status = ?", "pending").First(&circles).Error
	return circles, err
}
func (ud *CircleDao) GetCircleByID(circleid int) (models.Circle, error) {
	var circle models.Circle
	err := database.DB.Where("id = ?", circleid).First(&circle).Error
	return circle, err
}
func (ud *CircleDao) DeleteCircleByID(circleid int) error {
	return database.DB.Where("id = ?", circleid).Delete(&models.Circle{}).Error
}
func (ud *CircleDao) UpdateCircle(circle *models.Circle,id int) error {
	return database.DB.Model(&models.Circle{}).Where("id=?",id).Updates(circle).Error
}
func (ud *CircleDao) SelectCircle() ([]models.Circle, error) {
	var circles []models.Circle
	err := database.DB.Where("status=?","approved").Find(&circles).Error
	return circles, err
}
func (ud *CircleDao) FollowCircle(circleid int, userid int) error {
	circle:=models.FollowCircle{
		Circleid:circleid,
		Userid:userid,
	}
	return database.DB.Create(&circle).Error
}
func (ud *CircleDao) GetIdByUser(name string) (int, error) {
	var id int
	err := database.DB.Model(&models.User{}).Where("name = ?", name).Select("id").First(&id).Error
	return id, err
}