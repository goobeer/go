package model

const (
	//权限类型
	Read    = 1
	Write   = 2
	Execute = 4
	All     = Read | Write | Execute
	None    = ^All

	//权限类别
	PrivilegeSystem = 1
	PrivilegeClass  = 2
	PrivilegeEntity = 4
)

type Privilege struct {
	*BaseModel    `xorm:"extends"`
	Name          string `xorm:"varchar(20)"`
	Val           int    //权限值,1,2,4,8
	PrivilegeType int
	Used          bool
}
