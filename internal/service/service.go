package service

import (
	model "kis/internal/models"
	"kis/internal/repository"
)

type Vacancy interface {
	GetVacancies(params model.QueryDTO) ([]model.VacancyDTO, error)
	CreateVacancy(obj model.Vacancy) (uint, error)
}

type Service struct {
	Vacancy
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Vacancy: NewVacancyService(repos.Vacancy, repos.Customer)}
}
