package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.html"
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	beego.Informational("LoginController  "+"        11111111111")
}
