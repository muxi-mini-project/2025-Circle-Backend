package models
type SearchHistory struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	SearchKey string 
	Userid int
}