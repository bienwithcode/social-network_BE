package route

import (
	"context"
	"social-network/middleware"

	messageBsn "social-network/modules/message/business"
	messageApiHdl "social-network/modules/message/delivery/api"
	messageRepo "social-network/modules/message/repository/mongo"
	"social-network/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserApiService interface {
	GetMessagesHdl() func(*gin.Context)
	GetConversationsHdl() func(*gin.Context)
	CreateMessagesHdl() func(*gin.Context)
	UpdateSeenMessageHdl() func(*gin.Context)
}

func initMessageApiService(ctx context.Context) UserApiService {
	// message dependencies
	messageRepository := messageRepo.NewMongoStorage(ctx.Value(utils.CtxMongodb).(*mongo.Database))
	messageBusiness := messageBsn.NewBusiness(messageRepository)
	messageApi := messageApiHdl.NewAPI(messageBusiness)
	return messageApi
}

func Setup(router *gin.RouterGroup, ctx context.Context) {
	userApi := initMessageApiService(ctx)
	userGroup := router.Group("/messages")
	{
		userGroup.GET("/", middleware.AuthRequire(), userApi.GetMessagesHdl())
		userGroup.GET("/conversations", middleware.AuthRequire(), userApi.GetConversationsHdl())
		userGroup.GET("/create", middleware.AuthRequire(), userApi.CreateMessagesHdl())
		userGroup.GET("/update-seen", middleware.AuthRequire(), userApi.UpdateSeenMessageHdl())
	}
}
