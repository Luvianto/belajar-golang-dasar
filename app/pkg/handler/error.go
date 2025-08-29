package handler

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status     bool       `json:"status"`
	StatusCode int        `json:"status_code"`
	Message    string     `json:"message"`
	Errors     []ApiError `json:"errors,omitempty"`
}

var errorResponsePool = sync.Pool{
	New: func() interface{} {
		return &ErrorResponse{}
	},
}

func Error(c *gin.Context, statusCode int, message string, errors []ApiError) {
	response := errorResponsePool.Get().(*ErrorResponse)

	response.Status = false
	response.StatusCode = statusCode
	response.Message = message
	response.Errors = errors

	c.AbortWithStatusJSON(statusCode, response)

	errorResponsePool.Put(response)
}
