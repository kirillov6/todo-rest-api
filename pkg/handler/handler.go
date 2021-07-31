package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getLists)
			lists.POST("/", h.createList)
			lists.GET("/:list_id", h.getList)
			lists.PUT("/:list_id", h.updateList)
			lists.DELETE("/:list_id", h.deleteList)
		}

		items := lists.Group(":list_id/items")
		{
			items.GET("/", h.getItems)
			items.POST("/", h.createItem)
			items.GET("/:item_id", h.getItem)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}
	}

	return router
}
