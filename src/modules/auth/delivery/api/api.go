package api

import (
	"context"
	"net/http"
	"social-network/modules/auth/model"
	"social-network/utils"

	"github.com/gin-gonic/gin"
)

type Business interface {
	Login(ctx context.Context, data *model.AuthEmailPassword) (*model.TokenResponse, error)
}

type api struct {
	// serviceCtx sctx.ServiceContext
	business Business
}

func NewAPI(business Business) *api {
	return &api{business: business}
}

func (api *api) LoginHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.AuthEmailPassword

		if err := c.ShouldBind(&data); err != nil {
			utils.WriteErrorResponse(c, http.StatusBadRequest, err)
			return
		}

		response, err := api.business.Login(c.Request.Context(), &data)

		if err != nil {
			utils.WriteErrorResponse(c, http.StatusBadRequest, err)
			return
		}
		utils.WriteSuccessResponse(c, "success", http.StatusOK, &response)
	}
}
