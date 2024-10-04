package repository

import (
	model "kis/internal/models"

	"gorm.io/gorm"
)

type Vacancy interface {
	GetById(id uint) (model.Vacancy, error)
	// updateVacancy(obj model.Vacancy)
	GetBySearchTerms(params model.VacancySearchTermsDTO) ([]model.Vacancy, error)
	Create(vacancy model.Vacancy) (uint, error)
}

type Customer interface {
	GetById(id uint) (model.Customer, error)
}

type Repository struct {
	Vacancy  Vacancy
	Customer Customer
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Vacancy: NewVacancyRepository(db), Customer: NewCustomerRepository(db)}
}
