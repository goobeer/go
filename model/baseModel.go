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
	//	bm     *BaseModel
	engine *xorm.Engine
)

func init() {
	var err error
	engine, err = NewDBEngine()
	if err != nil {
		panic(err)
	}

	engine.Sync2(new(LogInfo))
	engine.Sync2(new(Users))
	engine.Sync2(new(Article))
	engine.Sync2(new(Privilege))
	engine.Sync2(new(ArticleType))

	user := &Users{BaseModel: &BaseModel{}}
	user.Name = "root"

	if !user.RepeatName() {
		user.Pwd = "123456"
		user.Used = true
		user.GeneratePwd()
		user.Add(user)
	}
}

func NewDBEngine() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(DbDriverName, DataSrcName)
	if err == nil {
		engine.SetMapper(core.SameMapper{})
	}
	return engine, err
}

func (bm *BaseModel) Add(m ...interface{}) (r int64, err error) {
	r, err = engine.Insert(m...)
	return
}

func (bm *BaseModel) Add1(m ...interface{}) (r int64, err error) {
	r, err = engine.Insert(m...)
	return
}

func (bm *BaseModel) Delete(bean, query interface{}, args ...interface{}) (r int64, err error) {
	r, err = bm.DbCommandSession(query, args...).Delete(bean)
	return
}

func (bm *BaseModel) UpdateByID(bean interface{}, incCols, omitCols string) (r int64, err error) {
	sess := engine.ID(bm.ID)
	if len(incCols) == 0 && len(omitCols) == 0 {
		sess = sess.AllCols()
	} else if len(incCols) > 0 {
		sess = sess.Cols(incCols)
	} else {
		sess = sess.Omit(omitCols)
	}
	r, err = sess.Update(bean)
	return
}

func (bm *BaseModel) Get(bean, query interface{}, args ...interface{}) (ok bool, err error) {
	ok, err = bm.DbCommandSession(query, args...).Get(bean)
	return
}

func (bm *BaseModel) GetWithCols(bean interface{}, incCols, omitCols string, query interface{}, args ...interface{}) (ok bool, err error) {
	var session *xorm.Session
	if len(incCols) > 0 {
		session = bm.DbCommandSession(query, args...).Cols(incCols)
	} else if len(omitCols) > 0 {
		session = bm.DbCommandSession(query, args...).Omit(omitCols)
	} else {
		session = bm.DbCommandSession(query, args...)
	}
	ok, err = session.Get(bean)
	return
}

func (bm *BaseModel) GetList(rowSlicePtr interface{}, pageNumber, pageSize uint, query interface{}, args ...interface{}) (err error) {
	err = bm.DbListCommandSession(pageNumber, pageSize, query, args...).Find(rowSlicePtr)
	return
}

func (bm *BaseModel) GetListWithCols(rowSlicePtr interface{}, incCols, omitCols string, pageNumber, pageSize uint, query interface{}, args ...interface{}) (err error) {
	var session *xorm.Session
	if len(incCols) > 0 {
		session = bm.DbListCommandSession(pageNumber, pageSize, query, args...).Cols(incCols)
	} else if len(omitCols) > 0 {
		session = bm.DbListCommandSession(pageNumber, pageSize, query, args...).Omit(omitCols)
	} else {
		session = bm.DbListCommandSession(pageNumber, pageSize, query, args...)
	}
	err = session.Find(rowSlicePtr)
	return
}

func (bm *BaseModel) DbCommandSession(query interface{}, args ...interface{}) (sess *xorm.Session) {
	sess = engine.Where(query, args...)
	return
}

func (bm *BaseModel) DbListCommandSession(pageNumber, pageSize uint, query interface{}, args ...interface{}) (sess *xorm.Session) {
	sess = engine.Where(query, args...).Limit(int(pageSize), int(pageNumber*pageSize))
	return
}
