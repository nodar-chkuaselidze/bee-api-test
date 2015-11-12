package main

import (
	_ "github.com/nodar-chkuaselidze/bee-api-test/docs"
	_ "github.com/nodar-chkuaselidze/bee-api-test/models"
	_ "github.com/nodar-chkuaselidze/bee-api-test/routers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DR_Postgres)
	orm.RegisterDataBase("default", "postgres", "user=nod dbname=BeeAPI sslmode=disable")
}

func main() {
	orm.RunCommand()

	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
