package claim

import (
	"github.com/gin-gonic/gin"
)

func GetClaimController() controller {
	return controller{}
}

type controller struct {
	// Put use case here
}


func (ctrl *controller) ClaimInsurance(ctx *gin.Context) {
	
}