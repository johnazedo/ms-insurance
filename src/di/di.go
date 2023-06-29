package di

import (
	infra "github.com/johnazedo/ms-insurance/src/infra"
	logs "github.com/johnazedo/ms-insurance/src/logs"
	phonelist "github.com/johnazedo/ms-insurance/src/offer/phonelist"
	simulation "github.com/johnazedo/ms-insurance/src/offer/simulation"
	"github.com/johnazedo/ms-insurance/src/recurrence"
	insurance "github.com/johnazedo/ms-insurance/src/xp/insurance"
	gorm "gorm.io/gorm"
)

func GetPhoneListController(db *gorm.DB) phonelist.Controller {
	repository := &infra.CellphonesAvailableRepositoryImpl{
		DB: db,
	}
	return phonelist.Controller {
		GetListOfBrandsUseCase: phonelist.GetListOfBrandsUseCase{
			CellphonesAvailableRepository: repository,
		},
		GetListOfModelsUseCase: phonelist.GetListOfModelsUseCase{
			CellphonesAvailableRepository: repository,
		},
	}
}

func GetInsuranceController(db *gorm.DB, logger logs.LogService) insurance.Controller {
	repository := &infra.InsuranceRepositoryImpl{
		DB: db,
	}
	return insurance.Controller {
		Logger: logger,
		GetInsuranceInfoUseCase: insurance.GetInsuranceInfoUseCase{
			InsuranceRepository: repository,
		},
		CancelInsuranceUseCase: insurance.CancelInsuranceUseCase{
			InsuranceRepository: repository,
		},
	}
}

func GetSimulationController(db *gorm.DB) simulation.Controller {
	phoneInfoRepository := &infra.CellphonesAvailableRepositoryImpl{
		DB: db,
	}
	return simulation.Controller{
		GetPhoneInformationUseCase: simulation.GetPhoneInformationUseCase{
			PhoneInfoRepository: phoneInfoRepository,
		},
		BuyInsuranceUseCase: simulation.BuyInsuranceUseCase{
			PaymentRepository:      &infra.PaymentRepositoryImpl{},
			BuyInsuranceRepository: &infra.InsuranceRepositoryImpl{
				DB: db,
			},
			PhoneInfoRepository:    phoneInfoRepository,
		},
	}
}

func GetRecurrenceUseCase(db *gorm.DB, logger logs.LogService) recurrence.RecurrenceUseCase {
	return recurrence.RecurrenceUseCase{
		Logger: logger,
		InsuranceRepository: &infra.InsuranceRepositoryImpl{
			DB: db,
		},
		PaymentRepository: &infra.PaymentRepositoryImpl{},
	}
}