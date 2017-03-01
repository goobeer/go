package model

type Privilege struct {
	ID   int64
	Name string `xorm:varchar(20)` //权限名称
	Val  int    //权限值,1,2,4,8
	Used bool   //是否启用
}
