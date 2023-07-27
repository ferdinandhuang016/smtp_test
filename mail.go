package main

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"log"
)

type mail struct {
	senderAddr   string   // 发件人地址
	senderName   string   // 发件人名称
	receiverAddr []string // 收件人地址，可以有多个收件人
	subject      string   // 邮件主题
	text         string   // 正文
	host         string   // 邮件服务器地址
	port         int      // 邮件服务器端口号
	username     string   // 用户名
	password     string   // 密码或授权码
}

func main() {
	m := mail{
		senderAddr:   "123@outlook.com",
		senderName:   "Yahaha",
		receiverAddr: []string{"123@hotmail.com"},
		subject:      "subject",
		text:         "test",
		host:         "outlook.office365.com",
		port:         993,
		username:     "123@outlook.com",
		password:     "123",
	}
	SendMail(&m)
}

func SendMail(s *mail) {
	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(s.senderAddr, s.senderName)}, // 发件人邮箱，发件人名称
		"To":      s.receiverAddr,                                // 多个收件人
		"Subject": {s.subject},                                   // 邮件主题
	})
	m.SetBody("text/plain", s.text)
	d := gomail.NewDialer(s.host, s.port, s.username, s.password) // 发送邮件服务器、端口号、发件人账号、发件人密码
	d.SSL = true
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Println("send mail err:", err)
	}
}
