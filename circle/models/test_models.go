package models
import "time"
type Test struct {
    Testid int `gorm:"primaryKey;autoIncrement"`	
	Name string
	Discription string
	Circle string
	Good int
	Allcomment int	
	Status string
	Createtime time.Time `gorm:"autoCreateTime"`
}
type TestQuestion struct {
	Testid int
	Questionid int `gorm:"primaryKey;autoIncrement"`	
	Content string `gorm:"type:text"`
	Difficulty string
	Answer string
	Variety string
	Imageurl string
}
type TestOption struct {
	Optionid int `gorm:"primaryKey;autoIncrement"`
	Content string `gorm:"type:text"`
	Practiceid int
	Option string
}
type Top struct {
	Topid int `gorm:"primaryKey;autoIncrement"`
	Userid int
	Correctnum int
	Time string
	Second int
	Testid int
}
type TestComment struct {
	Commentid int `gorm:"primaryKey;autoIncrement"`
	Content string `gorm:"type:text"`
	Testid int
	Userid int
}
type Testhistory struct {
	Testhisrotyid int `gorm:"primaryKey;autoIncrement"`
	Userid int
	Testid int
}