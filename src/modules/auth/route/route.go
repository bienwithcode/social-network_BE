package route

import (
	"social-network/proto/pb"

	authBsn "social-network/modules/auth/business"
	authHdl "social-network/modules/auth/delivery/api"
	authRpc "social-network/modules/auth/repository/rpc"
	"social-network/utils"

	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthService interface {
	LoginHdl() func(*gin.Context)
}

func initAuthGrpcService(addr string) AuthService {
	// auth grpc client
	authGrpcClientConn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer authGrpcClientConn.Close()

	// auth dependencies
	authGrpcClient := pb.NewUserServiceClient(authGrpcClientConn)
	authRepo := authRpc.NewGrpcServiceClient(authGrpcClient)
	authBusiness := authBsn.NewBusiness(authRepo)
	authApi := authHdl.NewAPI(authBusiness)
	return authApi
}

func Setup(router *gin.RouterGroup) {
	authGrpc := initAuthGrpcService(utils.GodotEnv("AUTH_GRPC_ADDR"))
	router.POST("/login", authGrpc.LoginHdl())
}
