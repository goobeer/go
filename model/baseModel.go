package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

type BaseModel struct {
}

var (
	engine *xorm.Engine
	bm     *BaseModel
)

func init() {
	var err error
	engine, err = xorm.NewEngine(DbDriverName, DataSrcName)
	if err != nil {
		panic(err)
	}
	engine.SetMapper(core.SameMapper{})
	bm = &BaseModel{}
}

func (bm *BaseModel) DbCommandSession(query interface{}, args ...interface{}) *xorm.Session {
	return engine.Where(query, args)
}

func (bm *BaseModel) DbListCommandSession(pageNumber, pageSize int, query interface{}, args ...interface{}) *xorm.Session {
	return engine.Asc("ID").Limit(pageSize, (pageNumber-1)*pageSize).Where(query, args)
}
