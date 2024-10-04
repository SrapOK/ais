package repository

import (
	model "kis/internal/models"
	"math"

	"gorm.io/gorm"
)

type VacancyRepository struct {
	db *gorm.DB
}

func NewVacancyRepository(db *gorm.DB) *VacancyRepository {
	return &VacancyRepository{db: db}
}

func (r *VacancyRepository) GetById(id uint) (model.Vacancy, error) {
	vacancy := model.Vacancy{}
	result := r.db.First(&vacancy, id).Error

	return vacancy, result
}

func (r *VacancyRepository) Create(vacancy model.Vacancy) (uint, error) {
	result := r.db.Create(&vacancy).Error

	return vacancy.ID, result
}

func (r *VacancyRepository) Update(vacancy model.Vacancy) error {
	result := r.db.Save(&vacancy)

	return result.Error
}

func (r *VacancyRepository) GetBySearchTerms(t model.VacancySearchTermsDTO) ([]model.Vacancy, error) {
	var vacancies = []model.Vacancy{}

	offset := int(math.Abs(float64(t.Page)-1)) * t.PageSize
	var order string
	if t.NewFirst {
		order = "desc"
	} else {
		order = "asc"
	}

	query := r.db.
		Where(&model.Vacancy{Country: t.Country, Region: t.Region, City: t.City, RoleName: t.Role, IsActive: t.Status}).
		Where("salary > ?", t.SalaryFrom)

	if t.SalaryTo != 0 {
		query.Where("salary < ?", t.SalaryTo)
	}

	result := query.
		Order("created_at " + order).
		Offset(offset).
		Preload("Grades").
		Preload("Candidates").
		Find(&vacancies).Error

	return vacancies, result
}
