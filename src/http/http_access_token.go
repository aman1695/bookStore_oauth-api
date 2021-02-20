package http

import (
	"github.com/aman1695/bookStore_oauth-api/src/domain/access_token"
	"github.com/aman1695/bookStore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	UpdateExpirationTime(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(strings.TrimSpace(c.Param("access_token_id")))
	if err !=  nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK,accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var accessToken access_token.AccessToken
	if err := c.ShouldBindJSON(&accessToken); err != nil {
		c.JSON(http.StatusBadRequest,errors.NewBadRequestError("Invalid Json Body!"))
		return
	}
    if err:= handler.service.Create(accessToken); err != nil {
    	c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK,accessToken)
}

func (handler *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	var accessToken access_token.AccessToken
	if err := c.ShouldBindJSON(&accessToken); err != nil {
		c.JSON(http.StatusBadRequest,errors.NewBadRequestError("Invalid Json Body!"))
		return
	}
	if err:= handler.service.UpdateExpirationTime(accessToken); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK,accessToken)
}