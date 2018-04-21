package controllers

import (
	"navsystem/models"
	"fmt"
	"encoding/json"
	"time"
)

type NavDicController struct {
	CommonController
}

//获取全部数据
func (this *NavDicController) GetAll() {
	navDics := models.GetAllNavDics()
	resultMap:=map[string]interface{}{"list": navDics};
	this.AjaxResponse(200,"success",resultMap)
}

//获取详情
func (this *NavDicController) Get() {
	nid := this.GetString(":nid")
	s := this.Input().Get("name")

	fmt.Println("参数s：",s)
	navDic := models.GetNavDic(nid)
	this.AjaxResponse(200,"success",navDic)
}

//新增
func (this *NavDicController) Post() {
	var navDic models.NavDic
	fmt.Println(this.Input())


	// 获取json参数
	json.Unmarshal(this.Ctx.Input.RequestBody, &navDic)
	fmt.Println(navDic)

	//获取form参数
	//this.ParseForm(&navDic)
	//fmt.Println(navDic)
	i, _ := models.AddNavDic(&navDic)
	this.AjaxResponse(200,"success",i)
}

//修改
func (this *NavDicController) Put() {
	//获取地址参数
	nid, _ := this.GetInt64(":nid")
	fmt.Println(nid)
	var navDic models.NavDic

	// get获取body参数
	s := this.GetString("intro")
	fmt.Println(s)

	//获取body参数
	bytes := this.Ctx.Input.RequestBody
	fmt.Println(string(bytes))

	// 获取json参数
	json.Unmarshal(this.Ctx.Input.RequestBody, &navDic)
	fmt.Println(navDic)

	//获取form参数
	this.ParseForm(&navDic)
	//navDic.Id=nid
	fmt.Println(navDic)

	i, _ := models.UpdateNavDic(&navDic)
	this.AjaxResponse(200,"success",i)
}


//删除
func (this *NavDicController) Delete() {
	nid, _ := this.GetInt64(":nid")
	fmt.Println(nid)
	navDicg := models.NavDic{Id:nid,Status:0,UpdateTime:time.Now()}
	fmt.Println(time.Time{})
	fmt.Println(navDicg)
	i, _ := models.DeleteNavDic(&navDicg)
	this.AjaxResponse(200,"success",i)
}