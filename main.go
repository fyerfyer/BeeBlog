package main

import (
	// "fmt"
	_ "beeblog/routers"
	"beeblog/utils"
	"log"
	// "time"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	utils.SQLinit()
	// o := orm.NewOrm()
	if err := orm.RunSyncdb("default", false, true); err != nil {
		log.Fatalf("Failed to sync database: %v", err)
	}

	beego.Run()
}
