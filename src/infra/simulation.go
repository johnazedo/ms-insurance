package infra

import . "github.com/johnazedo/ms-insurance/src/offer/simulation"


type PaymentRepositoryImpl struct{}

func (r *PaymentRepositoryImpl) MakePayment(userID string, value float64) (*PaymentInfo, error) {
	return &PaymentInfo{
		ID: "uuasdjf-aidfnkd-adsfksn",
	}, nil
}