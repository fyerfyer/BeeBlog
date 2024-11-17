package utils

import (
	"beeblog/models"

	bee "github.com/beego/beego/v2/server/web"
)

func SetFlash(c *bee.Controller, typ string, msg string) {
	flash := bee.NewFlash()
	if typ == "success" {
		flash.Success(msg)
	}
	// if other condition
	flash.Store(c)
}

func RenderFlash(c *bee.Controller, tplName string) bool {
	if tplName == "" {
		return false
	}

	c.TplName = tplName
	flash := bee.ReadFromRequest(c)

	if flash.Data == nil {
		return false
	}

	if success, ok := flash.Data["success"]; ok {
		c.Data["Flash"] = map[string]interface{}{"success": success}
	}

	return true
}

func RenderAuthenticated(c *bee.Controller, tplName string) {
	if tplName != "" {
		c.TplName = tplName
	}
	c.Data["IsAuthenticated"] = IsAuthenticated(c)
}

func SetTagString(article *models.Article) {
	switch len(article.Tags) {
	case 0:
		article.TagString = ""
	case 1:
		article.TagString = article.Tags[0].Name
	default:
		article.TagString = article.Tags[0].Name + "," + article.Tags[1].Name
	}
}
