package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type CommonController struct {
	beego.Controller
}

func success(data interface{}) map[string]interface{}{
	fmt.Println(data)
	return map[string]interface{}{"code": 200, "message": "请求成功", "data":data}
}

//ajax 返回
func (this *CommonController) AjaxResponse(resultCode int64, resultString string, data interface{}) {
	response := struct {
		StatusCode       int64
		StatusText string
		Data interface{}
	}{
		StatusCode:       resultCode,
		StatusText: resultString,
		Data: data,
	}

	this.Data["json"] = response
	this.ServeJSON()
}