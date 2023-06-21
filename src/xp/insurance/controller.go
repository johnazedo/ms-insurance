package insurance

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInsuranceController() Controller {
	repository := InsuranceRepositoryImpl{}
	return Controller{
		getInsuranceInfoUseCase: GetInsuranceInfoUseCase{
			insuranceRepository: &repository,
		},
		cancelInsuranceUseCase: CancelInsuranceUseCase{
			insuranceRepository: &repository,
		},
	}
}

type Controller struct {
	getInsuranceInfoUseCase GetInsuranceInfoUseCase
	cancelInsuranceUseCase  CancelInsuranceUseCase
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
func (ctrl *Controller) GetInsuranceInformation(ctx *gin.Context) {
	var request Request

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	insuranceInfo, err := ctrl.getInsuranceInfoUseCase.Invoke(request.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "message: Could not get information about the insurance")
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
func (ctrl *Controller) CancelInsurance(ctx *gin.Context) {
	var request Request

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, "message: Wrong informations sent")
		return
	}

	insuranceInfo, err := ctrl.cancelInsuranceUseCase.Invoke(request.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "message: Could not cancel this insurance")
		return
	}

	response := MapperFromInsuranceInfoToResponse(insuranceInfo)

	ctx.IndentedJSON(http.StatusOK, response)
}
