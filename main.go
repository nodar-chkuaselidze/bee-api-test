package main

import (
	_ "github.com/nodar-chkuaselidze/bee-api-test/docs"
	_ "github.com/nodar-chkuaselidze/bee-api-test/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
