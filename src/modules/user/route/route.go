package route

import (
	"context"
	"net"
	"social-network/domain"
	"social-network/proto/pb"

	userBsn "social-network/modules/user/business"
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
}

func initUserGrpcServer(addr string, ctx context.Context) {
	// user grpc server
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	// user dependencies
	UserRepository := userRepo.NewMongoStorage(ctx.Value(utils.CtxMongodb).(*mongo.Database))
	userBusiness := userBsn.NewBusiness(UserRepository)
	userRpc := userHdl.NewGrpcService(userBusiness)
	pb.RegisterUserServiceServer(server, userRpc)

	go func() {
		// Start the gRPC server
		if err := server.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()
}

func Setup(router *gin.RouterGroup, ctx context.Context) {
	initUserGrpcServer(utils.GodotEnv("USER_GRPC_ADDR"), ctx)
}
