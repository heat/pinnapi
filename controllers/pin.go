package controllers

import (
	"encoding/json"
	"pinnapi/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about object
type PinController struct {
	beego.Controller
}

type JsonApi struct {
	Code  int
	Error error
	Pin   models.Pin `json:"data"`
}

// @router / [post]
func (u *PinController) Post() {
	var pin models.Pin
	json.Unmarshal(u.Ctx.Input.RequestBody, &pin)
	next := models.Next()

	pin.ID = strconv.Itoa(next)

	pin.Codigo = pin.ID

	code, err := pin.Insert()

	j := JsonApi{
		Code:  code,
		Error: err,
		Pin:   pin,
	}
	u.Data["json"] = &j
	u.ServeJSON()
}

// @router /:id
func (c *PinController) Get() {

	pincCode := c.Ctx.Input.Param(":id")

	pin := models.Pin{}

	if code, err := pin.FindById(pincCode); err != nil {
		beego.Error("FindUserById:", err)

		j_ := JsonApi{
			Code:  code,
			Error: err,
		}
		c.Data["json"] = &j_
		c.ServeJSON()
		return
	}

	c.Data["json"] = &pin

	c.ServeJSON()
}
