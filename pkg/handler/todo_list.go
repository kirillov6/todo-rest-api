package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kirillov6/todo-rest-api"
)

type getListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getListsResponse{
		Data: lists,
	})
}

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var list todo.TodoList
	if err := c.BindJSON(&list); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, list)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := getListIdParam(c)
	if err != nil {
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := getListIdParam(c)
	if err != nil {
		return
	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoList.UpdateById(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := getListIdParam(c)
	if err != nil {
		return
	}

	err = h.services.TodoList.DeleteById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}

func getListIdParam(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param(listIdParamKey))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return 0, err
	}

	return id, err
}
