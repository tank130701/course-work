package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/tank130701/course-work/todo-app/back-end/docs"
	"github.com/tank130701/course-work/todo-app/back-end/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = true // разрешить все источники, измените это в соответствии с вашими требованиями безопасности
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	router.Use(cors.New(config))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/refresh", h.refreshToken)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
