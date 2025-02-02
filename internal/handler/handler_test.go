package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/quanergyO/sigma_evolution/internal/handler"
	"github.com/quanergyO/sigma_evolution/internal/service"
	mocks "github.com/quanergyO/sigma_evolution/mock"
	"github.com/quanergyO/sigma_evolution/types"
	"github.com/stretchr/testify/assert"
)

func TestSkillsCreate_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockISkils(ctrl)

	mockService.EXPECT().
		Create(types.Skill{Name: "Programming"}).
		Return(1, nil)

	handler := handler.NewHandler(&service.Service{ISkils: mockService})

	router := gin.Default()
	router.POST("/api/v1/skills", handler.SkillsCreate)

	input := types.Skill{Name: "Programming"}
	payload, _ := json.Marshal(input)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/skills", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/Json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]int
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshar response body: %v", err)
	}

	assert.Equal(t, 2, response["id"])

}

func TestSkillsCreate_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockISkils(ctrl)

	mockService.EXPECT().
		Create(types.Skill{Name: "Programming"}).
		Return(0, errors.New("service error"))

	handler := handler.NewHandler(&service.Service{ISkils: mockService})
	router := gin.Default()
	router.POST("/api/v1/skills", handler.SkillsCreate)

	input := types.Skill{Name: "Programming"}
	payload, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/skills", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, "service error", response["message"])
}

func TestSkillsCreate_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockISkils(ctrl)
	handler := handler.NewHandler(&service.Service{ISkils: mockService})

	router := gin.Default()
	router.POST("/api/v1/skills", handler.SkillsCreate)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/skills", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, "Invalid request body", response["message"])
}
