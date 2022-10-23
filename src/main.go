package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	usertrpt "cmc/module/users/transport"
)

func main()  {
	dsn := "root:root@tcp(127.0.0.1:3306)/cmc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
        fmt.Println("err")
    }
	fmt.Println(&db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/user", usertrpt.HanleCreateItem(db))  // updated
		//...
	}
}