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
	ID         int64
	Name       string    `xorm:"char(20)"`
	Pwd        string    `xorm:"char(50)"`
	CreateTime time.Time `xorm:"created"`
	Used       bool
}

func (u *Users) Add(users ...*Users) (r int64, err error) {
	var ms []interface{}
	for _, v := range users {
		v.ChangePwd()
		ms = append(ms, v)
	}

	r, err = bm.Add(ms...)
	return
}

func (u *Users) UpdateByID(incCols, omitCols string) (r int64, err error) {
	r, err = bm.UpdateByID(u.ID, u, incCols, omitCols)
	return
}

func (u *Users) Delete(query interface{}, args ...interface{}) (r int64, err error) {
	r, err = bm.DbCommandSession(query, args...).Delete(u)
	return
}

func (u *Users) Get(query interface{}, args ...interface{}) (ok bool, err error) {
	ok, err = bm.DbCommandSession(query, args...).Get(u)
	return
}

func (u *Users) GetList(pageNumber, pageSize uint, query interface{}, args ...interface{}) (data []Users, err error) {
	err = bm.DbListCommandSession(pageNumber, pageSize, query, args...).Find(&data)
	return
}

func (u *Users) ChangePwd() {
	salt, _ := common.CreateRandom(5, 4)
	hash := md5.New()
	hash.Write([]byte(u.Pwd + salt))
	chash := hash.Sum(nil)
	u.Pwd = fmt.Sprintf("%s:%s", hex.EncodeToString(chash), salt)
}

func (user *Users) VerfyUser(name, pwd string) (has bool, err error) {
	has, err = user.Get("Name=?", name)

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
