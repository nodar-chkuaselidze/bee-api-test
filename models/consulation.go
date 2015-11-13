package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Consultation struct {
	Id               int
	Name             string
	PainSite         string
	PainNote         string
	Age              uint8
	Location         string
	ExaminationNotes string
	Date             time.Time
	Patient          *Patient `orm:"rel(fk)"`
}

func (c *Consultation) TableName() string {
	return "consultation"
}

func GetAllConsultations() []*Consultation {
	return make([]*Consultation, 1)
}

func init() {
	orm.RegisterModel(new(Consultation))
}
