package service

import (
	model "kis/internal/models"
	"kis/internal/repository"
)

type Vacancy interface {
	GetVacancies(params model.QueryDTO) ([]model.VacancyDTO, error)
	CreateVacancy(obj model.Vacancy) (uint, error)
	SearchVacancies(employeeId uint, query string) ([]model.VacancyDTO, int, error)
	UpdateVacancyField(id uint, op string, path string, value any) error
}

type Bookmark interface {
	CreateBookmark(employeeId, vacancyId uint) error
	DeleteBookmark(employeeId, vacancyId uint) error
}

type Service struct {
	Vacancy
	Bookmark
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Vacancy:  NewVacancyService(repos.Vacancy, repos.Customer, repos.Employee),
		Bookmark: NewBookmarkService(repos.Bookmark),
	}
}
