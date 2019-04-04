package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"myshop/models"
	"myshop/util"
	"strconv"
)

// AddressController operations for Address
type AddressController struct {
	beego.Controller
}

// URLMapping ...
func (c *AddressController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetByUid", c.GetByUid)
}

// Post ...
// @Title Post
// @Description create Address
// @Param	body		body 	models.Address	true		"body for Address content"
// @Success 201 {int} models.Address
// @Failure 403 body is empty
// @router / [post]
func (c *AddressController) Post() {
	var v models.Address
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddAddress(&v); err == nil {
			//c.Ctx.Output.SetStatus(201)
			c.Data["json"] = util.Success(v)
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
// @Description get Address by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Address
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AddressController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAddressById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @router /uid/:id [get]
func (c *AddressController) GetByUid() {
	idStr := c.Ctx.Input.Param(":id")
	v := make(map[string]string)
	v["Uid"] = idStr
	vv, err := models.GetAllAddress(v,nil,nil,nil,0,10)
	if err != nil {
		c.Data["json"] = util.Error(err)
	} else {
		c.Data["json"] = util.Success(vv)
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Address
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Address	true		"body for Address content"
// @Success 200 {object} models.Address
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AddressController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Address{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateAddressById(&v); err == nil {
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
// @Description delete the Address
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AddressController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAddress(id); err == nil {
		c.Data["json"] = util.Success("Ok")
	} else {
		c.Data["json"] = util.Error(err)
	}
	c.ServeJSON()
}
