package insurance

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInsuranceController() controller {
	repository := InsuranceRepositoryImpl{}
	return controller{
		GetInsuranceInfoUseCase: GetInsuranceInfoUseCase{
			InsuranceRepository: &repository,
		},
		CancelInsuranceUseCase: CancelInsuranceUseCase{
			InsuranceRepository: &repository,
		},
	}
}

type controller struct {
	GetInsuranceInfoUseCase
	CancelInsuranceUseCase
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

	response := MapperFromInsuranceInfoToResponse(insuranceInfo)

	ctx.IndentedJSON(http.StatusOK, response)
}


// CancelInsurance godoc
// @Summary Cancel insurance.
// @Description Cancel insurance.
// @Schemes
// @Tags Insurance Management
// @Accept */*
// @Produce json
// @Param body body Request true "Send user id to cancel insurance"
// @Success 200 {object} Response
// @Router /cancel [post]
func (ctrl *controller) CancelInsurance(ctx *gin.Context) {
	var request Request

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	insuranceInfo, err := ctrl.CancelInsuranceUseCase.Invoke(request.UserID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := MapperFromInsuranceInfoToResponse(insuranceInfo)

	ctx.IndentedJSON(http.StatusOK, response)
}