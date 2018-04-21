package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

func init() {
	fmt.Println("NavNavigation.go")
	//orm.Debug=true
	//orm.RegisterDataBase("default","mysql","xiaozhi:xiaozhi123@tcp(193.112.46.144:3306)/xiaozhi?charset=utf8")
	orm.RegisterModel(new(NavNavigation))
}
// 导航数据
type NavNavigation struct {
	Id int64
	Lid string
	Name string
	Url string
	GroupLid string
	GroupName string
	FaviconUrl string
	Intro string `form:"intro"`
	Comments string
	Status int64
	OrderNum int64
	ClickTimes int64
	CreateTime time.Time
	UpdateTime time.Time
}

func GetAllNavs() []NavNavigation {
	o :=orm.NewOrm()
	raw := o.Raw("select * from nav_navigation")
	var navs []NavNavigation
	rows, error := raw.QueryRows(&navs)
	if(error!= nil) {
		panic(error)
	}
	fmt.Println(rows)
	return navs
}


func Get(id string) NavNavigation {
	o:=orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("nav_navigation").Where("id=?").Limit(1)
	raw := o.Raw(qb.String(),id)
	var nav NavNavigation
	row := raw.QueryRow(&nav)
	fmt.Println(row)
	return nav
}


func Add(navgation *NavNavigation) (int64,error){
	o:=orm.NewOrm()
	id, error := o.Insert(navgation)
	return id,error
}


func UpdateNav(navgation *NavNavigation) (int64,error){
	o:=orm.NewOrm()
	id, error := o.Update(navgation)
	return id,error
}


func DeleteNav(navgation *NavNavigation) (int64,error){
	//id, error := o.Delete(navgation)
	o:=orm.NewOrm()
	id, error := o.Update(&navgation,"Status","UpdateTime")
	return id,error
}