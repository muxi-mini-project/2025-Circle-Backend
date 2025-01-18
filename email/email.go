package email

import (
	"net/smtp"
	"github.com/jordan-wright/email"
)


func Getemail(ee string,VerificationCode string)  {
	html := "<h1>验证码：" + VerificationCode + "</h1>"
	// 创建一个新的邮件对象
	e := email.NewEmail()
	e.From = "luohuixi <2388287244@qq.com>"       // 发件人
	e.To = []string{ee}          // 收件人(可多人)
	e.Subject = "验证码"                             // 邮件主题
	e.Text = []byte("This is a plain text body.") // 文本内容
	e.HTML = []byte(html)                         // HTML 内容

	// 设置 SMTP 服务器
	smtpHost := "smtp.qq.com"                                                     // QQ SMTP 服务器
	smtpPort := "587"                                                             // SMTP 端口
	auth := smtp.PlainAuth("", "2388287244@qq.com", "cmuusgyezivbeccj", smtpHost) // 替换为你的 QQ 邮箱和授权码

	// 发送邮件
	e.Send(smtpHost+":"+smtpPort, auth)  //log会让程序终止不能用！！！
}
