package controllers

import (
	//"encoding/json"
	"github.com/nodar-chkuaselidze/bee-api-test/models"

	"github.com/astaxie/beego"
)

type ConsultationController struct {
	beego.Controller
}

// @Title Get
// @Description get all Consultations
// @Success 200 {object} models.Consultation
// @router / [get]
func (c *ConsultationController) GetAll() {
	consultations := models.GetAllConsultations()
	c.Data["json"] = consultations
	c.ServeJson()
}
