package controllers

import (
	//"encoding/json"
	"github.com/nodar-chkuaselidze/bee-api-test/models"

	"github.com/astaxie/beego"
	"strconv"
)

type PatientController struct {
	beego.Controller
}

// @Title Get
// @Description get patient by ID
// @Success 200 {object} models.Patient
// @router / [get]
func (p *PatientController) Get() {
	requestParam := p.Ctx.Input.Param(":id")

	requestId, err := strconv.Atoi(requestParam)

	if err != nil {
		beego.Debug(err)
		p.Abort("400")
	}

	p.Data["json"], err = models.GetPatient(requestId)

	if err != nil {
		p.Abort("404")
	}
	p.ServeJson()
}
