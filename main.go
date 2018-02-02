package main

import (
	_ "pinnapi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "POST", "GET"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	beego.Run()
}
