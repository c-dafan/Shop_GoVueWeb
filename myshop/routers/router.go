// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"myshop/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/", &controllers.MainController{}),
		beego.NSRouter("/login",&controllers.UserController{},"post:Login"),
		beego.NSNamespace("/address",
			beego.NSInclude(
				&controllers.AddressController{},
			),
		),
		beego.NSNamespace("/cart",
			beego.NSInclude(
				&controllers.CartController{},
			),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),

		beego.NSNamespace("/good",
			beego.NSInclude(
				&controllers.GoodController{},
			),
		),

		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),

		beego.NSNamespace("/order_item",
			beego.NSInclude(
				&controllers.OrderItemController{},
			),
		),

		beego.NSNamespace("/seller",
			beego.NSInclude(
				&controllers.SellerController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
