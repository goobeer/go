package model

const (
	Read    = 1
	Write   = 2
	Execute = 4
	All     = Read | Write | Execute
	None    = ^All

	//权限类型
	PrivilegeClass  = 1
	PrivilegeEntity = 2
)

type Privilege struct {
	ID   int64
	Name string `xorm:varchar(20)` //权限名称
	Val  int    //权限值,1,2,4,8
	Used bool   //是否启用
}
