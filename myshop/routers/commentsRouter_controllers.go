package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["myshop/controllers:AddressController"] = append(beego.GlobalControllerRouter["myshop/controllers:AddressController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:AddressController"] = append(beego.GlobalControllerRouter["myshop/controllers:AddressController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:AddressController"] = append(beego.GlobalControllerRouter["myshop/controllers:AddressController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:AddressController"] = append(beego.GlobalControllerRouter["myshop/controllers:AddressController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:AddressController"] = append(beego.GlobalControllerRouter["myshop/controllers:AddressController"],
        beego.ControllerComments{
            Method: "GetByUid",
            Router: `/uid/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:CartController"] = append(beego.GlobalControllerRouter["myshop/controllers:CartController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:CartController"] = append(beego.GlobalControllerRouter["myshop/controllers:CartController"],
        beego.ControllerComments{
            Method: "Patch",
            Router: `/`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:CartController"] = append(beego.GlobalControllerRouter["myshop/controllers:CartController"],
        beego.ControllerComments{
            Method: "GetByUid",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:CartController"] = append(beego.GlobalControllerRouter["myshop/controllers:CartController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:CommentController"] = append(beego.GlobalControllerRouter["myshop/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:CommentController"] = append(beego.GlobalControllerRouter["myshop/controllers:CommentController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:CommentController"] = append(beego.GlobalControllerRouter["myshop/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:CommentController"] = append(beego.GlobalControllerRouter["myshop/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:GoodController"] = append(beego.GlobalControllerRouter["myshop/controllers:GoodController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:GoodController"] = append(beego.GlobalControllerRouter["myshop/controllers:GoodController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:GoodController"] = append(beego.GlobalControllerRouter["myshop/controllers:GoodController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:GoodController"] = append(beego.GlobalControllerRouter["myshop/controllers:GoodController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:GoodController"] = append(beego.GlobalControllerRouter["myshop/controllers:GoodController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:OrderController"] = append(beego.GlobalControllerRouter["myshop/controllers:OrderController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:OrderController"] = append(beego.GlobalControllerRouter["myshop/controllers:OrderController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:OrderController"] = append(beego.GlobalControllerRouter["myshop/controllers:OrderController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:OrderController"] = append(beego.GlobalControllerRouter["myshop/controllers:OrderController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:OrderController"] = append(beego.GlobalControllerRouter["myshop/controllers:OrderController"],
        beego.ControllerComments{
            Method: "GetOrderUid",
            Router: `/uid/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:OrderItemController"] = append(beego.GlobalControllerRouter["myshop/controllers:OrderItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:OrderItemController"] = append(beego.GlobalControllerRouter["myshop/controllers:OrderItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:OrderItemController"] = append(beego.GlobalControllerRouter["myshop/controllers:OrderItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:OrderItemController"] = append(beego.GlobalControllerRouter["myshop/controllers:OrderItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:SellerController"] = append(beego.GlobalControllerRouter["myshop/controllers:SellerController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:SellerController"] = append(beego.GlobalControllerRouter["myshop/controllers:SellerController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:SellerController"] = append(beego.GlobalControllerRouter["myshop/controllers:SellerController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:SellerController"] = append(beego.GlobalControllerRouter["myshop/controllers:SellerController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:UserController"] = append(beego.GlobalControllerRouter["myshop/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:UserController"] = append(beego.GlobalControllerRouter["myshop/controllers:UserController"],
        beego.ControllerComments{
            Method: "Patch",
            Router: `/`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:UserController"] = append(beego.GlobalControllerRouter["myshop/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:UserController"] = append(beego.GlobalControllerRouter["myshop/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myshop/controllers:UserController"] = append(beego.GlobalControllerRouter["myshop/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
