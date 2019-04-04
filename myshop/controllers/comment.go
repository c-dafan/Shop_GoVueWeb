package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"myshop/models"
	"myshop/util"
	"strconv"
)

// CommentController operations for Comment
type CommentController struct {
	beego.Controller
}

// URLMapping ...
func (c *CommentController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Comment
// @Param	body		body 	models.Comment	true		"body for Comment content"
// @Success 201 {int} models.Comment
// @Failure 403 body is empty
// @router / [post]
func (c *CommentController) Post() {
	var v models.Comment
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if cid, err := models.AddComment(&v); err == nil {
			//c.Ctx.Output.SetStatus(201)
			ot, err := models.GetOrderItemById(v.Otid.Id)
			if err == nil {
				v.Id = int(cid)
				ot.Cid =  &v
				if err = models.UpdateOrderItemById(ot);err==nil{
					c.Data["json"] = util.Success("OK")
				}else{
					c.Data["json"] = util.Error(err)
				}
			}else{
				c.Data["json"] = util.Error(err)
			}
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
// @Description get Comment by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comment
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CommentController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCommentById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Comment
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Comment	true		"body for Comment content"
// @Success 200 {object} models.Comment
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CommentController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Comment{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateCommentById(&v); err == nil {
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
// @Description delete the Comment
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CommentController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteComment(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
