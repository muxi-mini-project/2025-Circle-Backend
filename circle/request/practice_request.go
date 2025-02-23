package request
import()
type Practice struct{
	Variety string  `json:"variety"`
	Difficulty string `json:"difficulty"`
	Circle string  `json:"circle"`
	Imageurl string  `json:"imageurl"`
	Content string  `json:"content"`
	Answer string  `json:"answer"`
	Explain string  `json:"explain"`
}
type Option struct{
	Practiceid int  `json:"practiceid"`
	Content string  `json:"content"`
	Option string  `json:"option"`
}
type GetPractice struct{
	Circle string `json:"circle"`
	Practiceid int `json:"practiceid"`
}
type Comment struct{
	Practiceid int `json:"practiceid"`
	Content string `json:"content"`
}
type CheckAnswer struct{
	Circle string `json:"circle"`
	Practiceid int `json:"practiceid"`
	Answer string `json:"answer"`
	Time int `json:"time"`
}