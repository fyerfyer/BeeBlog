package utils

import (
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

func RenderFlash(c *bee.Controller, tplName string) {
	if tplName != "" {
		c.TplName = tplName
	}
	flash := bee.ReadFromRequest(c)
	if flash.Data != nil {
		if success, ok := flash.Data["success"]; ok {
			c.Data["Flash"] = map[string]interface{}{"success": success}
		}
	}
}

func RenderAuthenticated(c *bee.Controller, tplName string) {
	if tplName != "" {
		c.TplName = tplName
	}
	c.Data["IsAuthenticated"] = IsAuthenticated(c)
}
