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
	var existLog []LogInfo
	err = engine.Sync2(log)

	err = engine.Where("ip=? and UserAgent=? and Url=? and CreateTime>date_add(now(),interval -2 hour)", log.IP, log.UserAgent, log.Url).Find(&existLog)
	if err != nil {
		return
	}
	if len(existLog) == 0 {
		r, err = engine.InsertOne(log)
	}
	return
}

func (log *LogInfo) Get() (ok bool, err error) {
	ok, err = engine.ID(log.ID).Get(log)
	return
}

func (log *LogInfo) GetList(pageNumber, pageSize int) (res []LogInfo, err error) {
	err = engine.Asc("ID").Limit(pageSize, pageNumber*pageSize).Find(&res)
	return
}

func (log *LogInfo) Update() (r int64, err error) {
	r, err = engine.ID(log.ID).Update(log)
	return
}

func (log *LogInfo) Delete() (r int64, err error) {
	r, err = engine.ID(log.ID).Delete(log)
	return
}
