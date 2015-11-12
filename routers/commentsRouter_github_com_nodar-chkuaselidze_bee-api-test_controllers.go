package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/nodar-chkuaselidze/bee-api-test/controllers:ConsultationController"] = append(beego.GlobalControllerRouter["github.com/nodar-chkuaselidze/bee-api-test/controllers:ConsultationController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/nodar-chkuaselidze/bee-api-test/controllers:PatientController"] = append(beego.GlobalControllerRouter["github.com/nodar-chkuaselidze/bee-api-test/controllers:PatientController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

}
