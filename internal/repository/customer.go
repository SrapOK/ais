package repository

import (
	model "kis/internal/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetById(id uint) (model.Customer, error) {
	customer := model.Customer{}
	res := r.db.Find(&customer, id).Error
	return customer, res
}
