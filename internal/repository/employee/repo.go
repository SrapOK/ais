package employee

import (
	model "kis/internal/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) GetById(id uint) (model.Employee, error) {
	employee := model.Employee{}
	err := r.db.Find(&employee, id).Error
	return employee, err
}

func (r *EmployeeRepository) Update(employee model.Employee) (model.Employee, error) {
	err := r.db.Save(&employee).Error
	return employee, err
}
