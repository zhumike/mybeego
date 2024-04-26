package controllers

import (
	"github.com/astaxie/beego"
)

type GoodsController struct { // 首字母要大写
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["title"] = "你好beego" // 绑定数据
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "goods.tpl"
}

func (c *GoodsController) DoEdit() {
	title := c.GetString("title")
	c.Ctx.WriteString("执行修改操作--" + title)
}
