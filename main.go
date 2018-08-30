package main

import (
	_ "BeeGoPj/routers"
	"github.com/astaxie/beego"
	_ "BeeGoPj/models"
)

func main() {
	//models.InsertUser()
	//
	//models.UpdateUser()
	beego.Run()

}

