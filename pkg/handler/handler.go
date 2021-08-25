package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kirillov6/todo-rest-api/pkg/services"
)

const (
	listIdParamKey = "list_id"
	itemIdParamKey = "item_id"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getLists)
			lists.POST("/", h.createList)

			listIdPath := fmt.Sprintf("/:%s", listIdParamKey)
			lists.GET(listIdPath, h.getList)
			lists.PUT(listIdPath, h.updateList)
			lists.DELETE(listIdPath, h.deleteList)

			items := lists.Group(":list_id/items")
			{
				items.GET("/", h.getItems)
				items.POST("/", h.createItem)
			}
		}

		items := api.Group("/items")
		{
			itemIdPath := fmt.Sprintf("/:%s", itemIdParamKey)
			items.GET(itemIdPath, h.getItem)
			items.PUT(itemIdPath, h.updateItem)
			items.DELETE(itemIdPath, h.deleteItem)
		}
	}

	return router
}
