package utils

import (
	"beeblog/models"
	"log"

	bee "github.com/beego/beego/v2/server/web"
)

type contextKey string

var contextKeyIsAuthenticated = contextKey("isAuthenticated")

func IsAuthenticated(c *bee.Controller) bool {
	sessionValue := c.GetSession("authenticatedUserID")
	if sessionValue == nil {
		return false
	}

	// if we get the id, we need to second check
	// if the id has been unauthenticated by the database
	// if we know from the context that the user is authenticated
	// then that's it
	contextValue := c.Ctx.Input.GetData(contextKeyIsAuthenticated)

	// otherwise, we set the context
	if authValue, ok := contextValue.(int); ok && authValue > 0 {
		return true
	}
	id, ok := sessionValue.(int)
	if !ok {
		log.Println("Failed to convert session value to string")
		return false
	}
	active, err := models.GetUserStatusById(id)
	if err != nil {
		log.Println("Failed to fetch the user in the database")
		return false
	}
	c.Ctx.Input.SetData(contextKeyIsAuthenticated, active > 0)

	return active > 0
}
