package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"myshop/models"
	"myshop/util"
	"strconv"
)

// OrderController operations for Order
type OrderController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrderController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Order
// @Param	body		body 	models.Order	true		"body for Order content"
// @Success 201 {int} models.Order
// @Failure 403 body is empty
// @router / [post]
func (c *OrderController) Post() {
	var v models.Order
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if id, err := models.AddOrder(&v); err == nil {
			//c.Ctx.Output.SetStatus(201)
			c.Data["json"] = util.Success(id)
		} else {
			c.Data["json"] = util.Error(err)
		}
	} else {
		c.Data["json"] = util.Error(err)
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Order by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Order
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OrderController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOrderById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @router /uid/:id [get]
func (c *OrderController) GetOrderUid() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOrderByUid(id)
	if err != nil {
		c.Data["json"] = util.Error(err)
	} else {
		var items []interface{}
		for _, vv := range v {
			vi, _ := models.GetOrderItemByOId(vv.Id)
			item := make(map[string]interface{})
			item["items"] = vi
			item["order"] = vv
			items = append(items, item)
		}
		c.Data["json"] = util.Success(items)
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Order
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Order	true		"body for Order content"
// @Success 200 {object} models.Order
// @Failure 403 :id is not int
// @router /:id [put]
func (c *OrderController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Order{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateOrderById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Order
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OrderController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOrder(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
