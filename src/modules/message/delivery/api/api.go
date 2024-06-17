package api

import (
	"context"
	"net/http"
	"social-network/domain"
	"social-network/utils"

	"github.com/gin-gonic/gin"
)

type Business interface {
	GetConversations(ctx context.Context, authUserId string) ([]*domain.Message, error)
}

type api struct {
	// serviceCtx sctx.ServiceContext
	business Business
}

func NewAPI(business Business) *api {
	return &api{business: business}
}

func (api *api) GetConversationsHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var authUserId string
		if auth, ok := c.Get("authData"); ok {
			authData, _ := auth.(map[string]interface{})
			authUserId = authData["id"].(string)
		}

		response, err := api.business.GetConversations(c.Request.Context(), authUserId)

		if err != nil {
			panic(err.Error())
		}
		utils.WriteSuccessResponse(c, "success", http.StatusOK, &response)

	}
}

func (api *api) GetMessagesHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		// TODO

	}
}

func (api *api) CreateMessagesHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		// TODO

	}
}

func (api *api) UpdateSeenMessageHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		// TODO

	}
}
