package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader       = "Authorization"
	userIdContextKey = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userIdContextKey, userId)
}

func getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userIdContextKey)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := userId.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id has invalid type")
		return 0, errors.New("user id has invalid type")
	}

	return idInt, nil
}
