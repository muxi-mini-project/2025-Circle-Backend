package dao
import (
	"circle/models"
	"circle/database"
)
func CreateCircle(circle *models.Circle) error {
	return database.DB.Create(circle).Error
}
func SelectPendingCircle() (models.Circle, error) {
	var circles models.Circle
	err := database.DB.Where("status = ?", "pending").First(&circles).Error
	return circles, err
}
func GetCircleByID(circleid int) (models.Circle, error) {
	var circle models.Circle
	err := database.DB.Where("id = ?", circleid).First(&circle).Error
	return circle, err
}
func DeleteCircleByID(circleid int) error {
	return database.DB.Where("id = ?", circleid).Delete(&models.Circle{}).Error
}
func UpdateCircle(circle *models.Circle,id int) error {
	return database.DB.Model(&models.Circle{}).Where("id=?",id).Updates(circle).Error
}
func SelectCircle() ([]models.Circle, error) {
	var circles []models.Circle
	err := database.DB.Where("status=?","approved").Find(&circles).Error
	return circles, err
}
func FollowCircle(circleid int, userid int) error {
	circle:=models.FollowCircle{
		Circleid:circleid,
		Userid:userid,
	}
	return database.DB.Create(&circle).Error
}