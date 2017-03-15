package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

type BaseModel struct {
	ID int64
}

var (
	engine *xorm.Engine
)

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:watersa@tcp(127.0.0.1:3306)/tdb?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	engine.SetMapper(core.SameMapper{})
}
