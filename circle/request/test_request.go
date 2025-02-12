package request
import (
)
type Test struct{
	Discription string `json:"discription"`
	Circle string `json:"circle"`
	Testname string `json:"testname"`
}
type TestQuestion struct{
	Testid int `json:"testid"`
	Content string `json:"content"`
	Difficulty string `json:"difficulty"`
	Answer string `json:"answer"`
	Variety string `json:"variety"`
	Imageurl string `json:"imageurl"`
	Explain string `json:"explain"`
}
type Gettest struct{
	Testid int `json:"testid"`
}
type Score struct{
	Testid int `json:"testid"`
	Time int `json:"time"` 
	Correctnum int `json:"correctnum"`
}
type Commenttest struct{
	Testid int `json:"testid"`
	Content string `json:"content"`
}
type GetCircle struct{
	Circle string `json:"circle"`
}