package model

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"
)

//用户信息
type Users struct {
	ID         int64
	Name       string `xorm:"char(20)"`
	Pwd        string `xorm:"char(50)"`
	CreateTime time.Time
	Used       bool
}

func (user *Users) VerfyUser(name, pwd string) (has bool, err error) {
	has, err = engine.Where("Name=?", name).Get(user)
	if has {
		pwdStr := strings.Split(user.Pwd, ":")
		salt := pwdStr[1]
		hash := md5.New()
		hash.Write([]byte(pwd + salt))
		chash := hash.Sum(nil)
		has = pwdStr[0] == hex.EncodeToString(chash)
	}
	return
}
