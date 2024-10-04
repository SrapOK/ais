package contracts

import (
	model "kis/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getVacanciesByQuery(c *gin.Context) {
	params := new(model.QueryDTO)

	err := c.BindQuery(params)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "не удалось прочитать параметры"})
		return
	}

	vacancies, err := h.services.GetVacancies(*params)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить вакансии"})
		return
	}

	res := newVacanciesResponseDTO(vacancies, params.Page)

	c.JSON(http.StatusOK, res)
}

func (h *Handler) createVacancy(c *gin.Context) {
	vacancy := new(model.Vacancy)

	err := c.BindJSON(vacancy)

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "не удалось прочиать тело запроса"})
		return
	}

	h.services.CreateVacancy(*vacancy)
}

func (h *Handler) searchVacancies(c *gin.Context) {

}

func (h *Handler) patchVacancy(c *gin.Context) {

}
