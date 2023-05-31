package cancel

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	// Put usecase here
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
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := Response {
		Status: "CANCELED",
	}

	ctx.IndentedJSON(http.StatusOK, response)
}