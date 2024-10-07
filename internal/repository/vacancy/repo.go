package vacancy

import (
	"kis/internal/consts"
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
	err := r.db.First(&vacancy, id).Error

	return vacancy, err
}

func (r *VacancyRepository) Create(vacancy model.Vacancy) (uint, error) {
	err := r.db.Create(&vacancy).Error

	return vacancy.ID, err
}

func (r *VacancyRepository) Update(vacancy model.Vacancy) error {
	err := r.db.Save(&vacancy).Error

	return err
}

func (r *VacancyRepository) GetBySearchTerms(t VacancySearchTermsDTO) ([]model.Vacancy, error) {
	var vacancies []model.Vacancy

	offset := int(math.Abs(float64(t.Page)-1)) * t.PageSize
	var order string
	if t.NewFirst {
		order = "desc"
	} else {
		order = "asc"
	}

	query := r.db.
		Where(&model.Vacancy{
			Country:  t.Country,
			Region:   t.Region,
			City:     t.City,
			RoleName: t.Role,
			IsActive: t.Status,
		}).
		Where("salary > ?", t.SalaryFrom)

	if t.SalaryTo != 0 {
		query.Where("salary < ?", t.SalaryTo)
	}

	err := query.
		Order("created_at " + order).
		Offset(offset).
		Limit(t.PageSize).
		Preload("Grades").
		Preload("Candidates").
		Find(&vacancies).Error

	return vacancies, err
}

func (r *VacancyRepository) SearchVacancies(employeeId uint, query string) ([]model.Vacancy, int, error) {
	var vacancies []model.Vacancy
	offset := 0

	err := r.db.Table("vacancy").
		Order("vacancy.created_at ASC").
		Offset(offset).
		Limit(consts.PAGE_SIZE).
		Joins("LEFT JOIN customer c ON vacancy.customer_id = c.id").
		Where(` setweight(to_tsvector('russian', vacancy.title), 'A') || 
				setweight(to_tsvector('russian', vacancy.role_name), 'B') || 
				setweight(to_tsvector('russian', c.name), 'C') ||
				setweight(to_tsvector('russian', vacancy.city), 'D')
				@@ to_tsquery('russian', ?)`, query).
		Preload("Grades").
		Preload("Candidates").
		Scan(&vacancies).Error

	return vacancies, 1, err
}
