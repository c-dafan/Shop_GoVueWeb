package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"myshop/models"
	"strconv"
)

// SellerController operations for Seller
type SellerController struct {
	beego.Controller
}

// URLMapping ...
func (c *SellerController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Seller
// @Param	body		body 	models.Seller	true		"body for Seller content"
// @Success 201 {int} models.Seller
// @Failure 403 body is empty
// @router / [post]
func (c *SellerController) Post() {
	var v models.Seller
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSeller(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Seller by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Seller
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SellerController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSellerById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Seller
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Seller	true		"body for Seller content"
// @Success 200 {object} models.Seller
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SellerController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Seller{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSellerById(&v); err == nil {
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
// @Description delete the Seller
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SellerController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSeller(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
