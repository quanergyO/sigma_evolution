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

	return router
}
