// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"navsystem/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
	//beego.Router("/navs",&controllers.NavgationController{},"get:GetAll")
	beego.Router("/nav",&controllers.NavgationController{},"get:GetAll;put:Put;post:Post")
	beego.Router("/nav/list",&controllers.NavgationController{},"get:GetAllTreeNavs")
	beego.Router("/nav/:nid",&controllers.NavgationController{},"get:Get;delete:Delete")
	beego.Router("/nav/init",&controllers.NavgationController{},"*:InitData")

	beego.Router("/navdic",&controllers.NavDicController{},"get:GetAll;post:Post;put:Put")
	beego.Router("/navdic/:nid",&controllers.NavDicController{},"get:Get;delete:Delete")
}
