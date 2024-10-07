package candidate

import (
	model "kis/internal/models"

	"gorm.io/gorm"
)

type CandidateRepository struct {
	db *gorm.DB
}

func NewCandidateRepository(db *gorm.DB) *CandidateRepository {
	return &CandidateRepository{db: db}
}

func (r *CandidateRepository) GetById(id uint) (model.Candidate, error) {
	candidate := model.Candidate{}
	err := r.db.Find(&candidate, id).Error

	return candidate, err
}
