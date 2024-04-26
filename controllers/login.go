package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) DoLogin() {

	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println(username, password)
	c.Redirect("/", 302)
}
