package controllers

import (
	"github.com/astaxie/beego"
)

type CmsController struct { // 首字母要大写
	beego.Controller
}

func (c *CmsController) Get() {
	cmsId := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("cms----:" + cmsId) // 直接给页面返回数据
}
