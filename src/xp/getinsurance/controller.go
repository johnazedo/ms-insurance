package getinsurance

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInsuranceController() controller {
	return controller{
		GetInsuranceInfoUseCase: GetInsuranceInfoUseCase{
			InsuranceRepository: &InsuranceRepositoryImpl{},
		},
	}
}

type controller struct {
	GetInsuranceInfoUseCase
}

// GetInsuranceInformation godoc
// @Summary Get Insurance information.
// @Description Get Insurance information.
// @Schemes
// @Tags Insurance Management
// @Accept */*
// @Produce json
// @Param body body Request true "Send user id to get insurance information"
// @Success 200 {object} Response
// @Router /insurance [post]
func (ctrl *controller) GetInsuranceInformation(ctx *gin.Context) {
	var request Request

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	insuranceInfo, err := ctrl.GetInsuranceInfoUseCase.Invoke(request.UserID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := Response{
		PhoneBrand:    insuranceInfo.PhoneBrand,
		PhoneModel:    insuranceInfo.PhoneModel,
		ValuePerMonth: insuranceInfo.ValuePerMonth,
		Franchise:     insuranceInfo.Franchise,
		Validity:      insuranceInfo.Validity,
		Status:        insuranceInfo.Status,
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
