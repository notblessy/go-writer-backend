package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseOK(ginCtx *gin.Context, response *Response) {
	defaultValueSuccess(response)
	ginCtx.JSON(http.StatusOK, response)
}

func ResponseError(ginCtx *gin.Context, response *Response) {
	defaultValueError(response)
	ginCtx.JSON(http.StatusInternalServerError, response)
}

func defaultValueSuccess(response *Response) {
	if response.Status == "" {
		response.Status = "SUCCESS"
	}

	if response.Message == "" {
		response.Message = "SUCCESS"
	}
}

func defaultValueError(response *Response) {
	if response.Status == "" {
		response.Status = "ERROR"
	}

	if response.Message == "" {
		response.Message = "ERROR"
	}
}
