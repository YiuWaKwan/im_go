package service

import (
	"errors"
	"fmt"
	"hello/model"
	"log"
	"xorm.io/xorm"
)

var DbEngin *xorm.Engine

func init() {
	drivename := "mysql"
	DsName := "root:123456@@(111.230.244.177:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, DsName)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}
	//是否显示SQL语句
	DbEngin.ShowSQL(false)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(2)

	//自动User
	DbEngin.Sync2(new(model.User),
		new(model.Contact),
		new(model.Community))
	//DbEngin = dbengin
	fmt.Println("init data base ok")
}
