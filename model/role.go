package model

type Role struct {
	*BaseModel `xorm:"extends"`
	Name       string `xorm:"varchar(20)"` //角色名称
	PrivID     int64  //权限组ID
	Used       bool
}
