package contracts

import (
	"kis/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")

	vacancies := api.Group("/vacancies")

	vacancies.GET("/", h.getVacanciesByQuery)
	vacancies.POST("/", h.createVacancy)
	vacancies.PATCH("/:id", h.patchVacancy)
	vacancies.GET("/search", h.searchVacancies)

	return router
}
