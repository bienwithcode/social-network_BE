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
	GetUsers(ctx context.Context, authUserId string, paging *utils.Pagination, filter *model.Filter) ([]*domain.User, error)
	GetOnlineUsers(ctx context.Context, authUserId string) ([]*domain.User, error)
	GetNewMembers(ctx context.Context, authUserId string, paging *utils.Pagination) ([]*domain.User, error)
	GetAuthUser(ctx context.Context, id string) (*domain.User, error)
}

type api struct {
	// serviceCtx sctx.ServiceContext
	business Business
}

func NewAPI(business Business) *api {
	return &api{business: business}
}

func (api *api) GetUserHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		// pagination
		var paging domain.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err.Error())
		}
		pagination := utils.NewPagination(paging.Page, paging.PerPage)

		// filter
		var filter model.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err.Error())
		}

		var authUserId string
		if auth, ok := c.Get("authData"); ok {
			authData, _ := auth.(map[string]interface{})
			authUserId = authData["id"].(string)
		}

		response, err := api.business.GetUsers(c.Request.Context(), authUserId, pagination, &filter)

		if err != nil {
			panic(err.Error())
		}
		utils.WriteSuccessResponse(c, "success", http.StatusOK, &response)

	}
}

func (api *api) GetOnlineUsersHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var authUserId string
		if auth, ok := c.Get("authData"); ok {
			authData, _ := auth.(map[string]interface{})
			authUserId = authData["id"].(string)
		}

		response, err := api.business.GetOnlineUsers(c.Request.Context(), authUserId)

		if err != nil {
			panic(err.Error())
		}
		utils.WriteSuccessResponse(c, "success", http.StatusOK, &response)

	}
}

func (api *api) GetNewMembersHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		// pagination
		var paging domain.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err.Error())
		}
		pagination := utils.NewPagination(paging.Page, paging.PerPage)

		var authUserId string
		if auth, ok := c.Get("authData"); ok {
			authData, _ := auth.(map[string]interface{})
			authUserId = authData["id"].(string)
		}

		response, err := api.business.GetNewMembers(c.Request.Context(), authUserId, pagination)

		if err != nil {
			panic(err.Error())
		}
		utils.WriteSuccessResponse(c, "success", http.StatusOK, &response)

	}
}

func (api *api) GetUserByIdHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("id")
		response, err := api.business.GetAuthUser(c.Request.Context(), userId)

		if err != nil {
			panic(err.Error())
		}
		utils.WriteSuccessResponse(c, "success", http.StatusOK, &response)

	}
}
