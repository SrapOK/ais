package contracts

import model "kis/internal/models"

type vacanciesResponseDTO struct {
	Status string             `json:"status"`
	Result []model.VacancyDTO `json:"result"`
	Page   int                `json:"page" validate:"required"`
}

func newVacanciesResponseDTO(vacancies []model.VacancyDTO, page int) *vacanciesResponseDTO {

	return &vacanciesResponseDTO{Status: "success", Result: vacancies, Page: page}
}
