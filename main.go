package main

import (
	"fmt"
	"net/http"
	_ "pinnapi/routers"

	"pinnapi/healthcheck"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/toolbox"
)

func pleaseDontPeek(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Please get out!")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{

		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	toolbox.AddHealthCheck("database", &healthcheck.DatabaseCheck{})
	beego.ErrorHandler("404", pleaseDontPeek)
	beego.Run()
}
