package request
import()
type CreateCircle struct {
	Name string `json:"name"`
    Discription string `json:"discription"`
    Imageurl string `json:"imageurl"`
}
type ApproveCircle struct {
	Circleid int `json:"circleid"`
    Decide string `json:"decide"`
}
type Circleid struct{
	Circleid int `json:"circleid"`
}