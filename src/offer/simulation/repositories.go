package simulation

type PhoneInfoRepositoryImpl struct {}

func (r *PhoneInfoRepositoryImpl) SearchByBrandAndModel(phoneBrand string, phoneModel string) (*PhoneInfo, error) {
	return &PhoneInfo{
		PhoneBrand: "Samsumg",
		PhoneModel: "S22 128GB",
		ValuePerMonth: 60.0,
		Franchise: 1500.0,
	}, nil
}

type PaymentRepositoryImpl struct {}

func (r *PaymentRepositoryImpl) MakePayment(userID string, value float64) (*PaymentInfo, error) {
	return &PaymentInfo{
		ID: "uuasdjf-aidfnkd-adsfksn",
	}, nil
}

type BuyInsuranceRepositoryImpl struct {}

func (r *BuyInsuranceRepositoryImpl) SaveInsurance(userID string, paymentID string, phoneBrand string, phoneModel string) error {
	return nil
}

