package mygo

import (
	"gopkg.in/gomail.v2"
)

type Mail struct {
	Enable int
	Host   string
	Port   int
	User   string
	Pass   string
	dialer *gomail.Dialer
}

type MailMessage struct {
	From    string
	To      []string
	Cc      []string
	Subject string
	Body    string
	Attach  string
	mail    *Mail
}

func NewMail(m *Mail) *Mail {
	m.dialer = gomail.NewDialer(m.Host, m.Port, m.User, m.Pass)
	return m
}

func (m *Mail) Send(mMsg *MailMessage) (err error) {
	if m.Enable == 0 {
		return
	}
	mMsg.mail = m
	msg := mMsg.NewMessage()
	err = m.dialer.DialAndSend(msg)
	if err != nil {
		return
	}
	return
}

func (mm *MailMessage) NewMessage() *gomail.Message {
	mailMsg := gomail.NewMessage()
	if mm.From == "" {
		mailMsg.SetHeader("From", mm.mail.User)
	} else {
		mailMsg.SetHeader("From", mm.From)
	}
	mailMsg.SetHeader("To", mm.To...)
	if len(mm.Cc) > 0 {
		mailMsg.SetHeader("Cc", mm.Cc...)
	}
	mailMsg.SetHeader("Subject", mm.Subject)
	mailMsg.SetBody("text/html", mm.Body)
	if mm.Attach != "" {
		mailMsg.Attach(mm.Attach)
	}
	return mailMsg
}
