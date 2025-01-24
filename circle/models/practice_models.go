package models
type Practice struct {
	Practiceid int `gorm:"primaryKey;autoIncrement"`
	Content string `gorm:"type:text"`
	Difficulty string
	Circle string
    Name string
	Answer string
	Variety string
	Imageurl string
	Status string
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
	Name string
}
type Practicehistory struct {
    Userid int
	Practiceid int
	Answer string
}
