package controllers

import (
	"strconv"
	"encoding/json"
	"pinnapi/models"
	"github.com/astaxie/beego"
)

type pinStruct struct {
  PinCode string `json:"pin_code"`
}

// Operations about object
type PinController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *PinController) Post() {
	var pin models.Pin
	json.Unmarshal(u.Ctx.Input.RequestBody, &pin)
	u.Data["json"] = pin
	u.ServeJSON()
}
// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *PinController) Get() {
	
	next := models.Next()

	palpites := []*models.Palpite{
		&models.Palpite{ Evento: "asdasd", Odd: "asdasdasd" },
	}
	
	pin := models.Pin{
		Codigo: strconv.Itoa(next),
		Cliente: "Teste",
		Valor: models.Dinheiro(12.23),
		Palpites: palpites,
	}
	
	o.Data["json"] = pin
	
	o.ServeJSON()
}