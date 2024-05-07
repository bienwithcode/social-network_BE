package api

import (
	"context"
	"net/http"
	"social-network/domain"
	"social-network/modules/auth/model"
	"social-network/utils"

	"github.com/gin-gonic/gin"
)

type Business interface {
	Login(ctx context.Context, data *model.AuthEmailPassword) (*model.TokenResponse, error)
	GetAuthUser(ctx context.Context, data *model.AuthUserId) (*domain.User, error)
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
			panic(err.Error())
		}

		response, err := api.business.Login(c.Request.Context(), &data)

		if err != nil {
			panic(err.Error())
		}
		utils.WriteSuccessResponse(c, "success", http.StatusOK, &response)
	}
}

func (api *api) AuthUserHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		data := model.AuthUserId{
			Id: "62a75af7d4327f33ccdf8010",
		}

		// if err := c.ShouldBind(&data); err != nil {
		// 	panic(err.Error())
		// }

		response, err := api.business.GetAuthUser(c.Request.Context(), &data)

		if err != nil {
			panic(err.Error())
		}
		utils.WriteSuccessResponse(c, "success", http.StatusOK, &response)
	}
}
