package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"myshop/models"
	"myshop/util"
	"strconv"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Login", c.Login)
	c.Mapping("Patch", c.Patch)
}

func (this *UserController) Login() {
	var v models.User
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &v); err == nil {
		if u, err := models.GetUserById(v.Id); err == nil {
			if u == nil {
				this.Data["json"] = util.Error(errors.New("用户不存在"))
			} else if v.Password == u.Password {
				res := make(map[string]interface{})
				res["Id"] = v.Id
				id, _ := uuid.NewV1()
				res["Token"] = fmt.Sprint(id)
				res["Name"] = u.Username
				this.Data["json"] = util.Success(res)
			} else {
				this.Data["json"] = util.Error(errors.New("密码不正确"))
			}
		} else {
			this.Data["json"] = util.Error(errors.New("用户不存在"))
		}
	} else {
		this.Data["json"] = util.Error(err)
	}
	this.ServeJSON()
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	var v models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println(v)
		if _, err := models.AddUser(&v); err == nil {
			fmt.Println(v)
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
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.User{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUserById(&v); err == nil {
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
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUser(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @router / [patch]
func (this *UserController) Patch() {
	v := &models.User{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, v); err == nil {
		if user, err := models.GetUserById(v.Id); err == nil {
			user.Aid = v.Aid
			if err = models.UpdateUserById(user); err == nil {
				this.Data["json"] = util.Success("Ok")
			} else {
				this.Data["json"] = util.Error(err)
			}
		} else {
			this.Data["json"] = util.Error(err)
		}
	} else {
		this.Data["json"] = util.Error(err)
	}
	this.ServeJSON()
}
