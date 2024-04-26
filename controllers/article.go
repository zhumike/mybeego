package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ArticleController struct { // 首字母要大写
	beego.Controller
}

func (c *ArticleController) Get() {
	c.Ctx.WriteString("新闻列表") // 直接给页面返回数据
}

func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("增加新闻")
}

func (c *ArticleController) EditArticle() {
	//id := c.GetString("id")
	//fmt.Printf("值：%v 类型：%T",id,id)
	//beego.Info(id)
	//c.Ctx.WriteString("修改新闻")

	/*
		传参做一层判断
	*/
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		c.Ctx.WriteString("传参错误，请检查！")
		logs.Info("end!")
		return
	}
	fmt.Printf("值：%v 类型：%T", id, id)
	c.Ctx.WriteString("修改新闻")

}
