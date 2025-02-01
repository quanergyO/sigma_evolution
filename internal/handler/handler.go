package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanergyO/sigma_evolution/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("api/v1")
	{
		skills := api.Group("/skills")
		{
			skills.GET("/", h.skillsGetAll)
			skills.GET("/:id", h.skillsGetById)
			skills.PATCH("/:id", h.skillsUpdate)
			skills.POST("/", h.skillsCreate)
			skills.DELETE("/:id", h.skillsDelete)
		}
	}

	return router
}
