package insurance

import (
	"net/http"
	"github.com/gin-gonic/gin"
	logs "github.com/johnazedo/ms-insurance/src/logs"
)

type Controller struct {
	Logger                  logs.LogService
	GetInsuranceInfoUseCase GetInsuranceInfoUseCase
	CancelInsuranceUseCase  CancelInsuranceUseCase
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
		log := logs.LogInput{
			Level:   "ERROR",
			Class:   "InsuranceController",
			Method:  "GetInsuranceInformation",
			Message: err.Error(),
		}
		go ctrl.Logger.SendLog(log)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	insuranceInfo, err := ctrl.GetInsuranceInfoUseCase.Invoke(request.UserID)
	// err = errors.New("This is a mock error")
	if err != nil {
		message := "message: Could not get information about the insurance"
		log := logs.LogInput{
			Level:   "ERROR",
			Class:   "InsuranceController",
			Method:  "GetInsuranceInformation",
			Message: message,
		}
		go ctrl.Logger.SendLog(log)
		ctx.JSON(http.StatusInternalServerError, message)
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
		message := "message: Wrong informations sent"
		log := logs.LogInput{
			Level:   "ERROR",
			Class:   "InsuranceController",
			Method:  "CancelInsurance",
			Message: message,
		}
		go ctrl.Logger.SendLog(log)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	insuranceInfo, err := ctrl.CancelInsuranceUseCase.Invoke(request.UserID)
	if err != nil {
		message := "message: Could not cancel this insurance"
		log := logs.LogInput{
			Level:   "ERROR",
			Class:   "InsuranceController",
			Method:  "CancelInsurance",
			Message: message,
		}
		go ctrl.Logger.SendLog(log)
		ctx.JSON(http.StatusInternalServerError, message)
		return
	}

	response := MapperFromInsuranceInfoToResponse(insuranceInfo)

	ctx.IndentedJSON(http.StatusOK, response)
}
