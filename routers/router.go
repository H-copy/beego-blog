package routers

import (
	"github.com/astaxie/beego"
	"github.com/ulricqin/beego-blog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.AutoRouter(&controllers.ApiController{})

	beego.Router("/article/:ident", &controllers.ArticleController{}, "get:Read")
	beego.Router("/catalog/:ident", &controllers.CatalogController{}, "get:ListByCatalog")

	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	beego.Router("/me", &controllers.MeController{}, "get:Default")
	beego.Router("/me/catalog/add", &controllers.CatalogController{}, "get:Add;post:DoAdd")
	beego.Router("/me/catalog/edit", &controllers.CatalogController{}, "get:Edit;post:DoEdit")
	beego.Router("/me/article/add", &controllers.ArticleController{}, "get:Add;post:DoAdd")
}
