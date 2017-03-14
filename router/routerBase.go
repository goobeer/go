package router

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"fasthttpweb/common"
	"fasthttpweb/model"

	"github.com/buaazp/fasthttprouter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var (
	R *fasthttprouter.Router = fasthttprouter.New()
)

func test() {

	var err error
	//	var engine *xorm.Engine
	engine, err := xorm.NewEngine("mysql", "dbUname:dbupwd@tcp(127.0.0.1:3306)/tdb?charset=utf8")
	engine.SetMapper(core.SameMapper{})

	common.PanicError(err)

	user := &model.Users{Name: "root", Used: true, CreateTime: time.Now()}

	engine.Sync2(user)

	salt, _ := common.CreateRandom(5, 4)
	hash := md5.New()
	hash.Write([]byte("123456" + salt))
	chash := hash.Sum(nil)

	user.Pwd = fmt.Sprintf("%s:%s", hex.EncodeToString(chash), salt)
	//	engine.Insert(user)
	var rootUser *model.Users = &model.Users{}
	//	b, qErr := engine.Where("Name=?", "root").Desc("id").Get(rootUser)
	//	panicError(qErr)
	//	fmt.Println(b)
	//	if b {
	//		fmt.Println(rootUser)
	//	}

	engine.ID(1).Get(rootUser)
	fmt.Println(rootUser)
	engine.ShowSQL(true)
	engine.ID(1).MustCols("Name", "Pwd").Update(user)
}
