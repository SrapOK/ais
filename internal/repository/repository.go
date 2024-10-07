package repository

import (
	model "kis/internal/models"
	"kis/internal/repository/bookmark"
	"kis/internal/repository/candidate"
	"kis/internal/repository/customer"
	"kis/internal/repository/employee"
	"kis/internal/repository/vacancy"

	"gorm.io/gorm"
)

type Vacancy interface {
	GetById(id uint) (model.Vacancy, error)
	Update(obj model.Vacancy) error
	SearchVacancies(employeeId uint, query string) ([]model.Vacancy, int, error)
	GetBySearchTerms(params vacancy.VacancySearchTermsDTO) ([]model.Vacancy, error)
	Create(vacancy model.Vacancy) (uint, error)
}

type Bookmark interface {
	Create(employeeId, vacancyId uint) error
	Delete(employeeId, vacancyId uint) error
}

type Customer interface {
	GetById(id uint) (model.Customer, error)
}

type Candidate interface {
	GetById(id uint) (model.Candidate, error)
}

type Employee interface {
	GetById(id uint) (model.Employee, error)
}

type Repository struct {
	Vacancy   Vacancy
	Customer  Customer
	Candidate Candidate
	Employee  Employee
	Bookmark  Bookmark
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Vacancy:   vacancy.NewVacancyRepository(db),
		Customer:  customer.NewCustomerRepository(db),
		Candidate: candidate.NewCandidateRepository(db),
		Employee:  employee.NewEmployeeRepository(db),
		Bookmark:  bookmark.NewBookmarkRepository(db),
	}
}
