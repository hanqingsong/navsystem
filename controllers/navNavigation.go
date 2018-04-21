package controllers

import (
	"navsystem/models"
	"fmt"
	"encoding/json"
)

type NavgationController struct {
	CommonController
}

//获取全部导航数据
func (nav *NavgationController) GetAll() {
	navs := models.GetAllNavs()
	resultMap:=map[string]interface{}{"list": navs};
	nav.AjaxResponse(200,"success",resultMap)
}

//获取详情
func (nav *NavgationController) Get() {
	nid := nav.GetString(":nid")
	s := nav.Input().Get("name")

	fmt.Println("参数s：",s)
	navgation := models.Get(nid)
	nav.AjaxResponse(200,"success",navgation)
}

//新增
func (this *NavgationController) Post() {
	var navNavigation models.NavNavigation
	// 获取json参数
	json.Unmarshal(this.Ctx.Input.RequestBody, &navNavigation)
	fmt.Println(navNavigation)
	//获取form参数
	//this.ParseForm(&navNavigation)
	//fmt.Println(navNavigation)
	i, _ := models.Add(&navNavigation)
	this.AjaxResponse(200,"success",i)
}

//修改
func (nav *NavgationController) Put() {
	//获取地址参数
	//nid, _ := nav.GetInt64(":nid")
	//fmt.Println(nid)
	var navNavigation models.NavNavigation

	// get获取body参数
	//s := nav.GetString("intro")
	//fmt.Println(s)

	//获取body参数
	//bytes := nav.Ctx.Input.RequestBody
	//fmt.Println(string(bytes))

	// 获取json参数
	json.Unmarshal(nav.Ctx.Input.RequestBody, &navNavigation)
	fmt.Println(navNavigation)

	//获取form参数
	//nav.ParseForm(&navNavigation)
	//navNavigation.Id=nid
	//fmt.Println(navNavigation)

	i, _ := models.UpdateNav(&navNavigation)
	nav.AjaxResponse(200,"success",i)
}


//删除
func (nav *NavgationController) Delete() {
	nid, _ := nav.GetInt64(":nid")
	fmt.Println(nid)
	navg := models.NavNavigation{Id:nid}
	i, _ := models.DeleteNav(&navg)
	nav.AjaxResponse(200,"success",i)
}