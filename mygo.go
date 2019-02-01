package mygo

import "database/sql"

type Mygo struct {
}

const (
	VERSION = "1.0.0"
)

var (
	G_mail *Mail
	Db     *sql.DB
)

func NewMygo() *Mygo {
	mygo := &Mygo{}
	return mygo
}

func (m *Mygo) RegisterMail() {
	mailConf := &Mail{
		Enable: G_config.Smtp.Enable,
		Host:   G_config.Smtp.Host,
		Port:   G_config.Smtp.Port,
		User:   G_config.Smtp.User,
		Pass:   G_config.Smtp.Pass,
	}
	G_mail = NewMail(mailConf)
}

func (m *Mygo) RegisterDb(){
	myDB := NewMyDB()
	err := myDB.Open()
	if err != nil {
		panic(err)
	}
	Db = myDB.db
}
