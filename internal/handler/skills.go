package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quanergyO/sigma_evolution/internal/handler/response"
	"github.com/quanergyO/sigma_evolution/types"
)

func (h *Handler) SkillsGetAll(c *gin.Context) {

	skills, err := h.service.ISkils.GetAll()
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Data": skills,
	})
}

func (h *Handler) SkillsGetById(c *gin.Context) {
	slog.Info("Not implemented")
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status":     "Not implemented",
		"MethodName": "SkillsGetById",
	})
}

func (h *Handler) SkillsCreate(c *gin.Context) {
	slog.Info("Not implemented")

	var input types.Skill
	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.ISkils.Create(input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SkillsUpdate(c *gin.Context) {
	slog.Info("Not implemented")
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status":     "Not implemented",
		"MethodName": "SkillsUpdate",
	})
}

func (h *Handler) SkillsDelete(c *gin.Context) {
	slog.Info("Not implemented")
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status":     "Not implemented",
		"MethodName": "SkillsDelete",
	})
}
