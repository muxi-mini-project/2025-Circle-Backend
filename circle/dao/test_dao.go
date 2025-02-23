package dao

import (
	"circle/models"
	"circle/database"
	"fmt"
	"gorm.io/gorm"
)
type TestDao struct {
	db *gorm.DB
}
func NewTestDao(db *gorm.DB) *TestDao {
	return &TestDao{db: db}
}

func (ud *TestDao) CreateTest(test *models.Test) (int, error) {
	err := database.DB.Create(test).Error
	if err != nil {
		return 0, fmt.Errorf("创建测试记录失败: %w", err)
	}
	return test.Testid, nil
}
func (ud *TestDao) CreateQuestion(question *models.TestQuestion) (int, error) {
	err := database.DB.Create(question).Error
	if err != nil {
		return 0, fmt.Errorf("创建问题记录失败: %w", err)
	}
	return question.Questionid, nil
}
func (ud *TestDao) CreateTestOption(option *models.TestOption) (int, error) {
	err := database.DB.Create(option).Error
	if err != nil {
		return 0, fmt.Errorf("创建测试选项记录失败: %w", err)
	}
	return option.Optionid, nil
}
func (ud *TestDao) GetTestByID(testid int) (models.Test, error) {
	var test models.Test
	err := database.DB.Where("testid = ?", testid).First(&test).Error
	if err != nil {
		return test, fmt.Errorf("获取测试信息失败: %w", err)
	}
	return test, nil
}
func (ud *TestDao) RecordTestHistory(testid int, userId int) error {
	testHistory := models.Testhistory{
		Testid: testid,
		Userid: userId,
	}
	err := database.DB.Create(&testHistory).Error
	if err != nil {
		return fmt.Errorf("记录测试历史失败: %w", err)
	}
	return nil
}
func (ud *TestDao) GetQuestionsByTestID(testid int) ([]models.TestQuestion, error) {
	var questions []models.TestQuestion
	err := database.DB.Where("testid = ?", testid).Find(&questions).Error
	if err != nil {
		return nil, fmt.Errorf("获取测试题目失败: %w", err)
	}
	return questions, nil
}
func (ud *TestDao) GetTestOptionsByPracticeID(practiceid int) ([]models.TestOption, error) {
	var options []models.TestOption
	err := database.DB.Where("practiceid = ?", practiceid).Find(&options).Error
	if err != nil {
		return nil, fmt.Errorf("获取测试选项失败: %w", err)
	}
	return options, nil
}
func (ud *TestDao) SaveTopRecord(top models.Top) error {
	err := database.DB.Create(&top).Error
	if err != nil {
		return fmt.Errorf("保存成绩记录失败: %w", err)
	}
	return nil
}
func (ud *TestDao) GetTopByTestID(testid int) ([]models.Top, error) {
	var tops []models.Top
	err := database.DB.Order("correctnum desc, time asc").
		Where("testid = ?", testid).
		Limit(10).
		Find(&tops).Error
	return tops, err
}
func (ud *TestDao) CreateTestComment(comment *models.TestComment) error {
	return database.DB.Create(comment).Error
}
func (ud *TestDao) GetTestComments(testid int) ([]models.TestComment, error) {
	var comments []models.TestComment
	if err := database.DB.Where("testid = ?", testid).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
func (ud *TestDao) GetTestByTestID(testid int) (models.Test,error){
	var test models.Test
	err := database.DB.Where("testid = ?", testid).First(&test).Error
	if err != nil {
		return test, fmt.Errorf("获取测试信息失败: %w", err)
	}
	return test, nil
}
func (ud *TestDao) UpdateTest(test *models.Test) error {
	return database.DB.Save(test).Error
}
func (ud *TestDao) RecommentTest(circle string) ([]models.Test){
    var test []models.Test
	if circle=="" {
		_= database.DB.Order("RAND()").Limit(10).Find(&test).Error
	}else{
	    _= database.DB.Where("circle = ?", circle).Order("RAND()").Limit(10).Find(&test).Error
	}
	return test
}
func (ud *TestDao) HotTest(circle string) ([]models.Test){
	var test []models.Test
	if circle=="" {
		_= database.DB.Order("good desc").Limit(10).Find(&test).Error
	}else{
		_= database.DB.Where("circle = ?", circle).Order("good desc").Limit(10).Find(&test).Error
	}
	return test
}
func (ud *TestDao) NewTest(circle string) ([]models.Test){
	var test []models.Test
	if circle=="" {
		_= database.DB.Order("createtime desc").Limit(10).Find(&test).Error
	}else{
		_= database.DB.Where("circle = ?", circle).Order("createtime desc").Limit(10).Find(&test).Error
	}
	return test
}
func (ud *TestDao) FollowCircleTest(userid int) ([]models.Test){
	var test []models.Test
	var circleid []int
	var circlename []string
	_= database.DB.Model(&models.FollowCircle{}).Where("userid = ?", userid).Pluck("circleid", &circleid).Error //pluck表示查询单个数据
	_= database.DB.Model(&models.Circle{}).Where("id in (?)", circleid).Pluck("name", &circlename).Error
	_= database.DB.Where("circle in (?)", circlename).Order("RAND()").Limit(10).Find(&test).Error  //in表示查询多个数据
	return test
}
func (ud *TestDao) GetIdByUser(name string) (int, error) {
	var id int
	err := database.DB.Model(&models.User{}).Where("name = ?", name).Select("id").First(&id).Error
	return id, err
}

func (ud *TestDao) GetUserByName(name string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("name = ?", name).First(&user).Error
	return &user, err
}