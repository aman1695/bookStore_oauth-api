package app

import (
	"github.com/aman1695/bookStore_oauth-api/src/domain/access_token"
	"github.com/aman1695/bookStore_oauth-api/src/http"
	"github.com/aman1695/bookStore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var(
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.PUT("/oauth/access_token", atHandler.UpdateExpirationTime)
	router.Run(":8080")
}
