package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tank130701/course-work/todo-app/back-end/internal/errs"
	"github.com/tank130701/course-work/todo-app/back-end/internal/models"

	"net/http"
	"strconv"
)

// @Summary Create Item
// @Security ApiKeyAuth
// @Tags items
// @Description create new item
// @ID create-item
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Param body body models.TodoItem true "Item Body"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/:id [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input models.TodoItem
	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Items
// @Security ApiKeyAuth
// @Tags items
// @Description get all items
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Success 200 {array} models.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/:id [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get Item By Id
// @Security ApiKeyAuth
// @Tags items
// @Description get item by id
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} models.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/:id [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update Item
// @Security ApiKeyAuth
// @Tags items
// @Description update item by id
// @ID update-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Param body body models.UpdateItemInput true "Update Item Body"
// @Success 200 {object} errs.StatusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/:id [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, errs.StatusResponse{Status: "ok"})
}

// @Summary Delete Item
// @Security ApiKeyAuth
// @Tags items
// @Description delete item by id
// @ID delete-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} errs.StatusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/:id [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, errs.StatusResponse{Status: "ok"})
}
