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
			skills.GET("/", h.SkillsGetAll)
			skills.GET("/:id", h.SkillsGetById)
			skills.PATCH("/:id", h.SkillsUpdate)
			skills.POST("/", h.SkillsCreate)
			skills.DELETE("/:id", h.SkillsDelete)
		}
	}

	return router
}
