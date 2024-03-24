package handler

import (
	"errors"
	"github.com/tank130701/course-work/todo-app/back-end/internal/errs"
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errs.errorResponse
// @Failure 500 {object} errs.errorResponse
// @Failure default {object} errs.errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		var myErr *errs.ErrorUsernameAlreadyExists
		if errors.As(err, &myErr) {
			errs.NewErrorResponse(c, http.StatusConflict, err.Error())
			return
		}
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errs.errorResponse
// @Failure 500 {object} errs.errorResponse
// @Failure default {object} errs.errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		var myErr *errs.ErrorNotFound
		if errors.As(err, &myErr) {
			errs.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Создание HTTP-only cookie для refresh token
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true, // использовать только при передаче через HTTPS
		Path:     "/",
	})

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": accessToken,
		// "refreshToken": refreshToken, // Удалите эту строку, если вы не хотите отправлять refresh token в JSON ответе
	})
}

// @Summary Logout
// @Tags auth
// @Description logout
// @ID logout
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Failure 400,404 {object} errs.errorResponse
// @Failure 500 {object} errs.errorResponse
// @Failure default {object} errs.errorResponse
// @Router /auth/logout [post]
func (h *Handler) logout(c *gin.Context) {
	// Удаление HTTP-only cookie, установив срок его действия в прошлом
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true, // использовать только при передаче через HTTPS
		Path:     "/",
	})

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
