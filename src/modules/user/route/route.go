package route

import (
	"context"
	"net"
	"social-network/domain"
	"social-network/middleware"
	"social-network/proto/pb"

	userBsn "social-network/modules/user/business"
	userApiHdl "social-network/modules/user/delivery/api"
	userHdl "social-network/modules/user/delivery/rpc"
	userRepo "social-network/modules/user/repository/mongo"
	"social-network/utils"

	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type UserService interface {
	GetAuth(ctx context.Context, email, password string) (*domain.User, error)
	GetAuthUser(ctx context.Context, id string) (*domain.User, error)
}

type UserApiService interface {
	GetUserHdl() func(*gin.Context)
	GetOnlineUsersHdl() func(*gin.Context)
	GetNewMembersHdl() func(*gin.Context)
}

func initUserGrpcServer(addr string, ctx context.Context) {
	// user grpc server
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	// user dependencies
	userRepository := userRepo.NewMongoStorage(ctx.Value(utils.CtxMongodb).(*mongo.Database))
	userBusiness := userBsn.NewBusiness(userRepository)
	userRpc := userHdl.NewGrpcService(userBusiness)
	pb.RegisterUserServiceServer(server, userRpc)

	go func() {
		// Start the gRPC server
		if err := server.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()
}

func initUserApiService(ctx context.Context) UserApiService {
	// auth dependencies
	userRepository := userRepo.NewMongoStorage(ctx.Value(utils.CtxMongodb).(*mongo.Database))
	userBusiness := userBsn.NewBusiness(userRepository)
	authApi := userApiHdl.NewAPI(userBusiness)
	return authApi
}

func Setup(router *gin.RouterGroup, ctx context.Context) {
	initUserGrpcServer(utils.GodotEnv("USER_GRPC_ADDR"), ctx)

	userApi := initUserApiService(ctx)
	userGroup := router.Group("/users")
	{
		userGroup.GET("/get-users", middleware.AuthRequire(), userApi.GetUserHdl())
		userGroup.GET("/online-users", middleware.AuthRequire(), userApi.GetOnlineUsersHdl())
		userGroup.GET("/new-members", middleware.AuthRequire(), userApi.GetNewMembersHdl())
	}
}
