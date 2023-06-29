package recurrence

import (
	"fmt"
	"time"

	"github.com/johnazedo/ms-insurance/src/logs"
	"github.com/johnazedo/ms-insurance/src/offer/simulation"
	"github.com/johnazedo/ms-insurance/src/xp/insurance"
)


type InsuranceRepository interface {
	GetInsurances() ([]*insurance.InsuranceInfo, error)
	UpdateValidityInsurance(id string, newValidity string) error
}

// PaymentRepository This repository will be implemented to communicate with payment microservice
type PaymentRepository interface {
	MakePayment(userID string, value float64) (*simulation.PaymentInfo, error)
}

type RecurrenceUseCase struct {
	Logger logs.LogService
	InsuranceRepository InsuranceRepository
	PaymentRepository PaymentRepository
}

func (uc *RecurrenceUseCase) Invoke() {
	for {
		time.Sleep(time.Hour * 24)
		insurances, err := uc.InsuranceRepository.GetInsurances()
		if err != nil {
			log := logs.LogInput{
				Level: "ERROR",
				Class: "InsuranceController",
				Method: "CancelInsurance",
				Message: "message: Error on get insurances",
			}
			uc.Logger.SendLog(log)
		}

		for _, insurance := range insurances {
			year, month, day := time.Now().Date()
			today := fmt.Sprintf("%d/%d/%d", day, month, year)
			if insurance.Status == "ACTIVE" && today == insurance.Validity {
				uc.PaymentRepository.MakePayment(insurance.UserID, insurance.ValuePerMonth)
				uc.InsuranceRepository.UpdateValidityInsurance(insurance.ID, time.Now().AddDate(0, 1, 0).String())
			}
		}
	} 
}
