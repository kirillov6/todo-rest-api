package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type (
	errorResponse struct {
		Message string `json:"message"`
	}

	statusResponse struct {
		Status string `json:"status"`
	}
)

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Fatal(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
