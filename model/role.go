package model

type Role struct {
	ID     int64
	Name   string `xorm:"varchar(20)"` //角色名称
	PriVal int64  //权限值
	Used   bool   //是否启用
}
