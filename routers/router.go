package routers

import (
	"beeblog/controllers"
	"beeblog/models"
	"beeblog/utils"
	beego "github.com/beego/beego/v2/server/web"
)

var articles []models.Article

func init() {
	utils.InitTemplate()
	beego.SetStaticPath("/static", "static")

	beego.Router("/?:page", &controllers.HomeController{Articles: &articles, HasInitiated: false})
	beego.Router("/beegoblog/:id", &controllers.ShowController{Articles: &articles})
	beego.Router("/beegoblog/create", &controllers.CreateController{Articles: &articles})

	beego.Router("/beegoblog/search", &controllers.SearchController{})
	beego.Router("/beegoblog/search/:tag:string", &controllers.SearchResultController{Articles: &articles})

	beego.Router("/user/signup", &controllers.SignUpController{})
	beego.Router("/user/login", &controllers.LoginController{})
	beego.Router("/user/logout", &controllers.LogoutController{})
	beego.Router("/user/profile", &controllers.ProfileController{})
}
