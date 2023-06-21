package simulation

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSimulationController() Controller {
	phoneInfoRepository := &PhoneInfoRepositoryImpl{}
	return Controller{
		getPhoneInformationUseCase: GetPhoneInformationUseCase{
			phoneInfoRepository: phoneInfoRepository,
		},
		buyInsuranceUseCase: BuyInsuranceUseCase{
			paymentRepository:      &PaymentRepositoryImpl{},
			buyInsuranceRepository: &BuyInsuranceRepositoryImpl{},
			phoneInfoRepository:    phoneInfoRepository,
		},
	}
}

type Controller struct {
	getPhoneInformationUseCase GetPhoneInformationUseCase
	buyInsuranceUseCase        BuyInsuranceUseCase
}

// GetInsuranceSimulation godoc
// @Summary Show simulation data.
// @Description Show simulation data.
// @Schemes
// @Tags Simulation
// @Accept */*
// @Produce json
// @Param body body Request true "Resquest to simulation info"
// @Success 200 {object} Response
// @Router /simulation [post]
func (ctrl *Controller) GetInsuranceSimulation(ctx *gin.Context) {
	var request Request
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	phoneInfo, err := ctrl.getPhoneInformationUseCase.Invoke(request.PhoneBrandCode, request.PhoneModelCode)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := MapperFromPhoneInfoToResponse(phoneInfo)
	ctx.IndentedJSON(http.StatusOK, response)
}

// BuyInsurance godoc
// @Summary Buy new insurance.
// @Description Buy new insurance.
// @Tags Simulation
// @Accept */*
// @Param body body Request true "Buy a new insurance"
// @Produce json
// @Success 200 {object} BuyResponse
// @Router /buy [post]
func (ctrl *Controller) BuyInsurance(ctx *gin.Context) {
	var request Request

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	phoneInfo, paymentInfo, err := ctrl.buyInsuranceUseCase.Invoke(request.UserID, request.PhoneBrandCode, request.PhoneModelCode)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := BuyResponse{
		Response:  MapperFromPhoneInfoToResponse(phoneInfo),
		PaymentID: paymentInfo.ID,
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
