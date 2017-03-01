package model

import (
	"time"
)

type Article struct {
	ID         int
	Name       string `xorm:varchar(20)`
	Content    string `xorm:text`
	CreateTime time.Time
}
