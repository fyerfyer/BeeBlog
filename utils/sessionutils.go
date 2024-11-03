package utils

import (
	"log"

	bee "github.com/beego/beego/v2/server/web"
)

func GetIntSession(c *bee.Controller, sessionField string) (bool, int) {
	sessionValue := c.GetSession(sessionField)
	if sessionValue == nil {
		return true, 0
	}

	if id, ok := sessionValue.(int); !ok {
		log.Println("Failed to convert sessionvalue into integer")
		return false, 0
	} else {
		return true, id
	}
}
