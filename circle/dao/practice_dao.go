package dao

import (
	"circle/models"
	"circle/database"
	"fmt"
	"gorm.io/gorm"
)
type PracticeDao struct {
	db *gorm.DB
}
func NewPracticeDao(db *gorm.DB) *PracticeDao {
	return &PracticeDao{db: db}
}

func (ud *PracticeDao) CreatePractice(practice *models.Practice) error {
	err := database.DB.Create(practice).Error
	if err != nil {
		return fmt.Errorf("创建 practice 失败: %w", err)
	}
	return nil
}
func (ud *PracticeDao) CreatePracticeOption(option *models.PracticeOption) error {
	err := database.DB.Create(option).Error
	if err != nil {
		return fmt.Errorf("创建 PracticeOption 失败: %w", err)
	}
	return nil
}
func (ud *PracticeDao) GetPracticeByCircle(circle string) (models.Practice, error) {
	var practice models.Practice
	err := database.DB.Where("status = ?", "approved").
		Where("circle = ?", circle).
		Order("RAND()").
		First(&practice).Error
	if err != nil {
		return practice, fmt.Errorf("查询 Practice 失败: %w", err)
	}
	return practice, nil
}
func (ud *PracticeDao) GetPracticeOptionsByPracticeID(practiceid int) ([]models.PracticeOption, error) {
	var options []models.PracticeOption
	err := database.DB.Where("practiceid = ?", practiceid).Find(&options).Error
	if err != nil {
		return nil, fmt.Errorf("查询 PracticeOptions 失败: %w", err)
	}
	return options, nil
}
func (ud *PracticeDao) CreatePracticeComment(comment *models.PracticeComment) error {
	err := database.DB.Create(comment).Error
	if err != nil {
		return fmt.Errorf("创建 PracticeComment 失败: %w", err)
	}
	return nil
}
func (ud *PracticeDao) GetPracticeCommentsByPracticeID(practiceid int) ([]models.PracticeComment, error) {
	var comments []models.PracticeComment
	err := database.DB.Where("practiceid = ?", practiceid).Find(&comments).Error
	if err != nil {
		return nil, fmt.Errorf("查询 PracticeComments 失败: %w", err)
	}
	return comments, nil
}
func (ud *PracticeDao) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("name = ?", username).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}
	return &user, nil
}
func (ud *PracticeDao) GetUserPracticeByUserID(userID int,circle string) (*models.UserPractice, error) {
	var userpractice models.UserPractice
	err := database.DB.Where("userid = ?", userID).Where("circle = ?", circle).First(&userpractice).Error
	if err == gorm.ErrRecordNotFound {
		userpractice = models.UserPractice{
			Userid: userID,
			Circle: circle,
			Practicenum: 0,
			Correctnum: 0,
			Alltime: 0,
		}
		database.DB.Create(&userpractice)
	}
	return &userpractice, nil
}
func (ud *PracticeDao) UpdateUserPractice(userpractice *models.UserPractice) error {
	err := database.DB.Save(userpractice).Error
	if err != nil {
		return fmt.Errorf("更新用户练习记录失败: %w", err)
	}
	return nil
}
func (ud *PracticeDao) CreatePracticeHistory(history *models.Practicehistory) error {
	err := database.DB.Create(history).Error
	if err != nil {
		return fmt.Errorf("创建练习历史失败: %w", err)
	}
	return nil
}
func (ud *PracticeDao) GetApprovedPracticesByCircle(circle string) ([]models.Practice, error) {
	var practices []models.Practice
	err := database.DB.Where("status = ?", "approved").Where("circle = ?", circle).Limit(5).Find(&practices).Error
	if err != nil {
		return nil, fmt.Errorf("获取练习记录失败: %w", err)
	}
	return practices, nil
}
func (ud *PracticeDao) Showrank(id int,circle string) int {
	var userPractices []models.UserPractice
    database.DB.Where("circle = ?", circle).
	   Order("CAST(correctnum AS float) / CAST(practicenum AS float) DESC").
       Order("practicenum DESC").
       Order("CAST(alltime AS float) / CAST(practicenum AS float) ASC").
       Find(&userPractices)
    var rank int
    for i, user := range userPractices {
      if user.Userid == id {
          rank = i + 1 
          break
        }  
    }
    return rank
}
func (ud *PracticeDao) GetPracticeByPracticeID(practiceid int) (models.Practice) {
	var practice models.Practice
	_ = database.DB.Where("practiceid = ?", practiceid).First(&practice)
	return practice
}
func (ud *PracticeDao) UpdatePractice(practice *models.Practice) error {
	err := database.DB.Save(practice).Error
	if err != nil {
		return fmt.Errorf("更新练习记录失败: %w", err)
	}
	return nil
}
func (ud *PracticeDao) GetIdByUser(name string) (int, error) {
	var id int
	err := database.DB.Model(&models.User{}).Where("name = ?", name).Select("id").First(&id).Error
	return id, err
}
