package dao

import (
	"circle/models"
	"circle/database"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func GetUserByName(name string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("name = ?", name).First(&user).Error
	return &user, err
}

func GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := database.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func CountUsersByEmail(email string) (int64, error) {
	var count int64
	err := database.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count, err
}

func CountUsersByName(name string) (int64, error) {
	var count int64
	err := database.DB.Model(&models.User{}).Where("name = ?", name).Count(&count).Error
	return count, err
}

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}

func GetIdByUser(name string) (int, error) {
	var id int
	err := database.DB.Model(&models.User{}).Where("name = ?", name).Select("id").First(&id).Error
	return id, err
}

func CreateUserpractice(userpractice *models.UserPractice) error {
	return database.DB.Create(userpractice).Error
}

func GetTestByUserid(userid int) ([]models.Test, error) {
	var usertest []models.Test
	err := database.DB.Where("userid = ?", userid).Find(&usertest).Error
	return usertest, err
}

func GetPracticeByUserid(userid int) ([]models.Practice, error) {
	var userpractice []models.Practice
	err := database.DB.Where("userid = ?", userid).Find(&userpractice).Error
	return userpractice, err
}

func GetHistoryTestByUserid(userid int) ([]models.Testhistory, error) {
	var historytest []models.Testhistory
	err := database.DB.Where("userid = ?", userid).Find(&historytest).Error
	return historytest, err
}

func GetHistoryPracticeByUserid(userid int) ([]models.Practicehistory, error) {
	var historypractice []models.Practicehistory
	err := database.DB.Where("userid = ?", userid).Find(&historypractice).Error
	return historypractice, err
}