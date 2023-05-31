package getinsurance

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	// Put usecase here
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

	response := Response {
		PhoneBrand: "Samsung",
		PhoneModel: "S22 128",
		ValuePerMonth: 60.00,
		Franchise: 1200.00,
		Validity: "11/05/2024",
		Status: "ACTIVE",
	}

	ctx.IndentedJSON(http.StatusOK, response)
}