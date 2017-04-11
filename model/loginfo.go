package model

import (
	"time"
)

const (
	Log_Info = iota
	Log_Warn
	Log_Error
)

//日志信息
type LogInfo struct {
	*BaseModel  `xorm:"extends"`
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
	r, err = log.DbCommandSession("ip=? and UserAgent=? and Url=? and CreateTime>date_add(now(),interval -2 hour)", log.IP, log.UserAgent, log.Url).Count(log)
	if err != nil {
		return
	}

	if r == 0 {
		r, err = log.BaseModel.Add(log)
	} else {
		r = 0
	}

	return
}
