package routers

import (
	"github.com/astaxie/beego"
	"zhuhello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/goods", &controllers.GoodsController{}) // 结构体地址
	beego.Router("goods/edit", &controllers.GoodsController{}, "put:DoEdit")

	beego.Router("/article", &controllers.ArticleController{})                         // 访问 Get 方法
	beego.Router("/article/add", &controllers.ArticleController{}, "get:AddArticle")   // 访问自定义方法
	beego.Router("/article/edit", &controllers.ArticleController{}, "get:EditArticle") // 结构体地址

	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "get:AddUser")
	beego.Router("/user/doAdd", &controllers.UserController{}, "POST:DoAddUser")

	beego.Router("/user/edit", &controllers.UserController{}, "get:EditUser")
	beego.Router("/user/doEdit", &controllers.UserController{}, "post:DoEditUser")
	beego.Router("/user/getUser", &controllers.UserController{}, "get:GetUser")

	//动态路由--伪静态
	beego.Router("/cms_:id([0-9]+).html", &controllers.CmsController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/doLogin", &controllers.LoginController{}, "post:DoLogin")

	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/doRegister", &controllers.RegisterController{}, "post:DoRegister")
}
