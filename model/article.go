package model

import (
	"time"
)

type Article struct {
	*BaseModel `xorm:"extends"`
	Title      string `xorm:"char(20)"`
	Content    string `xorm:"text"`
	ArtType    int
	CreateTime time.Time `xorm:"created"`
	Used       bool
}
