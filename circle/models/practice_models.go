package models
type Practice struct {
	Practiceid int `gorm:"primaryKey;autoIncrement"`
	Content string `gorm:"type:text"`
	Difficulty string
	Circle string
    Userid int
	Answer string
	Variety string
	Imageurl string
	Status string
	Explain string `gorm:"type:text"`
	Good int
}
type PracticeOption struct {
	Optionid int `gorm:"primaryKey;autoIncrement"`
	Content string `gorm:"type:text"`
	Practiceid int
	Option string
}
type PracticeComment struct {
	Commentid int `gorm:"primaryKey;autoIncrement"`
	Content string `gorm:"type:text"`
	Practiceid int
	Userid int
}
type Practicehistory struct {
    Userid int
	Practiceid int
	Answer string
}
type UserPractice struct {
	Id int `gorm:"primaryKey;autoIncrement"` //有主键位才有save,可更新可插入
    Userid int   
    Practicenum int
	Correctnum int
	Alltime int
	Circle string
}
