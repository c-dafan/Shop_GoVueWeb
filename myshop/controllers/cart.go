package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"myshop/models"
	"myshop/util"
	"strconv"
)

type CartController struct {
	beego.Controller
}

func (c *CartController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetByUid", c.GetByUid)
	//c.Mapping("GetAll", c.GetAll)
	c.Mapping("PATCH", c.Patch)
	c.Mapping("Delete", c.Delete)
}

// @router / [post]
func (this *CartController) Post() {
	var cart models.Cart
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cart); err == nil {
		if v, err := models.FindCardByUidAndGid(&cart);err == nil{
			v.Num += cart.Num
			if err = models.UpdateCartById(v);err == nil{
				this.Data["json"] = util.Success("")
				this.ServeJSON()
				return
			}else {
				this.Data["json"] = util.Error(err)
			}
		}
		if num, err := models.AddCart(&cart); err == nil {
			//if cart2,err := models.GetCartById(int(num));err == nil{
			//	this.Data["json"] = util.Success(cart2)
			//}else {
			//	this.Data["json"] = util.Error(err)
			//}
			this.Data["json"] = util.Success(num)
		} else {
			this.Data["json"] = util.Error(err)
		}
	} else {
		this.Data["json"] = util.Error(err)
	}
	this.ServeJSON()
}

// @router /:id [get]
func (this *CartController) GetByUid() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if v, err := models.GetCartByUid(id); err == nil {
		this.Data["json"] = util.Success(v)
	}
	this.ServeJSON()
}

// @router /:id [delete]
func (this *CartController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCart(id); err == nil {
		this.Data["json"] = util.Success("")
	} else {
		this.Data["json"] = util.Success(err)
	}
	this.ServeJSON()
}

// @router / [patch]
func (this *CartController) Patch() {
	var cart models.Cart
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cart); err == nil {
		fmt.Println(cart)
		if err = models.UpdateCartById(&cart); err == nil{
			this.Data["json"] = util.Success("")
		}else{
			this.Data["json"] = util.Error(err)
		}
	}else{
		this.Data["json"] = util.Error(err)
	}
	this.ServeJSON()
}