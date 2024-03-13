package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tank130701/course-work/todo-app/back-end/internal/errs"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		errs.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		errs.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		errs.NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		errs.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func (h *Handler) refreshToken(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		errs.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		errs.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		errs.NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	// Используйте новый метод RefreshToken для обновления access токена
	newAccessToken, err := h.services.Authorization.RefreshToken(headerParts[1])
	if err != nil {
		errs.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Отправьте новый access токен обратно клиенту
	c.JSON(http.StatusOK, map[string]string{
		"access_token": newAccessToken,
	})
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
