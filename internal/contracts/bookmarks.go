package contracts

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookmarkParams struct {
	VacancyId  int `form:"vacancy_id" json:"vacancy_id" binding:"required"`
	EmployeeId int `form:"employee_id" json:"employee_id" binding:"required"`
}

func (h *Handler) createBookmark(c *gin.Context) {
	body := new(bookmarkParams)

	if err := c.ShouldBindJSON(body); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withBadRequest()).Values())
		return
	}

	if err := h.services.Bookmark.CreateBookmark(uint(body.EmployeeId), uint(body.VacancyId)); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withInternalServerError()).Values())
		return
	}

	c.JSON(newResponse(withSuccess(http.StatusCreated, "Вы успешно добавили вакансию в избранное")).Values())
}

func (h *Handler) deleteBookmark(c *gin.Context) {
	body := new(bookmarkParams)

	if err := c.ShouldBindJSON(body); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withBadRequest()).Values())
		return
	}

	if err := h.services.Bookmark.DeleteBookmark(uint(body.EmployeeId), uint(body.VacancyId)); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(newResponse(withInternalServerError()).Values())
		return
	}

	c.JSON(newResponse(withSuccess(http.StatusOK, "Вы успешно убрали вакансию из избранного")).Values())
}
