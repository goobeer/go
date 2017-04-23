package model

import (
	"crypto/md5"
	"encoding/hex"
	"fasthttpweb/common"
	"fmt"
	"strings"
	"time"
)

//用户信息
type Users struct {
	*BaseModel `xorm:"extends"`
	Name       string    `xorm:"char(20)"`
	Pwd        string    `xorm:"char(50)"`
	CreateTime time.Time `xorm:"created"`
	Used       bool
}

func (u *Users) Add(users ...*Users) (r int64, err error) {
	var ms []interface{}
	for _, v := range users {
		ms = append(ms, v)
	}

	r, err = u.BaseModel.Add(ms...)
	return
}

func (u *Users) GeneratePwd() {
	salt, _ := common.CreateRandom(5, 4)
	hash := md5.New()
	hash.Write([]byte(u.Pwd + salt))
	chash := hash.Sum(nil)
	u.Pwd = fmt.Sprintf("%s:%s", hex.EncodeToString(chash), salt)
}

func (u *Users) RepeatName() (isRepeat bool) {
	count, err := u.DbCommandSession("Name=?", u.Name).Count(u)
	if err == nil {
		isRepeat = count > 0
	}
	return
}

func (user *Users) VerfyUser(name, pwd string) (has bool, err error) {
	has, err = user.Get(user, "Name=?", name)

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
