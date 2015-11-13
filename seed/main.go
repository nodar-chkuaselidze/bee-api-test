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
	beego.AppConfigPath = "../conf"
}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	patient := GeneratePatient(1, 10)

	o.Insert(patient)
	o.InsertMulti(10, patient.Consultations)
}
