package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) DoRegister() {

	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println(username, password)
	c.TplName = "success.html"
}
