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

	vacancies := api.Group("/vacancies").Use(gin.BasicAuth(gin.Accounts{"bib": "qwert"}))

	vacancies.GET("/", h.getVacanciesByQuery)
	vacancies.POST("/", h.createVacancy)
	vacancies.GET("/search", h.searchVacancies)
	vacancies.PATCH("/:id", h.patchVacancies)

	bookmarks := api.Group("/").Use(gin.BasicAuth(gin.Accounts{"bib": "qwert"}))

	bookmarks.POST("/add-bookmark", h.createBookmark)
	bookmarks.DELETE("/delete-bookmark", h.deleteBookmark)

	return router
}
