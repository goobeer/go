package model

import (
	"time"
)

type Article struct {
	*BaseModel `xorm:"extends"`
	Title      string    `xorm:"char(20)"`
	Content    string    `xorm:"text"`
	CreateTime time.Time `xorm:"created"`
	Used       bool
}

func (art *Article) Add(arts ...*Article) (r int64, err error) {
	var ms []interface{}
	for _, v := range arts {
		ms = append(ms, v)
	}

	r, err = art.BaseModel.Add(ms...)
	return
}
