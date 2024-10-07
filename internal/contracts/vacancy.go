package contracts

import (
	model "kis/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SearchParams struct {
	EmployeeId uint   `form:"employee_id" json:"employee_id" binding:"required"`
	Query      string `form:"q" json:"q" binding:"required"`
}

type PatchParams struct {
	Op    string `form:"op" json:"op" binding:"required"`
	Path  string `form:"path" json:"path" binding:"required"`
	Value string `form:"value" json:"value" binding:"required"`
}

type PatchUri struct {
	ID int `uri:"id" binding:"required"`
}

func (h *Handler) getVacanciesByQuery(c *gin.Context) {
	params := new(model.QueryDTO)

	if err := c.BindQuery(params); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withBadRequest()).Values())
		return
	}

	vacancies, err := h.services.GetVacancies(*params)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withInternalServerError()).Values())
		return
	}

	c.JSON(newResponse(withResult(vacancies, params.Page)).Values())
}

func (h *Handler) createVacancy(c *gin.Context) {
	vacancy := new(model.Vacancy)

	err := c.BindJSON(vacancy)

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withBadRequest()).Values())
		return
	}

	id, err := h.services.CreateVacancy(*vacancy)
	if err != nil {
		c.AbortWithStatusJSON(newResponse(withInternalServerError()).Values())
	}

	c.JSON(newResponse(withSuccess(http.StatusCreated, strconv.Itoa(int(id)))).Values())
}

func (h *Handler) patchVacancies(c *gin.Context) {
	query := PatchParams{}
	if err := c.ShouldBindJSON(&query); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withBadRequest()).Values())
		return
	}

	uri := PatchUri{}
	if err := c.ShouldBindUri(&uri); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withBadRequest()).Values())
		return
	}

	if query.Op != "replace" {
		c.AbortWithStatusJSON(newResponse(withBadRequest()).Values())
		return
	}

	err := h.services.Vacancy.UpdateVacancyField(uint(uri.ID), query.Op, query.Path, query.Value)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withInternalServerError()).Values())
		return
	}

	c.JSON(newResponse(withSuccess(http.StatusOK, "Вы успешно сменили статус вакансии")).Values())
}

func (h *Handler) searchVacancies(c *gin.Context) {
	query := SearchParams{}

	if err := c.ShouldBindQuery(&query); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withBadRequest()).Values())
		return
	}

	vacancies, page, err := h.services.Vacancy.SearchVacancies(uint(query.EmployeeId), query.Query)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withInternalServerError()).Values())
		return
	}

	c.JSON(newResponse(withResult(vacancies, page)).Values())
}
