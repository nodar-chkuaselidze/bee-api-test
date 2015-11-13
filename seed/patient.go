package main

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/nodar-chkuaselidze/bee-api-test/models"
)

func GeneratePatient(id int, consultations int) *models.Patient {
	patient := &models.Patient{
		Id: id,
	}

	patient.Gender = "Female"
	patient.Email = randomdata.Email()

	patient.First_name = randomdata.FirstName(randomdata.Female)
	patient.Last_name = randomdata.LastName()

	patient.Date_birth = time.Now()
	patient.City = randomdata.City()
	patient.Country = randomdata.Country(randomdata.TwoCharCountry)
	patient.Money = uint64(randomdata.Number(100, 1000))
	patient.Height = uint8(randomdata.Number(50, 220))
	patient.Weight = uint8(randomdata.Number(20, 300))
	patient.Consultations = make([]*models.Consultation, consultations)

	for i := 1; i <= consultations; i++ {
		consult := GenerateConsultation(i, patient)
		patient.Consultations[i-1] = consult
	}

	return patient
}

func GenerateConsultation(id int, p *models.Patient) *models.Consultation {
	consultation := &models.Consultation{
		Id: id,
	}

	consultation.Age = uint8(randomdata.Number(1, 200))
	consultation.Date = time.Now()
	consultation.Location = randomdata.Address()
	consultation.PainNote = "something .." + randomdata.Letters(3)
	consultation.PainSite = "meredy .." + randomdata.Letters(3)
	consultation.Patient = p

	return consultation
}
