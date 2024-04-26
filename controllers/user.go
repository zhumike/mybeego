package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("用户中心！")

}

func (c *UserController) AddUser() {
	c.TplName = "user.html"
}

// 处理post请求时，对应的方法
func (c *UserController) DoAddUser() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id 必须是 int 类型")
		return
	}
	fmt.Printf("%v---%T", id, id)

	username := c.GetString("username")
	password := c.GetString("password")
	hobby := c.GetStrings("hobby")
	fmt.Printf("值: %v---类型: %T", hobby, hobby)
	c.Ctx.WriteString("用户中心--" + strconv.Itoa(id) + username + password)

}

func (c *UserController) EditUser() {
	c.TplName = "userEdit.html"
}

type User struct {
	Username string   `form:"username"` // html 中 name 是小写，所以这里需要使用 `` 来定义 tag 标签
	Password string   `form:"password"`
	Hobby    []string `form:"hobby"`
}

func (c *UserController) DoEditUser() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		c.Ctx.WriteString("post提交失败")
		logs.Debug("请查看入参提交")
		return
	}
	fmt.Printf("%#v", u)
	c.Ctx.WriteString("解析 post 数据成功")

}

// 在 beego 中，如果我们需要返回一个 json 数据，需要把数据放在结构体中
func (c *UserController) GetUser() {
	u := User{
		Username: "mike",
		Password: "123456",
		Hobby:    []string{"", "2"},
	}
	// 返回一个 json 数据
	c.Data["json"] = u
	c.ServeJSON()
}
