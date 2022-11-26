package main

import (
	usertrpt "cmc/module/users/delivery"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(mysqlcontainer)/cmc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/user", usertrpt.HanleCreateUser(db)) // updated
		//...
	}
	router.Run(":9000")
}
