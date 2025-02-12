package request
import()
type Email struct {
	Email string `json:"email"`
	Code string `json:"code"`
}
type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
type Newpassword struct {
	Newpassword string `json:"newpassword"`
}
type Newusername struct {
	Newusername string `json:"newusername"`
}
type Imageurl struct {
	Imageurl string `json:"imageurl"`
}
type Discription struct {
	Discription string `json:"discription"`
}
type Userid struct {
	Userid int `json:"id"`
}