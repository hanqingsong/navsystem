package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
)

func init() {
	fmt.Println("base.go")
	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	port, err := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")
	if nil != err {
		port = 3306
	}
	orm.Debug=true

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname)
	fmt.Println(dataSource)
	orm.RegisterDataBase("default","mysql", dataSource)
}