package common

import (
	"net/http"
	"test-mnc/shared/model"

	"github.com/gin-gonic/gin"
)

func SendCreateResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusCreated, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusCreated,
			Message: message,
		},
		Data: data,
	})
}

func SendLogoutResponse(c *gin.Context, message string) {
	c.JSON(http.StatusCreated, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusCreated,
			Message: message,
		},
	})
}


func SendSingleResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

func SendErrorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, &model.Status{
		Code:    code,
		Message: message,
	})
}

func SendNoContentResponse(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}