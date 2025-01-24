package models
type User struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"unique"`
	Password string
	Email string `gorm:"unique"`
	Imageurl string
	Discription string `gorm:"type:text"`
}
type UserPractice struct {
    Userid int `gorm:"primaryKey"`  //有主键位才有save,可更新可插入
    Practicenum int
	Correctnum int
}