package main

import (
	"context"
	"fmt"
	"net/http"
	authBsn "social-network/modules/auth/business"
	authHdl "social-network/modules/auth/delivery/api"
	authRepo "social-network/modules/users/repository/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://bienpn:iMv2MfJNVg6Efim6@cluster0.cgjsh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println("DB err!")
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("social-network")

	// auth dependencies
	authRepository := authRepo.NewMongoStorage(db)
	authBusiness := authBsn.NewBusiness(authRepository)
	authApi := authHdl.NewAPI(authBusiness)

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
