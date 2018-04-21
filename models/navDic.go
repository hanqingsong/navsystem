package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"navsystem/utils"
)

func init() {
	fmt.Println("navdic.go")
	//orm.Debug=true
	//orm.RegisterDataBase("default","mysql","xiaozhi:xiaozhi123@tcp(193.112.46.144:3306)/xiaozhi?charset=utf8")
	orm.RegisterModel(new(NavDic))
}
// 导航数据
type NavDic struct {
	Id int64
	Lid string
	Code string
	Name string
	Value string
	Type int64
	Comments string
	Pcode string
	Status int64
	OrderNum int64
	CreateTime time.Time
	UpdateTime time.Time
}

func GetAllNavDics() []NavDic {
	o :=orm.NewOrm()
	raw := o.Raw("select * from nav_dic where status=1")
	var navs []NavDic
	rows, error := raw.QueryRows(&navs)
	if(error!= nil) {
		panic(error)
	}
	fmt.Println(rows)
	return navs
}


func GetNavDic(id string) NavDic {
	o:=orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("nav_dic").Where("id=?").Limit(1)
	raw := o.Raw(qb.String(),id)
	var nav NavDic
	row := raw.QueryRow(&nav)
	fmt.Println(row)
	return nav
}


func AddNavDic(navgation *NavDic) (int64,error){
	o:=orm.NewOrm()
	navgation.Status=1
	navgation.Lid=utils.GenerateUUID()
	navgation.CreateTime=time.Now()
	navgation.UpdateTime=time.Now()
	id, error := o.Insert(navgation)
	return id,error
}


func UpdateNavDic(navgation *NavDic) (int64,error){
	o:=orm.NewOrm()
	id, error := o.Update(navgation)
	return id,error
}


func DeleteNavDic(navgation *NavDic) (int64,error){
	navDic2 := NavDic{Id:navgation.Id,Status:0,UpdateTime:time.Now()}
	o:=orm.NewOrm()
	id, error := o.Update(&navDic2,"Status","UpdateTime")
	return id,error
}