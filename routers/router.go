// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"new-beego-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
			beego.NSRouter("", &controllers.ObjectController{}, "get:GetAll"),
			beego.NSRouter("/:objectId", &controllers.ObjectController{}, "get:Get"),
			beego.NSRouter("", &controllers.ObjectController{}, "post:Post"),
			beego.NSRouter("/:objectId", &controllers.ObjectController{}, "put:Put"),
			beego.NSRouter("/:objectId", &controllers.ObjectController{}, "delete:Delete"),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
			beego.NSRouter("", &controllers.UserController{}, "get:GetAll"),
			beego.NSRouter("/:uid", &controllers.UserController{}, "get:Get"),
			beego.NSRouter("", &controllers.UserController{}, "post:Post"),
			beego.NSRouter("/:uid", &controllers.UserController{}, "put:Put"),
			beego.NSRouter("/:uid", &controllers.UserController{}, "delete:Delete"),
		),
	)
	beego.AddNamespace(ns)
}
