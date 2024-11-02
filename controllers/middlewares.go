package controllers

import (
	"beeblog/utils"
	bee "github.com/beego/beego/v2/server/web"
	"net/http"
)

type BaseController struct {
	bee.Controller
}

// Prepare runs before each controller action
func (c *BaseController) Prepare() {
	// middleware 1
	utils.RenderAuthenticated(&c.Controller, "base.layout.tpl")

	// middleware 2
	url := c.Ctx.Request.URL.Path
	if url == "/beegoblog/create" || url == "/beegoblog/"+c.Ctx.Input.Param(":id") || url == "/user/logout" {
		if !utils.IsAuthenticated(&c.Controller) {
			c.Redirect("/user/login", http.StatusSeeOther)
		}
	}
}
