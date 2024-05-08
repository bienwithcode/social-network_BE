package main

import (
	"context"
	"social-network/db"
	"social-network/middleware"
	authRoute "social-network/modules/auth/route"
	userRoute "social-network/modules/user/route"
	"social-network/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	dbCollection := db.InitMongodb(utils.GodotEnv("MONGO_DB_CON_STR"), utils.GodotEnv("MONGO_DB_COLLECTION"))
	ctx := context.Background()
	ctx = context.WithValue(ctx, utils.CtxMongodb, dbCollection)

	// router
	router := gin.Default()
	router.Use(middleware.Recover(), middleware.Cors())
	routerGroup := router.Group("/")
	authRoute.Setup(routerGroup)
	userRoute.Setup(routerGroup, ctx)

	router.Run(":9000")
}
