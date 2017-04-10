package model

type Role struct {
	ID    int64
	Name  string `xorm:"varchar(20)"` //角色名称
	PriID int64  //权限组ID
	Used  bool   //是否启用
}
