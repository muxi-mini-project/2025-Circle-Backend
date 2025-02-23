package service
import (
	"circle/request"
	"circle/dao"
	"circle/models"

	"net/smtp"
	"time"
	"math/rand"
	"fmt"
	"sync"
	"io/ioutil"
	"encoding/json"

	"github.com/jordan-wright/email"
)
type UserServices struct {
	ud *dao.UserDao
}
func NewUserServices(ud *dao.UserDao) *UserServices {
	return &UserServices{
		ud: ud,
	}
}
var lock sync.Mutex
var m=make(map[string]string)
type Config struct {
	Email string `json:"email"`
}
func Getemail(ee string,VerificationCode string)  {
	data, _ := ioutil.ReadFile("data2.json")
	var config Config
	_ = json.Unmarshal(data, &config)
	m:=config.Email
	html := "<h1>验证码：" + VerificationCode + "</h1>"
	e := email.NewEmail()
	e.From = "luohuixi <2388287244@qq.com>"    
	e.To = []string{ee}         
	e.Subject = "验证码"                            
	e.Text = []byte("This is a plain text body.") 
	e.HTML = []byte(html)                    
	smtpHost := "smtp.qq.com"                                           
	smtpPort := "587"                                                             
	auth := smtp.PlainAuth("", "2388287244@qq.com", m, smtpHost)
	e.Send(smtpHost+":"+smtpPort, auth)  
}
func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano()) 
	code := rand.Intn(9000) + 1000   
	return fmt.Sprintf("%04d", code)
}
func (us *UserServices) Getcode(email request.Email){
	code := GenerateVerificationCode()
    Getemail(email.Email, code)
	lock.Lock()
    defer lock.Unlock()
    m[email.Email] = code
}
func (us *UserServices) Checkcode(email request.Email) bool {
	lock.Lock()
	defer lock.Unlock()
	return m[email.Email] == email.Code
}
func (us *UserServices) Register(user request.User) (string,bool) {
	count, err := us.ud.CountUsersByEmail(user.Email)
    if err != nil {
        return "查询数据库失败", false
    }
    if count > 0 {
        return "该邮箱已注册", false
    }
    totalUsers, err := us.ud.CountUsersByName("")
    if err != nil {
        return "查询数据库失败", false
    }
    name := "Circle_" + fmt.Sprintf("%04d", totalUsers+1)
    newuser := models.User{
        Email:    user.Email,
        Password: user.Password,
        Name:     name,
		Discription: "这里空空如也",
    }
    if err := us.ud.CreateUser(&newuser); err != nil {
        return "创建用户失败", false
    }
	return "注册成功", true
}
func (us *UserServices) Login(user request.User) (string,bool) {
	users, err := us.ud.GetUserByEmail(user.Email)
    if err != nil {
        return "该邮箱未注册",false
    }
    if users.Password != user.Password {
        return "密码错误", false
    }
    token,err:= GenerateToken(users.Name)
	if err != nil {
		return "生成 Token 失败", false
	}
	return token, true
}
// func (us *UserServices) Logout(token string) {
// 	lock.Lock()
// 	defer lock.Unlock()
// 	delete(WhitelistedTokens,token)
// }
func (us *UserServices) Changepassword(newpassword request.Newpassword,name string) (string,bool) {
	user, _ := us.ud.GetUserByName(name)
	user.Password=newpassword.Newpassword
	_=us.ud.UpdateUser(user)
	return "密码修改成功", true
}
func (us *UserServices) Changeusername(newusername request.Newusername,name string) (string,bool) {
	count,_:=us.ud.CountUsersByName(newusername.Newusername)
	if count>0{
		return "用户名已存在", false
	}
	user, err := us.ud.GetUserByName(name)
	if err != nil {
		return "用户查询失败", false
	}
	user.Name = newusername.Newusername
	err = us.ud.UpdateUser(user)
	if err != nil {
		return "用户名修改失败", false
	}
	newtoken,err:=GenerateToken(newusername.Newusername)
	if err != nil {
		return "生成 Token 失败", false
	}
	return newtoken, true
}
func (us *UserServices) Setphoto (name string, imageurl string) (string,bool) {
	user, err := us.ud.GetUserByName(name)
	if err != nil {
		return "用户查询失败", false
	}
	user.Imageurl = imageurl
	err = us.ud.UpdateUser(user)
	if err != nil {
		return "头像修改失败", false
	}
	return "头像修改成功", true
}
func (us *UserServices) Setdiscription (name string, discription string) (string,bool) {
	user, err := us.ud.GetUserByName(name)
	if err != nil {
		return "用户查询失败", false
	}
	user.Discription = discription
	err = us.ud.UpdateUser(user)
	if err != nil {
		return "简介修改失败", false
	}
	return "简介修改成功", true
}
func (us *UserServices) Getname (id request.Userid) (string,bool) {
	user, err := us.ud.GetUserByID(id.Userid)
	if err != nil {
		return "用户查询失败", false
	}
	return user.Name, true
}
func (us *UserServices) Mytest(name string) ([]models.Test) {
	userid, _ := us.ud.GetIdByUser(name)
	test,_:=us.ud.GetTestByUserid(userid)
	return test
}
func (us *UserServices) Mypractice(name string) ([]models.Practice) {
	userid, _ := us.ud.GetIdByUser(name)
	practice,_:=us.ud.GetPracticeByUserid(userid)
	return practice
}
func (us *UserServices) MyDoTest(name string) ([]models.Testhistory) {
	userid, _ := us.ud.GetIdByUser(name)
	test,_:=us.ud.GetHistoryTestByUserid(userid)
	return test
}
func (us *UserServices) MyDoPractice(name string) ([]models.Practicehistory) {
	userid, _ := us.ud.GetIdByUser(name)
	practice,_:=us.ud.GetHistoryPracticeByUserid(userid)
	return practice
}
func (us *UserServices) MyUser(name string) (models.User) {
	user, _ := us.ud.GetUserByName(name)
	return *user
}