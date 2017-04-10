package model

import (
	"time"
)

//日志信息
type LogInfo struct {
	ID          int64
	IP          string    `xorm:"char(15)"`
	UserAgent   string    `xorm:"varchar(200)"`
	ExpMsg      string    `xorm:"text"`
	LoggerLevel string    `xorm:"char(20)"`
	LogCatelog  string    `xorm:"varchar(50)"`
	Msg         string    `xorm:"text"`
	Url         string    `xorm:"varchar(200)"`
	CreateTime  time.Time `xorm:"created"`
}

func (log *LogInfo) Add() (r int64, err error) {
	r, err = bm.DbCommandSession("ip=? and UserAgent=? and Url=? and CreateTime>date_add(now(),interval -2 hour)", log.IP, log.UserAgent, log.Url).Count(log)
	if err != nil {
		return
	}

	if r == 0 {
		r, err = bm.Add(log)
	} else {
		r = 0
	}

	return
}

func (log *LogInfo) UpdateByID(incCols, omitCols string) (r int64, err error) {
	r, err = bm.UpdateByID(log.ID, log, incCols, omitCols)
	return
}

func (log *LogInfo) Delete(query interface{}, args ...interface{}) (r int64, err error) {
	r, err = bm.DbCommandSession(query, args...).Delete(log)
	return
}

func (log *LogInfo) Get(query interface{}, args ...interface{}) (ok bool, err error) {
	ok, err = bm.DbCommandSession(query, args...).Get(log)
	return
}

func (log *LogInfo) GetList(pageNumber, pageSize uint, query interface{}, args ...interface{}) (res []LogInfo, err error) {
	err = bm.DbListCommandSession(pageNumber, pageSize, query, args...).Find(&res)
	return
}
