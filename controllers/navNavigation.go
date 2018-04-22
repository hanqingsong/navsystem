package controllers

import (
	"navsystem/models"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"regexp"
	"net/http"
	"strings"
)

type NavgationController struct {
	CommonController
}

//获取全部导航数据
func (nav *NavgationController) GetAll() {
	navs := models.GetAllNavs()
	resultMap:=map[string]interface{}{"list": navs}
	nav.AjaxResponse(200,"success",resultMap)
}


//获取全部导航数据
func (nav *NavgationController) GetAllTreeNavs() {
	navs := models.GetAllTreeNavs()
	resultMap:=map[string]interface{}{"list": navs}
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

//初始化数据
func (nav *NavgationController) InitData(){
	url:=""
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	sHtml, _ := ioutil.ReadAll(resp.Body)
	rex2:=regexp.MustCompile(`<div class="layui-tab layui-tab-card" id="bar\d*">([\s\S]*?)</div>`)
	result2:=rex2.FindAllStringSubmatch(string(sHtml),-1)
	fmt.Println(len(result2))
	//fmt.Println(result2[0][1])
	titleRex:=regexp.MustCompile(`<li class="layui-tab-li">(.*)</li>`)
	itemRex:=regexp.MustCompile(`<a class="box-item" href="(.*)" target="_blank">(.*)</a>`)
	for _, value := range result2 {
		contentHtml:=value[1]
		//fmt.Println(key,contentHtml)
		title:=titleRex.FindAllStringSubmatch(contentHtml,-1)
		item:=itemRex.FindAllStringSubmatch(contentHtml,-1)
		//fmt.Println(title[0][1])
		navDic := addGroup(title[0][1])
		groupLid := navDic.Lid
		for _, value := range item {
			//fmt.Print(value[1],"\t")
			//fmt.Println(value[2])
			addItem(groupLid,value[2],value[1])
		}

		fmt.Println("----------")
	}


}

// 插入群组
func addGroup(groupName string) models.NavDic{
	fmt.Println(groupName)
	var navDic models.NavDic
	navDic.Name=groupName
	navDic.Code=groupName
	navDic.Type=1
	models.AddNavDic(&navDic)
	return navDic
}

// 插入地址
func addItem(groupLid string,name string,url string){
	fmt.Print(name,"\t")
	fmt.Println(url)
	var navNavigation models.NavNavigation
	navNavigation.Name=name
	navNavigation.Url=url
	if !strings.HasSuffix(url,"/"){
		url += "/"
	}
	navNavigation.FaviconUrl=url+"favicon.ico"
	navNavigation.GroupLid=groupLid
	models.Add(&navNavigation)
}