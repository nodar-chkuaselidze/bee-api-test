package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Patient struct {
	Id            int
	First_name    string
	Last_name     string
	Gender        string
	Date_birth    time.Time
	Country       string
	City          string
	Street        string
	Title         string
	Color         string
	Url           string
	Blog          string
	Height        uint8
	Weight        uint8
	Email         string
	Frequence     uint32
	Money         uint64
	Phone         string
	Consultations []*Consultation `orm:"reverse(many)"`
}

func (p *Patient) TableName() string {
	return "patient"
}

func GetPatient(id int) (*Patient, error) {
	o := orm.NewOrm()
	o.Using("default")

	patient := &Patient{
		Id: id,
	}

	err := o.Read(patient)

	if err != nil {
		return nil, err
	}

	beego.Trace(patient)

	return patient, nil
}

func GetPatientWithConsultations(id, consultations int) (*Patient, error) {
	o := orm.NewOrm()
	o.Using("default")

	patient, err := GetPatient(id)

	if err != nil {
		return nil, err
	}

	num, err := o.LoadRelated(patient, "Consultations")

	if err != nil {
		return nil, err
	}

	beego.Debug("Loaded consultations: %d", num)

	return patient, nil
}

func init() {
	orm.RegisterModel(new(Patient))
}
