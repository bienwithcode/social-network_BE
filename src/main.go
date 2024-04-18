package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	authBsn "social-network/modules/auth/business"
	authHdl "social-network/modules/auth/delivery/api"
	authRpc "social-network/modules/auth/repository/rpc"

	userBsn "social-network/modules/user/business"
	userHdl "social-network/modules/user/delivery/rpc"
	userRepo "social-network/modules/user/repository/mongo"
	pb "social-network/proto/pb"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://bienpn:iMv2MfJNVg6Efim6@cluster0.cgjsh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	db := client.Database("social-network")

	// auth grpc client
	authGrpcClientConn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer authGrpcClientConn.Close()

	// user grpc server
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Tạo một gRPC server
	type userServiceServer struct{}
	server := grpc.NewServer()

	// auth dependencies
	authGrpcClient := pb.NewUserServiceClient(authGrpcClientConn)
	authRepo := authRpc.NewGrpcServiceClient(authGrpcClient)
	authBusiness := authBsn.NewBusiness(authRepo)
	authApi := authHdl.NewAPI(authBusiness)

	// user dependencies
	UserRepository := userRepo.NewMongoStorage(db)
	userBusiness := userBsn.NewBusiness(UserRepository)
	userRpc := userHdl.NewGrpcService(userBusiness)
	pb.RegisterUserServiceServer(server, userRpc)

	go func() {
		// Start the gRPC server
		if err := server.Serve(listen); err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/login", authApi.LoginHdl()) // updated
		//...
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	fmt.Println("init server")
	router.Run(":9000")
}
