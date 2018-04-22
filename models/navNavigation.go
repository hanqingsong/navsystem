package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"navsystem/utils"
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
	FaviconUrl string
	Intro string `form:"intro"`
	Comments string
	Status int64
	OrderNum int64
	ClickTimes int64
	CreateTime time.Time
	UpdateTime time.Time
}


type NavNavigationVO struct {
	NavNavigation
	GroupName string
}

type NavNavigationTreeVO struct {
	NavDic
	Children []NavNavigation
}



func GetAllNavs() []NavNavigationVO {
	o :=orm.NewOrm()
	raw := o.Raw("select na.*,nd.`name` group_name from nav_navigation na " +
		" join nav_dic nd on na.`group_lid`=nd.lid where na.`status`=1 and nd.status=1 order by nd.order_num ")
	var navs []NavNavigationVO
	rows, error := raw.QueryRows(&navs)
	if(error!= nil) {
		panic(error)
	}
	fmt.Println(rows)
	return navs
}

func GetAllTreeNavs() []NavNavigationTreeVO {
	o:=orm.NewOrm()
	raw := o.Raw("select na.`group_lid` lid,nd.`name` from `nav_navigation` na  " +
		" join nav_dic nd on na.`group_lid`=nd.lid and nd.`status`=1 where na.status=1 group by na.`group_lid` order by nd.order_num")
	var navTreeVOs []NavNavigationTreeVO
	rows, error := raw.QueryRows(&navTreeVOs)
	if error!= nil {
		panic(error)
	}
	if rows> 0 {
		for index, navTreeVO := range navTreeVOs {
			groupLid:= navTreeVO.Lid
			var navNavigation []NavNavigation
			raw := o.Raw("select lid,name,url,favicon_url from `nav_navigation` na where na.`group_lid`=? order by order_num",groupLid)
			queryRows, error := raw.QueryRows(&navNavigation)
			if error!=nil {
				panic(error)
			}
			fmt.Println(queryRows)
			if queryRows>0 {
				navTreeVOs[index].Children=navNavigation
			}
		}
	}

	return navTreeVOs
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
	navgation.Status=1
	navgation.Lid=utils.GenerateUUID()
	navgation.CreateTime=time.Now()
	navgation.UpdateTime=time.Now()
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