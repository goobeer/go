package model

import (
	"time"
)

type Article struct {
	ID         int64
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
	r, err = bm.Add(ms...)
	return
}

func (art *Article) Delete(query interface{}, args ...interface{}) (r int64, err error) {
	r, err = bm.Delete(art, query, args...)
	return
}

func (art *Article) UpdateByID(incCols, omitCols string) (r int64, err error) {
	r, err = bm.UpdateByID(art.ID, art, incCols, omitCols)
	return
}

func (art *Article) Get(query interface{}, args ...interface{}) (ok bool, err error) {
	ok, err = bm.DbCommandSession(query, args...).Get(art)
	return
}

func (art *Article) GetList(pageNumber, pageSize uint, query interface{}, args ...interface{}) (res []Article, err error) {
	err = bm.DbListCommandSession(pageNumber, pageSize, query, args...).Find(&res)
	return
}
