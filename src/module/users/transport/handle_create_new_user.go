package usertrpt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	userbusiness "cmc/module/users/business"
	usermodel "cmc/module/users/model"
	userstorage "cmc/module/users/storage"
)

func HanleCreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem usermodel.User

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// preprocess title - trim all spaces
		dataItem.Username = strings.TrimSpace(dataItem.Username)

		// setup dependencies
		storage := userstorage.NewMySQLStorage(db)
		biz := userbusiness.NewCreateUserBiz(storage)

		if err := biz.CreateNewUser(c.Request.Context(), &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem.Id})
	}
}
