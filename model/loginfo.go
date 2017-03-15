package model

import (
	"time"
)

//日志信息
type LogInfo struct {
	ID           int64
	IP           string    `xorm:"char(15)"`
	UserAgent    string    `xorm:"varchar(max)"`
	ExceptionMsg string    `xorm:"text"`         //异常信息
	Logger       string    `xorm:char(20)`       //日志源
	LogCatelog   string    `xorm:"varchar(50)"`  //日志类别
	Msg          string    `xorm:"text"`         //日志信息
	Url          string    `xorm:"varchar(max)"` //请求地址
	CreateTime   time.Time `xorm:"created"`
}

func (log *LogInfo) Add() (r int64, err error) {
	var existLog *LogInfo
	engine.Where("ip=? and UserAgent=? and Url=? and CreateTime", log.IP, log.UserAgent, log.Url)
	r, err = engine.InsertOne(log)
	return
}
