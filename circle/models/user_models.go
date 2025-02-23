package models
type User struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"unique"`
	Password string
	Email string `gorm:"unique"`
	Imageurl string
	Discription string `gorm:"type:text"`
}